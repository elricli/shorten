package server

import (
	"context"
	"net"
	"net/http"
	"strings"
	"time"

	_ "net/http/pprof"

	"github.com/drrrMikado/shorten/internal/service"
	"github.com/drrrMikado/shorten/pkg/log"
	"github.com/drrrMikado/shorten/public/static"
)

type Server struct {
	*http.Server
	svc *service.Service
	opt option
}

const (
	_defaultAddr = ":8080"
)

func NewServer(svc *service.Service, opts ...Option) (*Server, func()) {
	opt := option{
		network: "tcp",
		address: _defaultAddr,
	}
	for _, o := range opts {
		o(&opt)
	}
	s := &Server{
		svc: svc,
		opt: opt,
	}
	s.initHandler()
	return s, s.stop
}

func (s *Server) Listen() {
	s.start()
	log.Infof("Server listening on %s...", s.opt.address)
	// pprof
	go func() {
		log.Info(http.ListenAndServe(":6060", nil))
	}()
}

func (s *Server) stop() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Info("Server forced to shutdown:", err)
	}
}

func (s *Server) start() {
	go func() {
		lis, err := net.Listen(s.opt.network, s.opt.address)
		if err != nil {
			log.Fatal(err)
		}
		if err := s.Serve(lis); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
}

func (s *Server) initHandler() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(static.FS))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		switch path {
		case "/":
			b, _ := static.FS.ReadFile("html/index.html")
			_, _ = w.Write(b)
		case "/api/shorten":
			s.shorten(w, r)
		default:
			shortUrl, err := s.svc.ShortUrl.Redirect(r.Context(), strings.Trim(path, "/"))
			if err != nil || shortUrl.URL == "" {
				http.Error(w, ErrLinkNotExist.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, shortUrl.URL, http.StatusMovedPermanently)
		}
		return
	})

	s.Server = &http.Server{
		Handler: mux,
	}
	if s.opt.middleware != nil {
		s.Handler = s.opt.middleware(mux)
	}
	return
}
