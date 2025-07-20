package inbound

import (
	"context"

	"1337b04rd/internal/domain"
)

type UseCase interface {
	Service
	CommentUsecase
	Ticker
}

type Service interface {
	ListOfActivePosts(context.Context) ([]domain.PostNonContent, error)
	GetPostImage(ctx context.Context, objName string) (*domain.OutputObject, error)
	CreatePost(ctx context.Context, form *domain.Form) error
	ListOfArchivePosts(context.Context) ([]domain.PostNonContent, error)
	UserUseCase
	GetArchivePost(context.Context, uint64) (*domain.ArchivePost, error)
	GetCommentImage(ctx context.Context, objName string) (*domain.OutputObject, error)
	GetActivePost(context.Context, uint64) (*domain.ActivePost, error)
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
