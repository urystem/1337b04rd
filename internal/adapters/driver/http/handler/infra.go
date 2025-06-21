package handler

import "1337b04rd/internal/ports/inbound"

type handler struct {
	middleware inbound.MiddlewareSessionContext
	use        inbound.UseCase
}

func InitHandler(middleware inbound.MiddlewareSessionContext, use inbound.UseCase) inbound.HandlerInter {
	return &handler{middleware, use}
}
