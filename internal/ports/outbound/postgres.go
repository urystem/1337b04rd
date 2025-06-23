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
	GetPosts(context.Context) ([]domain.PostNonContent, error)
}
