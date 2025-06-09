package main

import (
	"context"
	"log/slog"

	"1337b04rd/internal/bootstrap"
	"1337b04rd/pkg/config"
)

func main() {
	ctxBack := context.Background()

	cfg := config.Load()

	_, err := bootstrap.InitApp(ctxBack, cfg)
	if err != nil {
		slog.Error(err.Error())
	}

	// go func() {
	// 	app.Run()
	// }()
	// <-ctxBack.Done()
	// app.Shutdown(ctxBack)
}
