package redis

import (
	"context"
	"time"

	"1337b04rd/internal/core/ports/outbound"

	"github.com/redis/go-redis/v9"
)

// inrfa
type sessionRedis struct {
	*redis.Client
	ttl time.Duration
}

func InitSessionRedis() outbound.SessionRedisInter {
	return &sessionRedis{}
}

// methods (adapters)
func (sr *sessionRedis) CheckSession(ctx context.Context, session string) (bool, error) {
	_, err := sr.Get(ctx, session).Result()
	if err == redis.Nil {
		return false, nil // сессия не найдена
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

func (sr *sessionRedis) SetSession(ctx context.Context, session string) error {
	// dura, err := sr.TTL(ctx, "dd").Result()
	// key, err := sr.RandomKey(ctx).Result()
	return sr.Set(ctx, session, nil, sr.ttl).Err()
}

