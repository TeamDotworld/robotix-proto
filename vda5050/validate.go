package vda5050

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/santhosh-tekuri/jsonschema/v6"
)

// Schemas holds the official VDA5050 3.0.0 JSON schemas, embedded so that
// validation never depends on files being present at runtime. They are the
// same files the Go types are generated from, which keeps the two in step.
//
//go:embed schemas/*.schema
var Schemas embed.FS

// schemaFile maps a topic to its schema, per §4.3 Table 2.
var schemaFile = map[Topic]string{
	TopicOrder:          "schemas/order.schema",
	TopicInstantActions: "schemas/instantActions.schema",
	TopicState:          "schemas/state.schema",
	TopicVisualization:  "schemas/visualization.schema",
	TopicConnection:     "schemas/connection.schema",
	TopicFactsheet:      "schemas/factsheet.schema",
	TopicZoneSet:        "schemas/zoneSet.schema",
	TopicResponses:      "schemas/responses.schema",
}

var (
	compileOnce sync.Once
	compiled    map[Topic]*jsonschema.Schema
	compileErr  error
)

func compileSchemas() {
	compiler := jsonschema.NewCompiler()
	compiled = make(map[Topic]*jsonschema.Schema, len(schemaFile))
	for topic, path := range schemaFile {
		raw, err := Schemas.ReadFile(path)
		if err != nil {
			compileErr = fmt.Errorf("vda5050: reading %s: %w", path, err)
			return
		}
		doc, err := jsonschema.UnmarshalJSON(bytes.NewReader(raw))
		if err != nil {
			compileErr = fmt.Errorf("vda5050: parsing %s: %w", path, err)
			return
		}
		url := "https://vda5050.org/" + string(topic) + ".schema"
		if err := compiler.AddResource(url, doc); err != nil {
			compileErr = fmt.Errorf("vda5050: adding %s: %w", path, err)
			return
		}
		sch, err := compiler.Compile(url)
		if err != nil {
			compileErr = fmt.Errorf("vda5050: compiling %s: %w", path, err)
			return
		}
		compiled[topic] = sch
	}
}

// SchemaFor returns the compiled JSON schema for a topic.
func SchemaFor(topic Topic) (*jsonschema.Schema, error) {
	compileOnce.Do(compileSchemas)
	if compileErr != nil {
		return nil, compileErr
	}
	sch, ok := compiled[topic]
	if !ok {
		return nil, fmt.Errorf("vda5050: no schema for topic %q", topic)
	}
	return sch, nil
}

// ValidationError reports that a payload does not conform to its schema.
type ValidationError struct {
	Topic  Topic
	Detail string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("vda5050: %s message failed schema validation: %s", e.Topic, e.Detail)
}

// ValidateRaw checks a raw JSON payload against the schema for a topic.
//
// Validation is opt-in per direction. Outgoing messages are best validated in
// tests and in development, where a schema violation is a bug worth failing
// loudly on. Incoming messages are better validated permissively: a vehicle
// that is fractionally out of spec is still more useful under supervision
// than one dropped from the fleet, so callers typically log the error and
// carry on rather than discarding the message.
func ValidateRaw(topic Topic, payload []byte) error {
	sch, err := SchemaFor(topic)
	if err != nil {
		return err
	}
	inst, err := jsonschema.UnmarshalJSON(bytes.NewReader(payload))
	if err != nil {
		return &ValidationError{Topic: topic, Detail: "payload is not valid JSON: " + err.Error()}
	}
	if err := sch.Validate(inst); err != nil {
		return &ValidationError{Topic: topic, Detail: err.Error()}
	}
	return nil
}

// Validate marshals a message and checks it against the schema for a topic.
// Use this on the fleet-control side before publishing an order, so that a
// malformed order is caught here rather than rejected by the vehicle.
func Validate(topic Topic, msg any) error {
	raw, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("vda5050: marshalling %s message: %w", topic, err)
	}
	return ValidateRaw(topic, raw)
}
