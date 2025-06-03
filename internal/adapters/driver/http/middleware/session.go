package middleware

import (
	"errors"
	"net/http"

	"1337b04rd/pkg/contextkeys"
)

func (s *session) CheckOrSetSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		createNew := func() {
			ses := s.ser.NewSession(ctx)
			if ses == nil {
				//
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:     s.cookieName,
				Value:    ses.Uuid.String(),
				Path:     "/",
				MaxAge:   int(s.ttl.Seconds()),
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteLaxMode,
			})
			ctx = contextkeys.NewContext(ctx, ses)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		cookie, err := r.Cookie(s.cookieName)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				createNew()
			} else {
				http.Error(w, "server error", http.StatusInternalServerError)
			}
			return
		}
		ses := s.ser.GetSession(ctx, cookie.Value)
		if ses != nil {
			ctx = contextkeys.NewContext(ctx, ses)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {
			// Кука есть, но сессия невалидна — создаём новую
			createNew()
		}
	})
}
