package server

import (
	"context"
	"github.com/drrrMikado/shorten/pkg/rate"
	"log"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/drrrMikado/shorten/internal/service"
	"github.com/drrrMikado/shorten/pkg/middleware"
)

type Server struct {
	srv        *http.Server
	Svc        *service.Service
	staticPath string
}

func NewServer(svc *service.Service, staticPath string) (*Server, func()) {
	s := &Server{
		Svc:        svc,
		staticPath: staticPath,
	}
	s.srv = &http.Server{
		Addr:    ":8080",
		Handler: s.initRouter(),
	}
	return s, s.stop
}

func (s *Server) Serve() {
	s.start()
}

func (s *Server) stop() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown:", err)
	}
}

func (s *Server) start() {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()
}

// HTTPServe server.
func (s *Server) initRouter() http.Handler {
	mux := http.NewServeMux()
	l := rate.NewLimiter(1000, 1000)
	mw := middleware.Chain(
		middleware.AcceptRequests(http.MethodGet, http.MethodPost),
		middleware.Limiter(l),
	)
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, s.staticPath+"/img/favicon.ico")
	})
	mux.HandleFunc("/api/shorten", s.shorten)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(s.staticPath))))
	mux.HandleFunc("/", errorWrap(s.defaultHandler))
	return mw(mux)
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
