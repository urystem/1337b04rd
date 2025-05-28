package redis

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"strconv"

	"1337b04rd/internal/ports/inbound"
	"1337b04rd/internal/ports/outbound"
	myerrors "1337b04rd/pkg/myErrors"

	"github.com/redis/go-redis/v9"
)

type rickAndMorty struct {
	*redis.Client
}

type rickWriterReader struct {
	name      string
	avatarURL string
}

func (rickRW *rickWriterReader) GetName() string {
	return rickRW.name
}

func (rickRW *rickWriterReader) GetAvatar() string {
	return rickRW.avatarURL
}

func InitRickRedis(red inbound.RedisConfig) outbound.RickAndMortyRedisInter {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redisDB:" + red.GetAddr(), // имя сервиса + порт                 // адрес Redis
		Password: red.GetPass(),              // пароль, если есть
		DB:       0,                          // номер БД (0 по умолчанию)
	})
	return &rickAndMorty{rdb}
}

func (rick *rickAndMorty) SetCharacter(ctx context.Context, character outbound.CharacterInputInter) error {
	idKey := strconv.FormatUint(character.GetID(), 10)
	var buf bytes.Buffer
	rickRW := rickWriterReader{character.GetName(), character.GetAvatar()}
	err := gob.NewEncoder(&buf).Encode(rickRW)
	if err != nil {
		return err
	}
	fmt.Println(buf.Bytes())
	return rick.Set(ctx, idKey, buf.Bytes(), 0).Err()
}

func (rick *rickAndMorty) GetAndDelRandomCharacter(ctx context.Context) (outbound.CharacterOutputInter, error) {
	key, err := rick.RandomKey(ctx).Result()
	if err == redis.Nil {
		return nil, myerrors.ErrRickSoldOut
	}

	b, err := rick.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}

	rickRW := &rickWriterReader{}
	err = gob.NewDecoder(bytes.NewReader(b)).Decode(rickRW)
	if err != nil {
		return nil, err
	}
	go rick.Del(ctx, key).Err()
	// err = rick.Del(ctx, key).Err()
	// if err != nil {
	// 	return nil, err
	// }
	return rickRW, nil
}
