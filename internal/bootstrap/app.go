package bootstrap

import (
	"context"
	"errors"
	"sync"

	"1337b04rd/internal/adapters/driver/http/server"
	"1337b04rd/internal/ports/inbound"
)

// DI container
type myApp struct {
	srv inbound.ServerInter
	wg  sync.WaitGroup
}

func InitApp(ctx context.Context, cfg inbound.Config) (inbound.AppInter, error) {
	// init server
	srvCfg := cfg.GetServerCfg()
	mySrv := server.InitServer(srvCfg)

	// init app
	app := &myApp{srv: mySrv}

	// init middleware to app
	sessionCfg := cfg.GetSessionConfig()
	redisCfg := cfg.GetRedisConfig()

	middleware, err := app.middleWare(ctx, sessionCfg, redisCfg)
	if err != nil {
		return nil, errors.Join(err, app.Shutdown(ctx))
	}

	
	middleware.CheckOrSetSession(nil)
	return app, nil
}

func (app *myApp) Shutdown(ctx context.Context) error {
	err := app.srv.ShutdownGracefully(ctx)
	if err == nil {
		app.wg.Wait()
	}
	return err
}

func (app *myApp) Run() error {
	return app.srv.Run()
}
