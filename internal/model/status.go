// SPDX-FileCopyrightText: 2025 Hector <hector@email.gnx>
// SPDX-License-Identifier: AGPL3
package model

type (
	ProxyStatus int

	ProxyEvent struct {
		ID      string
		Port    string
		AuthURL string
		Status  ProxyStatus
	}
)

const (
	ProxyStatusInitializing ProxyStatus = iota
	ProxyStatusStarting
	ProxyStatusAuthenticating
	ProxyStatusRunning
	ProxyStatusStopping
	ProxyStatusStopped
	ProxyStatusError
)

var proxyStatusStrings = []string{
	"Initializing",
	"Starting",
	"Authenticating",
	"Running",
	"Stopping",
	"Stopped",
	"Error",
}

func (s *ProxyStatus) String() string {
	return proxyStatusStrings[int(*s)]
}
