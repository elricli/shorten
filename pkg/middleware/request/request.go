package request

import (
	"github.com/drrrMikado/shorten/pkg/middleware"
	"net/http"
)

const maxURILength = 5000

// Accept serves 405 (Method Not Allowed) for any method not on the
// given list and 414 (Method Request URI Too Long) for any URI that exceeds
// the maxURILength.
func Accept(methods ...string) middleware.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if len(r.URL.String()) >= maxURILength {
				http.Error(w, http.StatusText(http.StatusRequestURITooLong), http.StatusRequestURITooLong)
				return
			}
			for _, m := range methods {
				if r.Method == m {
					h.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		})
	}
}
