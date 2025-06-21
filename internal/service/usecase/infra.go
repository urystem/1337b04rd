package usecase

import (
	"1337b04rd/internal/ports/inbound"
	"1337b04rd/internal/ports/outbound"
)

type usecase struct {
	db outbound.PgxPost
	s3 outbound.MinIoInterPost
}

func InitUsecase(db outbound.PgxPost, s3 outbound.MinIoInterPost) inbound.UseCase {
	return &usecase{db: db, s3: s3}
}
