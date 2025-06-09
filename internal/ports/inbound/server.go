package inbound

import (
	"context"
	"net/http"
)

type ServerInter interface {
	SetHandler(hand http.Handler)
	Run() error
	ShutdownGracefully(ctx context.Context) error
	RegisterOnShutDown(f func())
	CloseServer() error
}
