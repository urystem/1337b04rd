package postgres

import (
	"context"
	"fmt"

	"1337b04rd/internal/domain"
)

func (db *poolDB) GetPosts(ctx context.Context) ([]domain.Post, error) {
	const query string = `
	SELECT post_id, title, has_image
		FROM posts
		ORDER BY post_time DESC`

	rows, err := db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query posts: %w", err)
	}
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		var p domain.Post
		var hasImage bool

		if err := rows.Scan(&p.ID, &p.Title, &hasImage); err != nil {
			return nil, fmt.Errorf("scan post: %w", err)
		}

		if hasImage {
			p.ImageLink = "1"
		}

		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}
	return posts, nil
}
