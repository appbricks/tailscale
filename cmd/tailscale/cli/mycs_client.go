package cli

import (
	"context"
	"runtime"
	"time"

	"tailscale.com/ipn/ipnstate"
	"tailscale.com/paths"
)

func RunUp(
	ctx context.Context, 
	server string,
	options map[string]interface{},
) error {
	
	rootArgs.socket = paths.DefaultTailscaledSocket()
	upArgs.server = server
	upArgs.forceReauth = true

	if runtime.GOOS == "windows" {
		upArgs.forceDaemon = true
	}

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

func RunPing(
	ctx context.Context, 
	name, ip string,
	timeout int,
) error {

	var (
		err error

		st   *ipnstate.Status
		peer *ipnstate.PeerStatus
	)

	to := time.Second * time.Duration(timeout)
	cctx, cancel := context.WithTimeout(ctx, to)

	FINDPEER:
	for err = cctx.Err(); err == nil; {
		if st, err = localClient.Status(cctx); err != nil {
			break FINDPEER
		}
		if st.Self.HostName == name {
			break FINDPEER
		}
		for _, peer = range st.Peer {
			if peer.HostName == name {
				break FINDPEER
			}
		}
		time.Sleep(time.Second)
	}
	cancel()
	if err != nil {
		return err
	}

	pingArgs.untilDirect = true
	pingArgs.num = timeout // assume 1s between failed pings
	pingArgs.timeout = time.Second * time.Duration(timeout)
	return runPing(ctx, []string{ip})
}
