package usecase

import (
	"context"
	"errors"
	"strconv"

	"1337b04rd/internal/domain"
)

func (u *usecase) ListOfActivePosts(ctx context.Context) ([]domain.PostNonContent, error) {
	return u.db.SelectActivePosts(ctx)
}

func (u *usecase) GetPostImage(ctx context.Context, objName string) (*domain.OutputObject, error) {
	return u.s3.GetPost(ctx, objName)
}

func (u *usecase) CreatePost(ctx context.Context, form *domain.Form) error {
	insert := &domain.InsertPost{}
	insert.Uuid = form.Uuid
	insert.Name = form.Name
	insert.Subject = form.Subject
	insert.Content = form.Content
	insert.HasImage = form.File != nil
	postId, err := u.db.InsertPost(ctx, insert)
	if err != nil {
		return err
	}
	if !insert.HasImage {
		return nil
	}

	form.File.ObjName = strconv.FormatUint(postId, 10)
	err = u.s3.PutPost(ctx, form.File)
	if err != nil {
		return errors.Join(err, u.db.DeletePost(ctx, postId))
	}
	return nil
}

func (u *usecase) ListOfArchivePosts(ctx context.Context) ([]domain.PostNonContent, error) {
	return u.db.SelectArchivePosts(ctx)
}

func (u *usecase) GetArchivePost(ctx context.Context, id uint64) (*domain.ArchivePost, error) {
	postX, err := u.db.GetPost(ctx, id)
	if err != nil {
		return nil, err
	}

	post := &domain.Post{id, *postX}

	comments, err := u.db.GetComments(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.ArchivePost{Post: *post, Comments: comments}, nil
}
