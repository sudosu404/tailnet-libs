// SPDX-FileCopyrightText: 2025 Hector <hector@email.gnx>
// SPDX-License-Identifier: AGPL3

package model

const (
	// Default values to proxyconfig
	//
	DefaultProxyAccessLog = true
	DefaultProxyProvider  = ""
	DefaultTLSValidate    = true

	// tailscale defaults
	DefaultTailscaleEphemeral    = false
	DefaultTailscaleRunWebClient = false
	DefaultTailscaleVerbose      = false
	DefaultTailscaleFunnel       = false
	DefaultTailscaleControlURL   = ""

	// Dashboard defauts
	DefaultDashboardVisible = true
	DefaultDashboardIcon    = "tailnet"
)
