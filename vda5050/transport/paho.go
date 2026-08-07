package transport

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Logger is the minimal logging surface the transport needs. herdIQ passes
// its logrus entry; tests pass nil.
type Logger interface {
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
}

type nopLogger struct{}

func (nopLogger) Infof(string, ...any)  {}
func (nopLogger) Warnf(string, ...any)  {}
func (nopLogger) Errorf(string, ...any) {}

// LastWill is the message the broker publishes if this client vanishes.
// A fleet control rarely needs one — it is the vehicles that must set a will
// on their `connection` topic — but a will lets herdIQ advertise its own
// availability to dashboards and standby instances.
type LastWill struct {
	Topic   string
	Payload []byte
	QoS     byte
	Retain  bool
}

// PahoBroker is a Broker backed by eclipse/paho.mqtt.golang, already a
// dependency of herdIQ for its gt_actions MQTT integration.
type PahoBroker struct {
	cfg    Config
	log    Logger
	client mqtt.Client

	mu   sync.RWMutex
	subs map[string]subscription
}

type subscription struct {
	qos byte
	h   Handler
}

// New returns a Broker for cfg. When cfg.URL is empty it returns Disabled, so
// that a deployment without an MQTT broker keeps working with VDA5050 simply
// switched off rather than failing to start.
func New(cfg Config, log Logger, will *LastWill) Broker {
	if !cfg.Enabled() {
		return Disabled{}
	}
	if log == nil {
		log = nopLogger{}
	}
	cfg = cfg.withDefaults()

	b := &PahoBroker{cfg: cfg, log: log, subs: make(map[string]subscription)}

	opts := mqtt.NewClientOptions().
		AddBroker(cfg.URL).
		SetClientID(cfg.ClientID).
		SetKeepAlive(cfg.KeepAlive).
		SetConnectTimeout(cfg.ConnectTimeout).
		SetCleanSession(cfg.CleanSession).
		SetAutoReconnect(true).
		SetMaxReconnectInterval(cfg.ReconnectInterval).
		SetResumeSubs(false).
		SetOrderMatters(false)

	if cfg.Username != "" {
		opts.SetUsername(cfg.Username)
		opts.SetPassword(cfg.Password)
	}
	if will != nil {
		opts.SetBinaryWill(will.Topic, will.Payload, will.QoS, will.Retain)
	}

	opts.SetConnectionLostHandler(func(_ mqtt.Client, err error) {
		b.log.Warnf("[vda5050] broker connection lost: %v", err)
	})
	// Subscriptions are re-established here rather than relying on the
	// broker's stored session: with CleanSession the broker forgets them, and
	// re-subscribing explicitly means the same code path runs on first
	// connect and on every reconnect.
	opts.SetOnConnectHandler(func(c mqtt.Client) {
		b.log.Infof("[vda5050] connected to broker %s as %s", cfg.URL, cfg.ClientID)
		b.resubscribe(c)
	})

	if cfg.CACert != "" || cfg.ClientCert != "" || cfg.InsecureSkipVerify {
		tlsCfg, err := buildTLS(cfg)
		if err != nil {
			// Returning Disabled rather than a broken client keeps a
			// misconfigured certificate from taking herdIQ down: VDA5050 goes
			// dark, native robots keep running, and the log says why.
			log.Errorf("[vda5050] TLS configuration failed, disabling VDA5050: %v", err)
			return Disabled{}
		}
		opts.SetTLSConfig(tlsCfg)
	}

	b.client = mqtt.NewClient(opts)
	return b
}

