package rickAndMorty

import (
	"context"
	"errors"

	"1337b04rd/internal/domain"
	"1337b04rd/internal/ports/inbound"
	"1337b04rd/internal/ports/outbound"
	myerrors "1337b04rd/pkg/myErrors"
)

type rick struct {
	rickApi outbound.RickAndMortyApi
	redis   outbound.RickAndMortyRedisInter
}

func InitRickAndMortyCase(api outbound.RickAndMortyApi, redis outbound.RickAndMortyRedisInter) inbound.UseCaseRickAndMorty {
	return &rick{rickApi: api, redis: redis}
}

func (rick *rick) GetRandomCharacterAndDel(ctx context.Context) (*domain.Character, error) {
	out, err := rick.redis.GetAndDelRandomCharacter(ctx)
	if err == nil {
		return out, nil
	} else if !errors.Is(err, myerrors.ErrRickSoldOut) {
		return nil, err
	}
	err = rick.setCharacters(ctx)
	if err != nil {
		return nil, err
	}
	return rick.redis.GetAndDelRandomCharacter(ctx)
}

func (rick *rick) setCharacters(ctx context.Context) error {
	for page := 1; ; page++ {
		characters, err := rick.rickApi.GetCharacters(context.TODO(), page)
		if err != nil {
			return err
		} else if len(characters) == 0 {
			return nil
		}
		for i := range characters {
			err = rick.redis.SetCharacter(ctx, &characters[i])
			if err != nil {
				return err
			}
		}
	}
}
