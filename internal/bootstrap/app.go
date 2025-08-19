package bootstrap

import (
	"context"
	"errors"
	"log/slog"
	"sync"

	"1337b04rd/internal/adapters/driver/http/handler"
	"1337b04rd/internal/adapters/driver/http/middleware"
	"1337b04rd/internal/adapters/driver/http/router"
	"1337b04rd/internal/adapters/driver/http/server"
	"1337b04rd/internal/ports/inbound"
)

// DI container
type myApp struct {
	ticker inbound.Ticker
	srv    inbound.ServerInter
	wg     sync.WaitGroup
}

func InitApp(ctx context.Context, cfg inbound.Config) (inbound.AppInter, error) {
	// init server
	srvCfg := cfg.GetServerCfg()
	mySrv := server.InitServer(srvCfg)

	// init app
	app := &myApp{srv: mySrv}

	// init session and middleware
	sessionCfg := cfg.GetSessionConfig()
	redisCfg := cfg.GetRedisConfig()

	session, err := app.initSession(ctx, sessionCfg, redisCfg)
	if err != nil {
		return nil, errors.Join(err, app.Shutdown(ctx))
	}
	middle := middleware.InitSession(sessionCfg, session)

	// init service(usecase)
	dbCfg := cfg.GetDBConfig()
	minIoCfg := cfg.GetMinIoConfig()
	useCase, err := app.initService(ctx, dbCfg, minIoCfg, session)
	if err != nil {
		return nil, errors.Join(err, app.Shutdown(ctx))
	}

	// init ticker
	app.ticker = useCase

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
	app.initTicker()
	return app.srv.ListenServe()
}