func buildTLS(cfg Config) (*tls.Config, error) {
	t := &tls.Config{InsecureSkipVerify: cfg.InsecureSkipVerify} //nolint:gosec // opt-in, documented
	if cfg.CACert != "" {
		pem, err := os.ReadFile(cfg.CACert)
		if err != nil {
			return nil, fmt.Errorf("reading CA certificate: %w", err)
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(pem) {
			return nil, fmt.Errorf("no certificates found in %s", cfg.CACert)
		}
		t.RootCAs = pool
	}
	if cfg.ClientCert != "" || cfg.ClientKey != "" {
		cert, err := tls.LoadX509KeyPair(cfg.ClientCert, cfg.ClientKey)
		if err != nil {
			return nil, fmt.Errorf("loading client key pair: %w", err)
		}
		t.Certificates = []tls.Certificate{cert}
	}
	return t, nil
}

// Connect establishes the session, returning as soon as it is up or when ctx
// is done. paho keeps retrying in the background afterwards.
func (b *PahoBroker) Connect(ctx context.Context) error {
	tok := b.client.Connect()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-tok.Done():
		if err := tok.Error(); err != nil {
			return fmt.Errorf("vda5050: connecting to broker %s: %w", b.cfg.URL, err)
		}
		return nil
	}
}

// Publish sends a message.
func (b *PahoBroker) Publish(ctx context.Context, topic string, qos byte, retain bool, payload []byte) error {
	if !b.client.IsConnected() {
		return fmt.Errorf("vda5050: not connected to broker, dropping publish to %s", topic)
	}
	tok := b.client.Publish(topic, qos, retain, payload)
	// QoS 0 is fire-and-forget: waiting for the token only measures the local
	// write, so return immediately and keep the publish path cheap. VDA5050
	// puts every high-rate topic at QoS 0 precisely for this reason.
	if qos == 0 {
		return nil
	}
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-tok.Done():
		return tok.Error()
	}
}

// Subscribe registers a handler for a topic filter. The handler is recorded
// so it can be reinstated after a reconnect.
func (b *PahoBroker) Subscribe(ctx context.Context, filter string, qos byte, h Handler) error {
	b.mu.Lock()
	b.subs[filter] = subscription{qos: qos, h: h}
	b.mu.Unlock()

	if !b.client.IsConnected() {
		// Not an error: the OnConnect handler will subscribe once the session
		// comes up, so callers can wire everything before connecting.
		return nil
	}
	tok := b.client.Subscribe(filter, qos, wrap(h))
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-tok.Done():
		if err := tok.Error(); err != nil {
			return fmt.Errorf("vda5050: subscribing to %s: %w", filter, err)
		}
		return nil
	}
}

// Unsubscribe removes a subscription.
func (b *PahoBroker) Unsubscribe(ctx context.Context, filter string) error {
	b.mu.Lock()
	delete(b.subs, filter)
	b.mu.Unlock()

	if !b.client.IsConnected() {
		return nil
	}
	tok := b.client.Unsubscribe(filter)
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-tok.Done():
		return tok.Error()
	}
}

// Connected reports the session state.
func (b *PahoBroker) Connected() bool { return b.client != nil && b.client.IsConnected() }

// disconnectGraceMillis is how long Close waits for in-flight writes before
// dropping the connection.
const disconnectGraceMillis = 250

// Close disconnects, allowing a short grace period for in-flight writes.
func (b *PahoBroker) Close() error {
	if b.client != nil && b.client.IsConnected() {
		b.client.Disconnect(disconnectGraceMillis)
	}
	return nil
}

func (b *PahoBroker) resubscribe(c mqtt.Client) {
	b.mu.RLock()
	filters := make(map[string]byte, len(b.subs))
	handlers := make(map[string]Handler, len(b.subs))
	for f, s := range b.subs {
		filters[f] = s.qos
		handlers[f] = s.h
	}
	b.mu.RUnlock()

	for f, qos := range filters {
		h := handlers[f]
		tok := c.Subscribe(f, qos, wrap(h))
		if tok.WaitTimeout(b.cfg.ConnectTimeout) && tok.Error() != nil {
			b.log.Errorf("[vda5050] re-subscribing to %s failed: %v", f, tok.Error())
			continue
		}
		b.log.Infof("[vda5050] subscribed to %s (qos %d)", f, qos)
	}
}

func wrap(h Handler) mqtt.MessageHandler {
	return func(_ mqtt.Client, m mqtt.Message) {
		h(Message{Topic: m.Topic(), Payload: m.Payload(), Retained: m.Retained()})
	}
}
