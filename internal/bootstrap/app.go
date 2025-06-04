package bootstrap

import (
	"context"

	"1337b04rd/internal/ports/inbound"
)

// DI container
type app struct{
	
}

func InitApp(ctx context.Context, cfg inbound.Config) any {
	return &app{}
}
