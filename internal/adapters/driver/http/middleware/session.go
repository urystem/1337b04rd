package middleware

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
)

func (s *session) CheckOrSetSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		createNew := func() {
			ses := s.ser.NewSession(ctx)
			if ses == nil {
				slog.Error("ses is nil")

				// errPage := &domain.ErrorPageData{
				// 	Code:    http.StatusInternalServerError,
				// 	Message: "cannot create a new session",
				// }

				// s.errhand.RenderError(w, errPage)
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
			ctx = context.WithValue(ctx, s.sessKey, ses)
			next.ServeHTTP(w, r.WithContext(ctx))
		}

		cookie, err := r.Cookie(s.cookieName)
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				createNew()
			} else {
				http.Error(w, "server error", http.StatusInternalServerError)
				// errPage := &domain.ErrorPageData{
				// 	Code:    http.StatusInternalServerError,
				// 	Message: "cannot get a cookie from brower",
				// }
				// s.errhand.RenderError(w, errPage)
			}
			return
		}
		ses := s.ser.GetSession(ctx, cookie.Value)
		if ses != nil {
			ctx = context.WithValue(ctx, s.sessKey, ses)
			next.ServeHTTP(w, r.WithContext(ctx))

		} else {
			// Кука есть, но сессия невалидна — создаём новую
			createNew()
		}
	})
}
