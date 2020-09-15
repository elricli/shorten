package server

import (
	"net/http"

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
	r.HandleFunc("/shorten", Shorten).Methods(http.MethodPost)
	r.HandleFunc("/{key:[0-9a-zA-Z]{10}}", Redirect).Methods(http.MethodGet)
	go http.ListenAndServe(":8182", r)
}
