package outbound

import (
	"context"

	"1337b04rd/internal/domain"

	"github.com/google/uuid"
)

type PostGres interface {
	PgxPost
	CloseDB()
	// CreateComment() error
}

type PgxPost interface {
	SelectActivePosts(context.Context) ([]domain.PostNonContent, error)
	InsertPost(ctx context.Context, post *domain.InsertPost) (uint64, error)
	DeletePost(ctx context.Context, id uint64) error
	InsertUser(context.Context, *domain.Session) error
	DeleteUser(ctx context.Context, sessionID uuid.UUID) error
	Archiver(ctx context.Context) error
	SelectArchivePosts(context.Context) ([]domain.PostNonContent, error)
}
