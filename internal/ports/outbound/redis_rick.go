package outbound

import (
	"context"

	"1337b04rd/internal/domain"
)

type RickAndMortyRedisInter interface {
	SetCharacter(ctx context.Context, character *domain.Character) error
	GetAndDelRandomCharacter(ctx context.Context) (*domain.Character, error)
}

type UseCaseRickAndMorty interface {
	GetCharacter(ctx context.Context) (*domain.Character, error)
}
