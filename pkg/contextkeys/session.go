package contextkeys

import (
	"context"

	"1337b04rd/internal/domain"
)

// type contextKey bool // уникальный тип

// var key contextKey // переменная, уникальное значение

type sessionKey struct{}

// var key = contextKey{}
var key sessionKey

func NewContext(ctx context.Context, s *domain.Session) context.Context {
	return context.WithValue(ctx, key, s)
}

func FromContext(ctx context.Context) (*domain.Session, bool) {
	s, ok := ctx.Value(key).(*domain.Session)
	return s, ok
	// return ctx.Value(key).(*domain.Session)
}
