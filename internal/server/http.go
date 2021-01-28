package server

import (
	"github.com/drrrMikado/shorten/pkg/rate"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/drrrMikado/shorten/internal/middleware"
	"github.com/drrrMikado/shorten/internal/service"
)

var (
	svc        *service.Service
	staticPath string
)

// HTTPServe server.
func HTTPServe(path string, s *service.Service) {
	svc = s
	staticPath = path
	mux := http.NewServeMux()
	l := rate.NewLimiter(1000, 1000)
	mw := middleware.Chain(
		middleware.AcceptRequests(http.MethodGet, http.MethodPost),
		middleware.Limiter(l),
	)
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, staticPath+"/img/favicon.ico")
	})
	mux.HandleFunc("/api/shorten", shorten)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(staticPath))))
	mux.HandleFunc("/", errorWrap(defaultHandler))
	go log.Fatalln(http.ListenAndServe(":8080", mw(mux)))
}

func errorWrap(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(string(debug.Stack()))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		if err := f(w, r); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
