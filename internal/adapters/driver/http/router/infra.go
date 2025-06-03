package router

import "1337b04rd/internal/ports/inbound"

type router struct {
	inbound.MiddleWareInter
	inbound.HandlerInter
}

func NewRoute(hand inbound.HandlerInter, middle inbound.MiddleWareInter) inbound.RouteInter {
	return &router{}
}
