package inbound

import (
	"context"
	"net/http"

	"1337b04rd/internal/domain"
)

// to router
type MiddleWareInter interface {
	CheckOrSetSession(next http.Handler) http.Handler
}

// to handler
type MiddlewareSessionContext interface {
	FromContext(ctx context.Context) (*domain.Session, bool)
}

type SessionMiddleware interface {
	MiddleWareInter
	MiddlewareSessionContext
}
