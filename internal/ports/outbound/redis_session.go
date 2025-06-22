package outbound

import (
	"context"

	"1337b04rd/internal/domain"

	"github.com/google/uuid"
)

type SessionRedisInter interface {
	GetUserBySession(ctx context.Context, session uuid.UUID) (*domain.Session, error)
	SetSession(ctx context.Context, session *domain.Session) error
	CloseRedis() error
}
