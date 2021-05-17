package recovery

import (
	"net/http"

	"github.com/drrrMikado/shorten/pkg/log"
	"github.com/drrrMikado/shorten/pkg/middleware"
)

func Recovery() middleware.Middleware {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					log.Error("catch a error：", err)
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()
			h.ServeHTTP(w, r)
			return
		})
	}
}
