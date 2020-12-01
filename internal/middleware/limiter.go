package middleware

import (
	"github.com/drrrMikado/shorten/internal/rate"
	"net/http"
)

func Limiter(r float64, burstSize int) Middleware {
	l := rate.NewLimiter(r, burstSize)
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !l.Allow() {
				http.Error(w,
					http.StatusText(http.StatusForbidden),
					http.StatusForbidden)
				return
			}
		})
	}

}
