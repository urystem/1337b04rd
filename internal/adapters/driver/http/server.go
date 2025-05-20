package http

import (
	"context"
	"log/slog"
	"net/http"

	"1337b04rd/internal/core/ports/inbound"
)

type myServer struct {
	logger *slog.Logger
	http.Server
}

func NewServer(logger *slog.Logger, addr string, handler http.Handler) inbound.ServerInter {
	srv := &myServer{logger: logger}
	srv.Addr = addr
	srv.Handler = handler
	return srv
}

// func (r *myServer) Run(ctx context.Context) error {
// 	srv := http.Server{Addr: r.Addr, Handler: r.Handler}
// 	errChan := make(chan error)
// 	go func() {
// 		if e := srv.ListenAndServe(); e != nil {
// 			errChan <- e
// 		}
// 	}()

// 	select {
// 	case err := <-errChan:
// 		return err
// 	case <-ctx.Done():
// 		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 		defer cancel()
// 		return srv.Shutdown(shutdownCtx)
// 	}
// }

func (srv *myServer) Serve() error {
	return srv.ListenAndServe()
}

func (srv *myServer) Shutdown(ctx context.Context) error {
	return srv.Shutdown(ctx)
}
