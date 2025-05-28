package usecase

import (
	"context"
	"errors"
	"log/slog"

	rickandmorty "1337b04rd/internal/adapters/driven/rickApi"
	"1337b04rd/internal/ports/outbound"
	myerrors "1337b04rd/pkg/myErrors"
)

type rick struct {
	outbound.RickAndMortyRedisInter
}

func InitRickAndMortyCase(redis outbound.RickAndMortyRedisInter) outbound.UseCaseRickAndMorty {
	return &rick{redis}
}

func (rick *rick) GetCharacter(ctx context.Context) (outbound.CharacterOutputInter, error) {
	out, err := rick.GetAndDelRandomCharacter(ctx)
	if err == nil {
		return out, nil
	} else if !errors.Is(err, myerrors.ErrRickSoldOut) {
		return nil, err
	}
	slog.Info("soldout rick and morty")
	err = rick.setCharacters(ctx)
	if err != nil {
		return nil, err
	}
	return rick.GetAndDelRandomCharacter(ctx)
}

func (rick *rick) setCharacters(ctx context.Context) error {
	slog.Info("starting set to redis")
	for page := 1; page < 3; page++ {
		characters, err := rickandmorty.GetCharacters(page)
		if err != nil {
			return err
		} else if len(characters) == 0 {
			slog.Info("page sold out")
			return nil
		}
		for i := range characters {
			err = rick.SetCharacter(ctx, characters[i])
			if err != nil {
				return err
			}
		}
	}
	return nil
}
