package ussd

import (
	"github.com/samora/ussd-go/sessionstores"
	"strings"
)

type session struct {
	store    sessionstores.Store
	routeKey string
}

func newSession(store sessionstores.Store, request *Request) *session {
	return &session{
		store:    store,
		routeKey: request.Mobile + "Route",
	}
}

func (s session) Set(r route) {
	route := r.Ctrl + "." + r.Action
	err := s.store.SetValue(s.routeKey, route)
	if err != nil {
		panicln("session: %v", err)
	}
}

func (s session) Get() route {
	rStr, err := s.store.GetValue(s.routeKey)
	if err != nil {
		panicln("session: %v", err)
	}
	routes := strings.Split(rStr, ".")
	if len(routes) != 2 {
		panicln("session: route not found")
	}
	return route{routes[0], routes[1]}
}

func (s session) Exists() bool {
	b, err := s.store.ValueExists(s.routeKey)
	if err != nil {
		panicln("session: %v", err)
	}
	return b
}

func (s session) Close() {
	s.store.DeleteValue(s.routeKey)
}
