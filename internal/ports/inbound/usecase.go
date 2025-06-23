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
}

type CommentUsecase interface {
	// ListOfPosts(context.Context) ([]domain.Post, error)
}
