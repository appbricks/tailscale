package controlclient

import "net/http"

// mycs space node hook provides mycs
// context to tailscale daemon services
type mycsNodeControlService interface {
	// configures TLS required
	// to connect to MyCS node
	// API
	ConfigureHTTPClient(url string, httpClient *http.Client) error
}

var MyCSNodeControlService mycsNodeControlService