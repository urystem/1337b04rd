package inbound

import (
	"context"

	"1337b04rd/internal/domain"
)

type UseCaseRickAndMorty interface {
	GetRandomCharacterAndDel(ctx context.Context) (*domain.Character, error)
}
