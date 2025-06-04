package outbound

import (
	"context"

	"1337b04rd/internal/domain"
)

type RickAndMortyApi interface {
	GetCharacters(ctx context.Context, p int) ([]domain.Character, error)
}
