package limiter

import (
	"github.com/drrrMikado/shorten/pkg/middleware"
	"github.com/drrrMikado/shorten/pkg/rate"
	"net/http"
)

func Limiter(l *rate.Limiter) middleware.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !l.Allow() {
				http.Error(w,
					http.StatusText(http.StatusForbidden),
					http.StatusForbidden)
				return
			}
			h.ServeHTTP(w, r)
			return
		})
	}
}
