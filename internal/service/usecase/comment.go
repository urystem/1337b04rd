package usecase

import (
	"context"

	"1337b04rd/internal/domain"
)

func (u *usecase) GetCommentImage(ctx context.Context, objName string) (*domain.OutputObject, error) {
	return u.s3.GetComment(ctx, objName)
}
