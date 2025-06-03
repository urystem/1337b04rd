package redis

import (
	"bytes"
	"context"
	"encoding/gob"
	"strconv"

	"1337b04rd/internal/domain"
	"1337b04rd/internal/ports/inbound"
	"1337b04rd/internal/ports/outbound"
	myerrors "1337b04rd/pkg/myErrors"

	"github.com/redis/go-redis/v9"
)

type rickAndMorty struct {
	*redis.Client
}

func InitRickRedis(ctx context.Context, red inbound.RedisConfig) (outbound.RickAndMortyRedisInter, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redisDB:" + red.GetAddr(), // имя сервиса + порт                 // адрес Redis
		Password: red.GetPass(),              // пароль, если есть
		DB:       0,                          // номер БД (0 по умолчанию)
	})
	return &rickAndMorty{rdb}, rdb.Ping(ctx).Err()
}

func (rick *rickAndMorty) SetCharacter(ctx context.Context, character *domain.Character) error {
	idKey := strconv.FormatUint(character.ID, 10)
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(character)
	if err != nil {
		return err
	}

	return rick.Set(ctx, idKey, buf.Bytes(), 0).Err()
}

func (rick *rickAndMorty) GetAndDelRandomCharacter(ctx context.Context) (*domain.Character, error) {
	key, err := rick.RandomKey(ctx).Result()
	if err == redis.Nil {
		return nil, myerrors.ErrRickSoldOut
	}

	b, err := rick.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	character := new(domain.Character)

	err = gob.NewDecoder(bytes.NewReader(b)).Decode(character)
	if err != nil {
		return nil, err
	}

	rick.Del(ctx, key)

	return character, nil
}
