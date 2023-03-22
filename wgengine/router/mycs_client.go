package router

import "net/netip"

// mycs space node hook provides mycs
// context to tailscale daemon services
type mycsClient interface {
	// hook called by userspaceBSDRouter.Set to
	// determine if a route provided by the
	// tailscale configuration should be configured.
	// this allows overriding tailscale configured
	// routes that cause issues.
	ExcludeRoute(pfx netip.Prefix) bool
}

var MyCSHook mycsClient
