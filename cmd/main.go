package main

import (
	"1337b04rd/internal/adapters/driven/redis"
	rickandmorty "1337b04rd/internal/adapters/driven/rickApi"
	"1337b04rd/internal/adapters/driver/http/middleware"
	rickAndMorty "1337b04rd/internal/service/rickCharacter"
	"1337b04rd/internal/service/session"
	"1337b04rd/pkg/config"
)

func main() {
	conf := config.InitConfig()

	redisConf := conf.GetRedisConfig()

	// init rick redis #1
	rickRedis := redis.InitRickRedis(redisConf)

	// init rickandmorty redisDB
	rickApi := rickandmorty.InitRickApi()
	// init rick service (first layer)
	rickService := rickAndMorty.InitRickAndMortyCase(rickApi, rickRedis)

	// init session config
	sessionConf := conf.GetSessionConfig()

	// init session redis #2
	sessionRedis := redis.InitSessionRedis(redisConf, sessionConf.GetDuration())

	// init rick service (second layer)
	sessionService := session.InitSession(sessionRedis, rickService)

	// session cookie middleware
	sessionMiddleware := middleware.InitSession(sessionConf, sessionService)

	conf.GetDBConfig()
}
