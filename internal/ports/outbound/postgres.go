package outbound

import (
	"context"

	"1337b04rd/internal/domain"

	"github.com/google/uuid"
)

type PostGres interface {
	Pgx
	CloseDB()
	// CreateComment() error
}

type Pgx interface {
	PgxPost
	PgxUser
	PgxComment
}

type PgxPost interface {
	SelectActivePosts(context.Context) ([]domain.PostNonContent, error)
	InsertPost(ctx context.Context, post *domain.InsertPost) (uint64, error)
	DeletePost(ctx context.Context, id uint64) error
	Archiver(ctx context.Context) error
	SelectArchivePosts(context.Context) ([]domain.PostNonContent, error)
	GetPost(ctx context.Context, id uint64) (*domain.PostX, error)
}

type PgxUser interface {
	InsertUser(context.Context, *domain.Session) error
	DeleteUser(ctx context.Context, sessionID uuid.UUID) error
}

type PgxComment interface {
	GetComments(ctx context.Context, postID uint64) ([]domain.Comment, error)
}
