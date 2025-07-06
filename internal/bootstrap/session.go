package bootstrap

import (
	"context"
	"time"

	"1337b04rd/internal/adapters/driven/redis"
	rickandmorty "1337b04rd/internal/adapters/driven/rickApi"
	"1337b04rd/internal/ports/inbound"
	rickCharacter "1337b04rd/internal/service/session/rickdal"
	session "1337b04rd/internal/service/session/sessionGenerator"
)

func (app *myApp) initSession(ctx context.Context, sessionCfg inbound.SessionConfig, redisCfg inbound.RedisConfig) (inbound.SessionInter, error) {
	rickRedis, err := redis.InitRickRedis(ctx, redisCfg)
	if err != nil {
		return nil, err
	}

	app.wg.Add(1)
	app.srv.RegisterOnShutDown(func() {
		defer app.wg.Done()
		rickRedis.CloseRedis()
	})

	// init rickandmorty api
	rickApi := rickandmorty.InitRickApi(10 * time.Second)

	// init rick service (first layer)
	rickService, err := rickCharacter.InitRickAndMortyRedis(ctx, rickApi, rickRedis)
	if err != nil {
		return nil, err
	}

	sessionRedis, err := redis.InitSessionRedis(ctx, redisCfg, sessionCfg.GetDuration())
	if err != nil {
		return nil, err
	}

	app.wg.Add(1)
	app.srv.RegisterOnShutDown(func() {
		defer app.wg.Done()
		sessionRedis.CloseRedis()
	})

	// init rick service (second layer)
	return session.InitSession(sessionRedis, rickService), nil
	// fmt.Println(sessionService.NewSession(ctx))
	// sessionMiddleware := middleware.InitSession(sessionCfg, sessionService)
	// return sessionMiddleware, nil
}
