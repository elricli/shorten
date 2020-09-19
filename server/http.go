package server

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/drrrMikado/shorten/service"
	"github.com/gorilla/mux"
)

var (
	svc *service.Service
)

// Serve server.
func Serve(s *service.Service) {
	svc = s
	r := mux.NewRouter()
	r.HandleFunc("/", recoverWrap(errorHandler(Index)))
	r.HandleFunc("/shorten", recoverWrap(errorHandler(Shorten))).
		Methods(http.MethodPost, http.MethodGet)
	r.HandleFunc("/{key:[0-9a-zA-Z]{10}}", Redirect).
		Methods(http.MethodGet)
	r.HandleFunc("/shorten", ShortenAPI).
		Methods(http.MethodPost).
		PathPrefix("/api/")
	go log.Fatal(http.ListenAndServe(":80", r))
}

func recoverWrap(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(string(debug.Stack()))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		f(w, r)
	}
}

func errorHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
