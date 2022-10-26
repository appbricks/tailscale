package controlbase

// mycs space node hook provides mycs
// context to tailscale daemon services
type mycsClient interface {
	// configures TLS required
	// to connect to MyCS node
	// API
	ConfigureHTTP(url string, ht interface{}) error
}

var MyCSHook mycsClient