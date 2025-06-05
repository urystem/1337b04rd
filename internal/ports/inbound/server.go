package inbound

import "context"

type ServerInter interface {
	Run() error
	ShutdownGracefully(ctx context.Context) error
	RegisterOnShutDown(f func())
}
