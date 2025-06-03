package middleware

import (
	"time"

	"1337b04rd/internal/ports/inbound"
)

type session struct {
	cookieName string
	ttl        time.Duration
	ser        inbound.SessionInter
}

func InitSession(conf inbound.SessionConfig, ser inbound.SessionInter) inbound.MiddleWareInter {
	return &session{cookieName: conf.GetCookieName(), ttl: conf.GetDuration(), ser: ser}
}
