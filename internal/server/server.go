package server

import (
	"context"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/drrrMikado/shorten/internal/service"
	"github.com/drrrMikado/shorten/public/static"
	"github.com/google/wire"
	"go.uber.org/zap"

	_ "net/http/pprof"
)

var ProviderSet = wire.NewSet(New)

type Server struct {
	*http.Server
	svc *service.Service
	opt option
	log *zap.SugaredLogger
}

const (
	_defaultAddr = ":8080"
)

func New(svc *service.Service, logger *zap.SugaredLogger, opts ...Option) (*Server, func()) {
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
		log: logger.Named("server"),
	}
	s.initHandler()
	return s, s.stop
}

func (s *Server) Listen() {
	s.start()
	s.log.Infof("Server listening on %s...", s.opt.address)
	// pprof
	go func() {
		s.log.Info(http.ListenAndServe(":6060", nil))
	}()
}

func (s *Server) stop() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		s.log.Info("Server forced to shutdown:", err)
	}
}

func (s *Server) start() {
	go func() {
		lis, err := net.Listen(s.opt.network, s.opt.address)
		if err != nil {
			s.log.Fatal(err)
		}
		if err := s.Serve(lis); err != nil && err != http.ErrServerClosed {
			s.log.Fatal(err)
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
			longurl, err := s.svc.Redirect(r.Context(), strings.Trim(path, "/"))
			if err != nil || longurl == "" {
				http.Error(w, ErrLinkNotExist.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, longurl, http.StatusMovedPermanently)
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
