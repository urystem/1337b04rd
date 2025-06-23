package rickAndMorty

import (
	"context"
	"sync"

	"1337b04rd/internal/domain"
	"1337b04rd/internal/ports/inbound"
	"1337b04rd/internal/ports/outbound"
)

type rick struct {
	rickApi outbound.RickAndMortyApi
	redis   outbound.RickAndMortyRedisInter
	mutex   sync.Mutex
	count   int
}

func InitRickAndMortyRedis(ctx context.Context, api outbound.RickAndMortyApi, redis outbound.RickAndMortyRedisInter) (inbound.UseCaseRickAndMorty, error) {
	myrick := &rick{rickApi: api, redis: redis}
	myrick.mutex.Lock()
	return myrick, myrick.setCharacters(ctx)
}

func (rick *rick) GetRandomCharacterAndDel(ctx context.Context) (*domain.Character, error) {
	rick.mutex.Lock()
	defer rick.unlockMutexAndCheck(ctx)
	return rick.redis.GetAndDelRandomCharacter(ctx)
}

func (rick *rick) unlockMutexAndCheck(ctx context.Context) {
	if rick.count--; rick.count == 0 {
		go rick.setCharacters(ctx)
	} else {
		rick.mutex.Unlock()
	}
}

func (rick *rick) setCharacters(ctx context.Context) error {
	defer rick.mutex.Unlock()
	for page := 1; ; page++ {
		characters, err := rick.rickApi.GetCharacters(ctx, page)
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
		rick.count += len(characters)
	}
}
