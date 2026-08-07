// Package vda5050 implements the VDA 5050 version 3.0.0 interface for the
// communication between mobile robots and a fleet control.
//
// GT Studio takes the *fleet control* (master control) role: herdIQ publishes
// `order`, `instantActions`, `zoneSet` and `responses`, and subscribes to
// `state`, `visualization`, `connection` and `factsheet`. This lets herdIQ
// drive any VDA5050-compliant vehicle alongside native GOAT robots.
//
// Layout:
//
//	vda5050            protocol types, topics, headers, schema validation
//	vda5050/transport  broker-agnostic MQTT transport (paho implementation)
//	vda5050/master     fleet-control runtime: registry, state store, orders
//
// The message types in types.gen.go are generated directly from the official
// JSON schemas published at https://github.com/VDA5050/VDA5050 so that they
// cannot drift from the specification. Regenerate with:
//
//	go generate ./pkg/vda5050/...
//
// # Optionality
//
// Every optional field in the specification is a Go pointer (or a nil-able
// slice) tagged `omitempty`. This matters: VDA5050 distinguishes "field
// absent" from "field present with value 0/false", and vehicles apply
// different defaults in each case. Never dereference an optional field
// without a nil check; use the Opt* helpers in helpers.go to construct them.
package vda5050

//go:generate python3 gen_types.py schemas types.gen.go
