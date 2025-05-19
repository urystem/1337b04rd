package http

import (
	"log/slog"
	"net/http"

	"1337b04rd/internal/core/ports/inbound"
)

type rest struct {
	logger  *slog.Logger
	handler inbound.HandlerInter
}

func NewRoute(logger *slog.Logger, hand inbound.HandlerInter) inbound.RouteInter {
	return &rest{}
}

func (r *rest) Serve() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", r.handler.Archive)
	return http.ListenAndServe("", mux)
}
