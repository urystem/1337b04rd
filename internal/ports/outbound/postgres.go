package outbound

import (
	"context"

	"1337b04rd/internal/domain"
)

type PostGres interface {
	GetPosts(context.Context) ([]domain.Post, error)
	CloseDB()
	// CreateComment() error
}
