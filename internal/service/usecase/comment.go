package usecase

import (
	"context"

	"1337b04rd/internal/domain"
)

func (u *usecase) GetCommentImage(ctx context.Context, objName string) (*domain.OutputObject, error) {
	return u.s3.GetComment(ctx, objName)
}

func (u *usecase) CreateComment(ctx context.Context, form *domain.CommentForm) error {
	insert := &domain.InsertComment{
		HasImage: form.File != nil,
	}

	insert.PostID = form.PostID
	insert.User = form.User
	insert.Content = form.Content
	commentID, err := u.db.InsertComment(ctx, insert)
	if err != nil {
		return err
	}
	return u.imageSaver(ctx, form.File, commentID)
}
