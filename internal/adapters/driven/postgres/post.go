package postgres

import (
	"context"

	"1337b04rd/internal/domain"
)

func (db *poolDB) GetPosts(ctx context.Context) ([]domain.Post, error) {
	return nil, nil
}
