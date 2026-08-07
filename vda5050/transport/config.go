package transport

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// Environment variables read by ConfigFromEnv.
const (
	EnvBrokerURL   = "VDA5050_BROKER_URL"
	EnvClientID    = "VDA5050_CLIENT_ID"
	EnvUsername    = "VDA5050_USERNAME"
	EnvPassword    = "VDA5050_PASSWORD"
	EnvCACert      = "VDA5050_CA_CERT"
	EnvClientCert  = "VDA5050_CLIENT_CERT"
	EnvClientKey   = "VDA5050_CLIENT_KEY"
	EnvInsecureTLS = "VDA5050_TLS_INSECURE"
	EnvKeepAlive   = "VDA5050_KEEPALIVE_SECONDS"
	EnvEnabled     = "VDA5050_ENABLED"
)

// ConfigFromEnv reads the broker configuration from the environment.
//
// VDA5050 is off unless a broker URL is present. That is the whole
// feature-flag mechanism: bring up an MQTT broker, point VDA5050_BROKER_URL
// at it, and herdIQ starts speaking VDA5050; leave it unset and herdIQ
// behaves exactly as it did before. VDA5050_ENABLED=false forces it off even
// when a URL is configured, which is useful for turning the integration off
// in one environment without editing the shared broker settings.
func ConfigFromEnv() Config {
	if v := os.Getenv(EnvEnabled); v != "" && !truthy(v) {
		return Config{}
	}
	url := strings.TrimSpace(os.Getenv(EnvBrokerURL))
	if url == "" {
		return Config{}
	}

	cfg := Config{
		URL:                url,
		ClientID:           os.Getenv(EnvClientID),
		Username:           os.Getenv(EnvUsername),
		Password:           os.Getenv(EnvPassword),
		CACert:             os.Getenv(EnvCACert),
		ClientCert:         os.Getenv(EnvClientCert),
		ClientKey:          os.Getenv(EnvClientKey),
		InsecureSkipVerify: truthy(os.Getenv(EnvInsecureTLS)),
		CleanSession:       true,
	}
	if s := os.Getenv(EnvKeepAlive); s != "" {
		if n, err := strconv.Atoi(s); err == nil && n > 0 {
			cfg.KeepAlive = time.Duration(n) * time.Second
		}
	}
	if cfg.ClientID == "" {
		cfg.ClientID = defaultClientID()
	}
	return cfg
}

// defaultClientID derives a broker client ID that is stable across restarts
// but distinct per herdIQ instance. Stability matters: a client ID that
// changed on every restart would leave orphaned sessions on the broker.
func defaultClientID() string {
	id := "herdiq"
	if slug := os.Getenv("HERD_SLUG_NAME"); slug != "" {
		id += "-" + slug
	}
	if host, err := os.Hostname(); err == nil && host != "" {
		id += "-" + host
	}
	// MQTT 3.1.1 brokers may reject client IDs longer than 23 characters.
	if len(id) > 23 {
		id = id[:23]
	}
	return id
}

func truthy(v string) bool {
	switch strings.ToLower(strings.TrimSpace(v)) {
	case "1", "true", "yes", "on", "enabled":
		return true
	}
	return false
}
