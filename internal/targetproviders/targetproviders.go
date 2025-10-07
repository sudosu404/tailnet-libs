// SPDX-FileCopyrightText: 2025 Hector <hector@email.gnx>
// SPDX-License-Identifier: AGPL3

package targetproviders

import (
	"context"

	"github.com/sudosu404/gnx-cli/internal/model"
)

type (
	// TargetProvider interface to be implemented by all target providers
	TargetProvider interface {
		WatchEvents(ctx context.Context, eventsChan chan TargetEvent, errChan chan error)
		GetDefaultProxyProviderName() string
		Close()
		AddTarget(id string) (*model.Config, error)
		DeleteProxy(id string) error
	}
)

const (
	ActionStartProxy ActionType = iota + 1
	ActionStopProxy
	ActionRestartProxy
	ActionStartProt
	ActionStopPrort
	ActionRestartPort
)

type (
	ActionType int

	TargetEvent struct {
		TargetProvider TargetProvider
		ID             string
		Action         ActionType
	}
)
