package inbound

import (
	"context"

	"1337b04rd/internal/domain"
)

type UseCase interface {
	PostUsecase
	CommentUsecase
}

type PostUsecase interface {
	ListOfPosts(context.Context) ([]domain.PostNonContent, error)
	GetPostImage(ctx context.Context, objName string) (*domain.OutputObject, error)

	CreatePost(ctx context.Context, form *domain.Form) error
}

type CommentUsecase interface {
	// ListOfPosts(context.Context) ([]domain.Post, error)
}
