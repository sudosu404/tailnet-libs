// SPDX-FileCopyrightText: 2025 Hector <hector@email.gnx>
// SPDX-License-Identifier: AGPL3

package dashboard

import (
	"sync"

	"github.com/sudosu404/gnx-cli/internal/core"
	"github.com/sudosu404/gnx-cli/internal/model"
	"github.com/sudosu404/gnx-cli/internal/proxymanager"
	"github.com/sudosu404/gnx-cli/internal/ui/pages"
	"github.com/sudosu404/gnx-cli/web"

	"github.com/rs/zerolog"
)

type Dashboard struct {
	Log        zerolog.Logger
	HTTP       *core.HTTPServer
	pm         *proxymanager.ProxyManager
	sseClients map[string]*sseClient
	mtx        sync.RWMutex
}

func NewDashboard(http *core.HTTPServer, log zerolog.Logger, pm *proxymanager.ProxyManager) *Dashboard {
	dash := &Dashboard{
		Log:        log.With().Str("module", "dashboard").Logger(),
		HTTP:       http,
		pm:         pm,
		sseClients: make(map[string]*sseClient),
	}

	go dash.streamProxyUpdates()

	return dash
}

// AddRoutes method add dashboard related routes to the http server
func (dash *Dashboard) AddRoutes() {
	dash.HTTP.Get("/stream", dash.streamHandler())
	dash.HTTP.Get("/", web.Static)
}

// index is the HandlerFunc to index page of dashboard
func (dash *Dashboard) renderList(ch chan SSEMessage) {
	dash.mtx.RLock()
	defer dash.mtx.RUnlock()

	// force remove elements of proxy-list inn case of client reconnect
	ch <- SSEMessage{
		Type:    EventRemoveMessage,
		Message: "#proxy-list>*",
	}

	proxies := dash.pm.GetProxies()
	_ = proxies
	for name, p := range dash.pm.Proxies {
		if p.Config.Dashboard.Visible {
			dash.renderProxy(ch, name, EventAppend)
		}
	}

	dash.streamSortList(ch)
}

func (dash *Dashboard) renderProxy(ch chan SSEMessage, name string, ev EventType) {
	p, ok := dash.pm.GetProxy(name)
	if !ok {
		return
	}

	status := p.GetStatus()

	url := p.GetURL()
	if status == model.ProxyStatusAuthenticating {
		url = p.GetAuthURL()
	}

	icon := p.Config.Dashboard.Icon
	if icon == "" {
		icon = model.DefaultDashboardIcon
	}

	label := p.Config.Dashboard.Label
	if label == "" {
		label = name
	}

	ports := make([]model.PortConfig, len(p.Config.Ports))
	i := 0
	for _, target := range p.Config.Ports {
		ports[i] = target
		i++
	}

	enabled := status == model.ProxyStatusAuthenticating || status == model.ProxyStatusRunning

	a := pages.ProxyData{
		Enabled:     enabled,
		Name:        name,
		URL:         url,
		ProxyStatus: status,
		Icon:        icon,
		Label:       label,
		Ports:       ports,
	}

	ch <- SSEMessage{
		Type: ev,
		Comp: pages.Proxy(a),
	}
}
