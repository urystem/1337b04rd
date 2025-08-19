package bootstrap

import (
	"context"

	"1337b04rd/internal/adapters/driven/minio"
	"1337b04rd/internal/adapters/driven/postgres"
	"1337b04rd/internal/ports/inbound"
	"1337b04rd/internal/service/usecase"
)

func (app *myApp) initService(ctx context.Context, dbCfg inbound.DBConfig, s3Cfg inbound.MinioCfg, session inbound.SessionSeviceInter) (inbound.UseCase, error) {
	db, err := postgres.InitDB(ctx, dbCfg)
	if err != nil {
		return nil, err
	}

	app.wg.Add(1)
	app.srv.RegisterOnShutDown(func() {
		defer app.wg.Done()
		db.CloseDB()
	})

	minio, err := minio.InitMinio(ctx, s3Cfg)
	if err != nil {
		return nil, err
	}
	return usecase.InitUsecase(db, minio, session), nil
}
