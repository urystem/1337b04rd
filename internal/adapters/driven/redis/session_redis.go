package redis

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"time"

	"1337b04rd/internal/domain"
	"1337b04rd/internal/ports/inbound"
	"1337b04rd/internal/ports/outbound"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// inrfa
type sessionRedis struct {
	*redis.Client
	ttl time.Duration
}

func InitSessionRedis(ctx context.Context, red inbound.RedisConfig, ttl time.Duration) (outbound.SessionRedisInter, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redisDB:" + red.GetAddr(),
		Password: red.GetPass(),
		DB:       1,
	})
	return &sessionRedis{rdb, ttl}, rdb.Ping(ctx).Err()
}

type save struct {
	UserName  string
	AvatarURL string
}

// methods (adapters)
func (sr *sessionRedis) GetUserBySession(ctx context.Context, session uuid.UUID) (*domain.Session, error) {
	b, err := sr.Get(ctx, session.String()).Bytes()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return nil, err // сессия не найдена
		}
		return nil, err
	}

	data := new(save)
	err = gob.NewDecoder(bytes.NewReader(b)).Decode(data)
	if err != nil {
		return nil, err
	}

	return &domain.Session{Uuid: session, Name: data.UserName, AvatarURL: data.AvatarURL}, nil
}

func (sr *sessionRedis) SetSession(ctx context.Context, session *domain.Session) error {
	keyUUID := session.Uuid.String()
	data := &save{session.Name, session.AvatarURL}

	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(data)
	if err != nil {
		return err
	}

	return sr.Set(ctx, keyUUID, buf.Bytes(), sr.ttl).Err()
}

func (sr *sessionRedis) CloseRedis() error {
	return sr.Close()
}
