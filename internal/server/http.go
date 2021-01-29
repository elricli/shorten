package server

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/drrrMikado/shorten/internal/service"
)

type Server struct {
	*http.Server
	svc *service.Service
	opt option
}

func NewServer(svc *service.Service, opts ...Option) (*Server, func()) {
	opt := option{
		staticPath: "public/static",
		address:    ":8080",
	}
	for _, o := range opts {
		o(&opt)
	}
	s := &Server{
		svc: svc,
		opt: opt,
	}
	s.newRouter()
	return s, s.stop
}

func (s *Server) Serve() {
	s.start()
}

func (s *Server) stop() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown:", err)
	}
}

func (s *Server) start() {
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()
}

func (s *Server) newRouter() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/shorten", s.shorten)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(s.opt.staticPath))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		switch path {
		case "/favicon.ico":
			http.ServeFile(w, r, s.GetStatic("/img/favicon.ico"))
		case "/":
			http.ServeFile(w, r, s.GetStatic("/html/index.html"))
		case "/api/shorten":
			s.shorten(w, r)
		default:
			shortUrl, err := s.svc.ShortUrl.Get(r.Context(), strings.Trim(path, "/"))
			if err != nil || shortUrl.LongUrl == "" {
				http.Error(w, ErrLinkNotExist.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, shortUrl.LongUrl, http.StatusMovedPermanently)
		}
		return
	})

	s.Handler = mux
	if s.opt.middleware != nil {
		s.Handler = s.opt.middleware(mux)
	}
	return
}

func (s *Server) GetStatic(path string) string {
	return s.opt.staticPath + path
}
