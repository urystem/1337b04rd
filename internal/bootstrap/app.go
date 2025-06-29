package bootstrap

import (
	"context"
	"errors"
	"log/slog"
	"sync"

	"1337b04rd/internal/adapters/driver/http/handler"
	"1337b04rd/internal/adapters/driver/http/router"
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

	middle, err := app.middleWare(ctx, sessionCfg, redisCfg)
	if err != nil {
		return nil, errors.Join(err, app.Shutdown(ctx))
	}

	// init service(usecase)
	dbCfg := cfg.GetDBConfig()
	minIoCfg := cfg.GetMinIoConfig()
	useCase, err := app.InitService(ctx, dbCfg, minIoCfg)
	if err != nil {
		return nil, errors.Join(err, app.Shutdown(ctx))
	}

	// init handler
	handler, err := handler.InitHandler(middle, useCase)
	if err != nil {
		return nil, errors.Join(err, app.Shutdown(ctx))
	}

	// init router
	router := router.NewRoute(middle, handler)

	app.srv.SetHandler(router)
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
	slog.Info("server starting")
	return app.srv.ListenServe()
}
