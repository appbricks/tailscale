package tlsdial

import "crypto/tls"

// mycs space node hook provides mycs
// context to tailscale daemon services
type mycsClient interface {
	// configures the TLS
	ConfigureTLS(host string, tc *tls.Config) error
}

var MyCSHook mycsClient
