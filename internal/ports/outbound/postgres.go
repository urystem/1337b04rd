package outbound

import (
	"context"

	"1337b04rd/internal/domain"
)

type PostGres interface {
	PgxPost
	CloseDB()
	// CreateComment() error
}

type PgxPost interface {
	SelectPosts(context.Context) ([]domain.PostNonContent, error)
	InsertPost(ctx context.Context, post *domain.InsertPost) (uint64, error)
	DeletePost(ctx context.Context, id uint64) error
}
