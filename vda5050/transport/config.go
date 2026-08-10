package transport

import (
	"hash/fnv"
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
// at it, and the service starts speaking VDA5050; leave it unset and it
// behaves exactly as before. VDA5050_ENABLED=false forces it off even when a
// URL is configured, which is useful for disabling the integration in one
// environment without editing shared broker settings.
//
// clientPrefix names the connecting service, e.g. "gtstudio". A broker evicts
// any existing session using the same client ID, so two services sharing a
// prefix on one host would repeatedly disconnect each other.
func ConfigFromEnv(clientPrefix string) Config {
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
		cfg.ClientID = defaultClientID(clientPrefix)
	}
	return cfg
}

// defaultClientID derives a broker client ID that is stable across restarts
// but distinct per service and host. Stability matters: an ID that changed on
// every restart would leave orphaned sessions on the broker.
func defaultClientID(prefix string) string {
	if prefix == "" {
		prefix = "vda5050"
	}
	id := prefix
	if host, err := os.Hostname(); err == nil && host != "" {
		id += "-" + host
	}

	// MQTT 3.1.1 brokers may reject client IDs longer than 23 characters, and
	// a broker evicts any existing session using the same ID. Plain truncation
	// would therefore make two long hostnames collide and repeatedly kick each
	// other off, so the tail is replaced with a hash of the full value.
	const maxLen = 23
	if len(id) > maxLen {
		sum := fnv.New32a()
		_, _ = sum.Write([]byte(id))
		suffix := "-" + strconv.FormatUint(uint64(sum.Sum32()), 36)
		keep := maxLen - len(suffix)
		if keep < 1 {
			keep = 1
		}
		id = id[:keep] + suffix
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
