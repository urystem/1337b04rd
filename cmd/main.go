package main

import (
	"context"
	"time"

	"1337b04rd/internal/adapters/driven/redis"
	rickandmorty "1337b04rd/internal/adapters/driven/rickApi"
	"1337b04rd/internal/adapters/driver/http/middleware"
	rickAndMorty "1337b04rd/internal/service/rickCharacter"
	"1337b04rd/internal/service/session"
	"1337b04rd/pkg/config"
)

func main() {
	ctxBack := context.Background()

	cfg := config.Load()

	redisConf := cfg.GetRedisConfig()

	// init rick redis #1
	rickRedis, err := redis.InitRickRedis(ctxBack, redisConf)
	// init rickandmorty redisDB
	rickApi := rickandmorty.InitRickApi(10 * time.Second)
	// init rick service (first layer)
	rickService := rickAndMorty.InitRickAndMortyCase(rickApi, rickRedis)

	// init session config
	sessionConf := cfg.GetSessionConfig()

	// init session redis #2
	sessionRedis, err := redis.InitSessionRedis(ctxBack, redisConf, sessionConf.GetDuration())

	// init rick service (second layer)
	sessionService := session.InitSession(sessionRedis, rickService)

	// session cookie middleware
	sessionMiddleware := middleware.InitSession(sessionConf, sessionService)

	cfg.GetDBConfig()
}
