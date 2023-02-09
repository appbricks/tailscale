package controlbase

import "net/http"

// mycs space node hook provides mycs
// context to tailscale daemon services
type mycsClient interface {
	// configures the http transport (for example with
	// TLS required to connect to MyCS node API)
	ConfigureHTTPTransport(url string, tr *http.Transport) error
}

var MyCSHook mycsClient