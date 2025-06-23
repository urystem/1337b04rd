package usecase

import (
	"context"

	"1337b04rd/internal/domain"
)

func (u *usecase) ListOfPosts(ctx context.Context) ([]domain.PostNonContent, error) {
	return u.db.GetPosts(ctx)
}

func (u *usecase) GetPostImage(ctx context.Context, objName string) (*domain.OutputObject, error) {
	return u.s3.GetPost(ctx, objName)
}
