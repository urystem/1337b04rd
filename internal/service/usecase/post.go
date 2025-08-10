package usecase

import (
	"context"
	"errors"
	"fmt"
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
	post, comments, err := u.getPostAndComment(ctx, id)
	if err != nil {
		return nil, err
	}

	return &domain.ArchivePost{Post: *post, Comments: comments}, nil
}

func (u *usecase) GetActivePost(ctx context.Context, id uint64) (*domain.ActivePost, error) {
	post, comments, err := u.getPostAndComment(ctx, id)
	if err != nil {
		return nil, err
	}
	fmt.Println(comments)
	// Шаг 1: Создаем плоский map с Comment.ID -> CommentTree
	commentMap := make(map[uint64]*domain.CommentTree)
	for _, c := range comments {
		treeComment := &domain.CommentTree{}
		treeComment.CommentID = c.CommentID
		treeComment.UserName = c.UserName
		treeComment.AvatarURL = c.AvatarURL
		treeComment.Content = c.Content
		treeComment.HasImage = c.HasImage
		treeComment.DataTime = c.DataTime

		// save to map
		commentMap[c.CommentID] = treeComment
	}

	// Шаг 2: Собираем дерево
	var roots []domain.CommentTree
	for _, c := range comments {
		node := commentMap[c.CommentID]

		if c.ReplyToID != nil {
			parentNode, ok := commentMap[*c.ReplyToID]
			if ok {
				parentNode.Replies = append(parentNode.Replies, *node)
			} else {
				return nil, fmt.Errorf("impossible error")
			}
		} else {
			roots = append(roots, *node)
		}
	}

	return &domain.ActivePost{Post: *post, CommentTries: roots}, nil
}

func (u *usecase) getPostAndComment(ctx context.Context, id uint64) (*domain.Post, []domain.Comment, error) {
	postX, err := u.db.GetPost(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	post := &domain.Post{ID: id, PostX: *postX}

	comments, err := u.db.GetComments(ctx, id)
	if err != nil {
		return nil, nil, err
	}
	return post, comments, nil
}
