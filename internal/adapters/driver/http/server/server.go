package server

import (
	"context"
	"fmt"
	"net/http"

	"1337b04rd/internal/ports/inbound"
)

type server struct {
	*http.Server
}

// type server *http.Server

func InitServer(hand http.Handler, cfg inbound.ServerCfg) any {
	addr := fmt.Sprintf("%d", cfg.GetPort())

	srv := http.Server{
		Addr:    addr,
		Handler: hand,
	}
	srv.ListenAndServe()
	srv.Close()
	srv.Shutdown(context.TODO())
	// srv.RegisterOnShutdown()
	return &srv
}

func (srv *server) Run() error {
	return srv.ListenAndServe()
}

func (srv *server) ShutdownGracefully(ctx context.Context) error {
	return srv.Shutdown(ctx)
}

func (srv *server) RegisterOnShutDown(f func()) {
	srv.RegisterOnShutdown(f)
}
