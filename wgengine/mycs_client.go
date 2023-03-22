package wgengine

import (
	"github.com/tailscale/wireguard-go/device"
)

// retrieves the underlying wireguard device
// for the given tailscale wireguard engine
func GetWireguardDevice(engine Engine) *device.Device {
	e := engine.(*userspaceEngine)
	return e.wgdev
}
