package bootstrap

import (
	"context"
	"time"

	"1337b04rd/internal/adapters/driven/redis"
	rickandmorty "1337b04rd/internal/adapters/driven/rickApi"
	"1337b04rd/internal/adapters/driver/http/middleware"
	"1337b04rd/internal/ports/inbound"
	rickCharacter "1337b04rd/internal/service/rickCharacter"
	"1337b04rd/internal/service/session"
)

// DI container
type app struct {
	inbound.ServerInter
}

func InitApp(ctx context.Context, cfg inbound.Config) (inbound.ServerInter, error) {
	
	redisConf := cfg.GetRedisConfig()

	// init rick redis #1
	rickRedis, err := redis.InitRickRedis(ctx, redisConf)
	if err != nil {
		return nil, err
	}

	// init rickandmorty redisDB
	rickApi := rickandmorty.InitRickApi(10 * time.Second)

	// init rick service (first layer)
	rickService := rickCharacter.InitRickAndMortyRedis(rickApi, rickRedis)

	// init session config
	sessionConf := cfg.GetSessionConfig()

	// init session redis #2
	sessionRedis, err := redis.InitSessionRedis(ctx, redisConf, sessionConf.GetDuration())

	// init rick service (second layer)
	sessionService := session.InitSession(sessionRedis, rickService)

	// session cookie middleware
	sessionMiddleware := middleware.InitSession(sessionConf, sessionService)
	return &app{}
}
