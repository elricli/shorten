package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/drrrMikado/shorten/internal/repo"
	"github.com/drrrMikado/shorten/internal/server"
	"github.com/drrrMikado/shorten/pkg/middleware"
	"github.com/drrrMikado/shorten/pkg/middleware/limiter"
	"github.com/drrrMikado/shorten/pkg/middleware/recovery"
	"github.com/drrrMikado/shorten/pkg/middleware/request"
	"github.com/drrrMikado/shorten/pkg/rate"
	"go.uber.org/zap"
)

func main() {
	flag.Parse()
	var logger *zap.SugaredLogger
	if l, err := zap.NewProduction(); err != nil {
		panic(err)
	} else {
		logger = l.Sugar().Named("main")
	}
	repoCfg := repo.Config{
		Dialect: "postgres",
		DSN:     os.Getenv("SHORTEN_DSN"),
	}
	srv, cf, err := Init(logger, repoCfg,
		server.Network("tcp"),
		server.Address(os.Getenv("SHORTEN_ADDR")),
		server.Middleware(middleware.Chain(
			request.Accept(http.MethodGet, http.MethodPost),
			limiter.Limiter(rate.NewLimiter(100, 100)),
			recovery.Recovery(recovery.WithLogger(logger)),
		)),
	)
	if err != nil {
		panic(err)
	}
	srv.Listen()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	logger.Info("Server shutting down by", s.String())
	cf()
	logger.Info("Server exiting")
}
