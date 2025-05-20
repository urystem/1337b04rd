package inbound

import "context"

type ServerInter interface {
	Serve() error
	Shutdown(ctx context.Context) error
}
