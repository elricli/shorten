package recovery

import (
	"github.com/drrrMikado/shorten/pkg/middleware"
	"log"
	"net/http"
	"runtime/debug"
)

func Recovery() middleware.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					log.Println(string(debug.Stack()))
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()
			h.ServeHTTP(w, r)
			return
		})
	}
}
