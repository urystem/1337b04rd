package usecase

import (
	"context"
	"errors"

	rickandmorty "1337b04rd/internal/adapters/driven/rickApi"
	"1337b04rd/internal/domain"
	"1337b04rd/internal/ports/outbound"
	myerrors "1337b04rd/pkg/myErrors"
)

type rick struct {
	outbound.RickAndMortyRedisInter
}

func InitRickAndMortyCase(redis outbound.RickAndMortyRedisInter) outbound.UseCaseRickAndMorty {
	return &rick{redis}
}

func (rick *rick) GetCharacter(ctx context.Context) (*domain.Character, error) {
	out, err := rick.GetAndDelRandomCharacter(ctx)
	if err == nil {
		return out, nil
	} else if !errors.Is(err, myerrors.ErrRickSoldOut) {
		return nil, err
	}
	err = rick.setCharacters(ctx)
	if err != nil {
		return nil, err
	}
	return rick.GetAndDelRandomCharacter(ctx)
}

func (rick *rick) setCharacters(ctx context.Context) error {
	for page := 1; ; page++ {
		characters, err := rickandmorty.GetCharacters(page)
		if err != nil {
			return err
		} else if len(characters) == 0 {
			return nil
		}
		for i := range characters {
			err = rick.SetCharacter(ctx, &characters[i])
			if err != nil {
				return err
			}
		}
	}
}
