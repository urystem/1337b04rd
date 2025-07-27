package inbound

import (
	"context"

	"1337b04rd/internal/domain"

	"github.com/google/uuid"
)

// for redis and middleware
type SessionInter interface {
	SessionMiddlewareInter // for middleware
	SessionSeviceInter     // for service
}

type SessionMiddlewareInter interface {
	NewSession(ctx context.Context) *domain.Session
	GetSession(ctx context.Context, id string) *domain.Session
}

type SessionSeviceInter interface {
	SetSavedUUID(ctx context.Context, id uuid.UUID) error
}
