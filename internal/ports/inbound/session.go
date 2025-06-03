package inbound

import (
	"context"

	"1337b04rd/internal/domain"
)

type SessionInter interface {
	NewSession(ctx context.Context) *domain.Session
	GetSession(ctx context.Context, id string) *domain.Session
}
