package usecase

import (
	"context"
	"errors"
	"strconv"

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
	if !insert.HasImage {
		return nil
	}
	form.File.ObjName = strconv.FormatUint(commentID, 10)
	err = u.s3.PutComment(ctx, form.File)
	if err != nil {
		return errors.Join(err, u.db.DeleteComment(ctx, commentID))
	}
	return nil
}
