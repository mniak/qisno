package main

import (
	"context"
	"github.com/caseymrm/menuet"
	"github.com/mniak/qisno/pkg/qisno"
)

var vpnState int
var vpnDisconnect qisno.DisconnectFunc

func (a _Application) generateVPNMenuItems(ctx context.Context) []menuet.MenuItem {
	switch vpnState {
	case 0: // Not connected
		return []menuet.MenuItem{{
			Text: "Connect to VPN",
			Clicked: func() {
				vpnState++
				var err error
				_, vpnDisconnect, err = a.VPNProvider.Connect()
				if err != nil {
					a.ShowMessage("Failed to connect to VPN")
				} else {
					vpnState++
				}
			},
		}}
	case 1: // Connecting
		return []menuet.MenuItem{{
			Text: "Connecting to VPN...",
		}}
	case 2: // Connected
		return []menuet.MenuItem{{
			Text: "VPN Connected",
			Clicked: func() {
				if vpnDisconnect != nil {
					err := vpnDisconnect()
					if err != nil {
						a.ShowMessage("Failed to disconnect")
					}
					vpnState = 0
				}
			},
			State: true,
		}}
	}
	return []menuet.MenuItem{}
}
