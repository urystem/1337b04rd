package http

import (
	"net/http"

	"1337b04rd/internal/core/ports/inbound"
)

type router struct {
	inbound.HandlerInter
}

func NewRoute(hand inbound.HandlerInter) inbound.RouteInter {
	return &router{hand}
}

func (r *router) NewServe() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", r.Archive)
	return mux
}
