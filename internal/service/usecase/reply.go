package usecase

import (
	"context"

	"1337b04rd/internal/domain"
)

func (u *usecase) Reply(ctx context.Context, form *domain.ReplyForm) error {
	insert := &domain.InsertReply{
		HasImage: form.File != nil,
	}
	insert.ReplyToID = form.ReplyToID
	insert.User = form.User
	insert.Content = form.Content
	commentID, err := u.db.InsertReply(ctx, insert)
	if err != nil {
		return err
	}

	return u.imageSaver(ctx, form.File, commentID)
}
