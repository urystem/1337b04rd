package inbound

import (
	"context"

	"1337b04rd/internal/domain"
)

type UseCase interface {
	PostUsecase
	CommentUsecase
	UserUseCase
	Ticker
}

type PostUsecase interface {
	ListOfActivePosts(context.Context) ([]domain.PostNonContent, error)
	GetPostImage(ctx context.Context, objName string) (*domain.OutputObject, error)
	CreatePost(ctx context.Context, form *domain.Form) error
	ListOfArchivePosts(context.Context) ([]domain.PostNonContent, error)
}

type CommentUsecase interface {
	// ListOfPosts(context.Context) ([]domain.Post, error)
}

type UserUseCase interface {
	AddUserToDB(ctx context.Context, ses *domain.Session) error
}

type Ticker interface {
	Archiver(ctx context.Context) error
}
