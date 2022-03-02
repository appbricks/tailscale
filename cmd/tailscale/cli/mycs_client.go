package cli

import (
	"context"
	"io"

	"tailscale.com/paths"
)

var MyCSOut io.Writer

func RunUp(
	ctx context.Context, 
	server string,
	options map[string]interface{},
) error {
	
	rootArgs.socket = paths.DefaultTailscaledSocket()
	upArgs.server = server
	upArgs.forceReauth = true

	for k, v := range options {
		switch k {
		case "reset":
			upArgs.reset = v.(bool)
		case "authKey":
			upArgs.authKeyOrFile = v.(string)
		case "hostname":
			upArgs.hostname = v.(string)
		case "acceptDNS":
			upArgs.acceptDNS = v.(bool)
		case "acceptRoutes":
			upArgs.acceptRoutes = v.(bool)
		case "advertiseRoutes":
			upArgs.advertiseRoutes = v.(string)
		case "advertiseDefaultRoute":
			upArgs.advertiseDefaultRoute = v.(bool)
		case "exitNodeIP":
			upArgs.exitNodeIP = v.(string)
		}
	}
	return runUp(ctx, []string{})
}

func RunLogout(
	ctx context.Context,
) error {
	return runLogout(ctx, []string{})
}

func RunDown(
	ctx context.Context,
) error {
	return runDown(ctx, []string{})
}
