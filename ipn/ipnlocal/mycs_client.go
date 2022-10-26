package ipnlocal

import (
	"net/netip"
)

// mycs space node hook provides mycs
// context to tailscale daemon services
type mycsClient interface {
	// Local DNS mappings that should resolve locally in MagicDNS.
	// This is used to override DNS resolutions that need to be
	// split to bypass the tunnel.
	ResolvedDNSNames() []MyCSDNSMapping
}

type MyCSDNSMapping struct {
	Name  string
	Addrs []netip.Prefix
}

var MyCSHook mycsClient
