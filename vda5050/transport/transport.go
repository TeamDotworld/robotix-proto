// Package transport carries VDA5050 messages over MQTT.
//
// The specification mandates MQTT (3.1.1 or later) but says nothing about
// which broker, so this package is deliberately broker-agnostic: Broker is an
// interface and the paho-backed implementation is one of potentially several.
// Customers routinely already run Mosquitto, EMQX, HiveMQ or a cloud IoT
// endpoint, and VDA5050 traffic should join whatever they have rather than
// force a second broker into the deployment.
package transport

import (
	"context"
	"errors"
	"time"
)

// Message is one received MQTT message.
type Message struct {
	Topic   string
	Payload []byte
	// Retained is true when the broker replayed a stored message rather than
	// delivering a live one. It matters on the `connection` topic, which is
	// always retained: a retained CONNECTION_BROKEN describes the past, not
	// necessarily the present.
	Retained bool
}

// Handler processes a received message. Handlers run on the transport's
// delivery goroutines and must not block for long; hand slow work to a queue.
type Handler func(Message)

// Broker is the minimal MQTT surface VDA5050 needs. Implementations must be
// safe for concurrent use.
type Broker interface {
	// Connect establishes the session and blocks until it is up or ctx ends.
	Connect(ctx context.Context) error
	// Publish sends a message. qos is 0 or 1; retain sets the MQTT retained
	// flag. VDA5050 fixes both per topic (§4.1, §6.5).
	Publish(ctx context.Context, topic string, qos byte, retain bool, payload []byte) error
	// Subscribe registers a handler for a topic filter, which may contain the
	// MQTT wildcards + and #.
	Subscribe(ctx context.Context, filter string, qos byte, h Handler) error
	// Unsubscribe removes a subscription.
	Unsubscribe(ctx context.Context, filter string) error
	// Connected reports the current session state.
	Connected() bool
	// Close disconnects and releases resources.
	Close() error
}

// ErrDisabled is returned by every Broker method on the no-op broker. herdIQ
// treats it as "VDA5050 is switched off", not as a failure: a deployment with
// no MQTT broker configured runs exactly as it did before, driving only
// native GOAT robots.
var ErrDisabled = errors.New("vda5050: transport is disabled (no broker configured)")

// Config describes how to reach a broker. The zero value is disabled.
type Config struct {
	// URL is the broker address, e.g. "tcp://mosquitto:1883",
	// "ssl://broker.example.com:8883" or "ws://broker:9001/mqtt". Empty
	// disables VDA5050 entirely.
	URL string
	// ClientID identifies this fleet control to the broker. Brokers evict an
	// existing session when a second client connects with the same ID, so
	// this must be unique per herdIQ instance.
	ClientID string

	Username string
	Password string

	// CACert, ClientCert and ClientKey are PEM file paths for TLS. Leave
	// empty for an unauthenticated local broker.
	CACert     string
	ClientCert string
	ClientKey  string
	// InsecureSkipVerify disables server certificate verification. Intended
	// for lab brokers with self-signed certificates only.
	InsecureSkipVerify bool

	// KeepAlive is the MQTT keep-alive interval. It bounds how quickly the
	// broker notices this client has gone and publishes its last will.
	KeepAlive time.Duration
	// ConnectTimeout bounds a single connection attempt.
	ConnectTimeout time.Duration
	// ReconnectInterval is the maximum backoff between reconnect attempts.
	ReconnectInterval time.Duration
	// CleanSession starts a fresh session rather than resuming a stored one.
	// True is right for a fleet control: it re-subscribes on every connect
	// and has no use for messages queued while it was away — a state message
	// from ten minutes ago is worse than no state message.
	CleanSession bool
}

// Enabled reports whether a broker has been configured.
func (c Config) Enabled() bool { return c.URL != "" }

// withDefaults fills in values suited to a fleet control on a plant network.
func (c Config) withDefaults() Config {
	if c.KeepAlive <= 0 {
		c.KeepAlive = 20 * time.Second
	}
	if c.ConnectTimeout <= 0 {
		c.ConnectTimeout = 10 * time.Second
	}
	if c.ReconnectInterval <= 0 {
		c.ReconnectInterval = 10 * time.Second
	}
	return c
}

// Disabled is a Broker that accepts no traffic. New returns it when no broker
// is configured, so callers can hold a non-nil Broker unconditionally instead
// of nil-checking at every use.
type Disabled struct{}

func (Disabled) Connect(context.Context) error { return ErrDisabled }
func (Disabled) Publish(context.Context, string, byte, bool, []byte) error {
	return ErrDisabled
}
func (Disabled) Subscribe(context.Context, string, byte, Handler) error { return ErrDisabled }
func (Disabled) Unsubscribe(context.Context, string) error              { return ErrDisabled }
func (Disabled) Connected() bool                                        { return false }
func (Disabled) Close() error                                           { return nil }
