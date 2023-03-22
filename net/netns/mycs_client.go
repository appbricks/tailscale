package netns

// mycs space node hook provides mycs
// context to tailscale daemon services
type mycsClient interface {
	// hook called by netns_darwin_tailscaled.controlLogf to
	// determine setsockoptInt should be called on the interface.
	// netns_darwin_tailscaled.controlLogf binds output requests
	// to the default interface. this causes issues with outbound
	// traffic to control server and/or wireguard tunnel packets
	// when the default route is a space node when it is set
	// as the exit node. i.e. when the default route is set to
	// be the tunnel interface
	IgnoreSetsockoptInt() bool
}

var MyCSHook mycsClient
