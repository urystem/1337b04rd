package middleware

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"1337b04rd/internal/ports/inbound"
	"1337b04rd/internal/ports/outbound"

	"github.com/google/uuid"
)

type session struct {
	cookieName string
	ttl        time.Duration
	redis      outbound.SessionRedisInter
	logger     *slog.Logger
}

func InitSession(name string, ttl time.Duration, redis outbound.SessionRedisInter, logger *slog.Logger) inbound.MiddleWareInter {
	return &session{name, ttl, redis, logger}
}

func (s *session) CheckOrSetSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		cookie, err := r.Cookie(s.cookieName)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				http.SetCookie(w, s.newSessionCookie(ctx))
				next.ServeHTTP(w, r)
				return
			}

			s.logger.Error("server", "error", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		has, err := s.redis.CheckSession(ctx, cookie.Value)
		if err != nil {
			s.logger.Error("server", "error", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		if !has {
			http.SetCookie(w, s.newSessionCookie(ctx))
			next.ServeHTTP(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *session) newSessionCookie(ctx context.Context) *http.Cookie {
	uuid := s.generateSessionID()
	s.redis.SetSession(ctx, uuid)
	return &http.Cookie{
		Name:     s.cookieName,
		Value:    uuid,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
}

func (s *session) generateSessionID() string {
	id := uuid.New() // Это UUIDv4 по умолчанию
	return id.String()
}
