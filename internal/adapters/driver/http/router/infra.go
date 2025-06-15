package router

import "1337b04rd/internal/ports/inbound"

type router struct {
	middleware inbound.MiddleWareInter
	handler    inbound.HandlerInter
}

func NewRoute(middle inbound.MiddleWareInter, hand inbound.HandlerInter) inbound.RouteInter {
	return &router{middle, hand}
}
