package wgengine

import (
	"golang.zx2c4.com/wireguard/device"
)

// retrieves the underlying wireguard device
// for the given tailscale wireguard engine
func GetWireguardDevice(engine Engine) *device.Device {
	e := engine.(*userspaceEngine)
	return e.wgdev
}
