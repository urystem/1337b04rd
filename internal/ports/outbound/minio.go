package outbound

import (
	"context"

	"1337b04rd/internal/domain"
)

type MinioInter interface {
	minIoInterPost
	minIoInterComment
}

type minIoInterPost interface {
	PutPost(ctx context.Context, obj *domain.InPutObject) error
	DelPost(ctx context.Context, objName string) error
	GetPost(ctx context.Context, objName string) (*domain.OutputObject, error)
}

type minIoInterComment interface {
	PutComment(ctx context.Context, obj *domain.InPutObject) error
	DelComment(ctx context.Context, objName string) error
	GetComment(ctx context.Context, objName string) (*domain.OutputObject, error)
}
