package outbound

import "1337b04rd/internal/domain"

type RickAndMortyApi interface {
	GetCharacters(p int) ([]domain.Character, error)
}
