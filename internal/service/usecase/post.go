package usecase

import (
	"context"
	"fmt"

	"1337b04rd/internal/domain"
)

func (u *usecase) ListOfPosts(ctx context.Context) ([]domain.Post, error) {
	posts, err := u.db.GetPosts(ctx)
	if err != nil {
		return nil, err
	}
	for i, post := range posts {
		if post.ImageLink != "" {
			url, err := u.s3.GetPost(ctx, fmt.Sprintf("%d", post.ID))
			if err != nil {
				return nil, err
			}
			posts[i].ImageLink = url
		}
	}
	return posts, nil
}

