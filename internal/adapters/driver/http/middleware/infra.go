package middleware

import (
	"context"
	"time"

	"1337b04rd/internal/domain"
	"1337b04rd/internal/ports/inbound"
)

type sessionKey struct{}

type session struct {
	cookieName string
	ttl        time.Duration
	ser        inbound.SessionInter
	// errhand    inbound.ErrorHandler
	sessKey sessionKey
}

func InitSession(conf inbound.SessionConfig, ser inbound.SessionInter) inbound.SessionMiddleware {
	return &session{cookieName: conf.GetCookieName(), ttl: conf.GetDuration(), ser: ser}
}

func (s *session) FromContext(ctx context.Context) (*domain.Session, bool) {
	model, ok := ctx.Value(s.sessKey).(*domain.Session)
	return model, ok
	// return ctx.Value(key).(*domain.Session)
}
