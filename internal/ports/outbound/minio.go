package outbound

import (
	"context"

	"1337b04rd/internal/domain"
)

type MinioInter interface {
	MinIoInterPost
	MinIoInterComment
}

type MinIoInterPost interface {
	PutPost(ctx context.Context, obj *domain.Object) error
	DelPost(ctx context.Context, objName string) error
	GetPost(ctx context.Context, objName string) (string, error)
}

type MinIoInterComment interface {
	PutComment(ctx context.Context, obj *domain.Object) error
	DelComment(ctx context.Context, objName string) error
	GetComment(ctx context.Context, objName string) (string, error)
}
