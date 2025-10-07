// SPDX-FileCopyrightText: 2025 Hector <hector@email.gnx>
// SPDX-License-Identifier: AGPL3

package docker

import (
	"time"
)

const (
	// Constants to be used in container labels
	LabelPrefix    = "tailnet."
	LabelIsEnabled = LabelEnable + "=true"

	// Container config labels.
	LabelEnable             = LabelPrefix + "enable"
	LabelName               = LabelPrefix + "name"
	LabelContainerAccessLog = LabelPrefix + "containeraccesslog"
	LabelProxyProvider      = LabelPrefix + "proxyprovider"
	LabelPort               = LabelPrefix + "port."
	// Tailscale
	LabelEphemeral    = LabelPrefix + "ephemeral"
	LabelRunWebClient = LabelPrefix + "runwebclient"
	LabelTsnetVerbose = LabelPrefix + "tsnet_verbose"
	LabelAuthKey      = LabelPrefix + "authkey"
	LabelAuthKeyFile  = LabelPrefix + "authkeyfile"
	LabelAutoDetect   = LabelPrefix + "autodetect"
	LabelTags         = LabelPrefix + "tags"
	// Legacy
	LabelContainerPort = LabelPrefix + "container_port"
	LabelScheme        = LabelPrefix + "scheme"
	LabelTLSValidate   = LabelPrefix + "tlsvalidate"
	// Legacy Tailscale
	LabelFunnel = LabelPrefix + "funnel"
	// Dashboard config labels
	LabelDashboardPrefix  = LabelPrefix + "dash."
	LabelDashboardVisible = LabelDashboardPrefix + "visible"
	LabelDashboardLabel   = LabelDashboardPrefix + "label"
	LabelDashboardIcon    = LabelDashboardPrefix + "icon"

	// docker only defaults
	DefaultTargetScheme = "http"

	// auto detect
	dialTimeout     = 2 * time.Second
	autoDetectTries = 5
	autoDetectSleep = 5 * time.Second

	// Port options
	PortOptionNoTLSValidate   = "no_tlsvalidate"
	PortOptionTailscaleFunnel = "tailscale_funnel"
)
