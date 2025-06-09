package handler

import "1337b04rd/internal/ports/inbound"

type handler struct {
	use inbound.Usecase
}

func InitHandler(use inbound.Usecase) inbound.HandlerInter {
	return &handler{use}
}
