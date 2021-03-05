package main

import (
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/drrrMikado/shorten/internal/di"
	"github.com/drrrMikado/shorten/internal/server"
	"github.com/drrrMikado/shorten/pkg/log"
	"github.com/drrrMikado/shorten/pkg/middleware"
	"github.com/drrrMikado/shorten/pkg/middleware/limiter"
	"github.com/drrrMikado/shorten/pkg/middleware/recovery"
	"github.com/drrrMikado/shorten/pkg/middleware/request"
	"github.com/drrrMikado/shorten/pkg/rate"
)

func main() {
	flag.Parse()
	srv, cf, err := di.InitServer(
		server.Network("tcp"),
		server.Address(os.Getenv("SHORTEN_ADDR")),
		server.Middleware(middleware.Chain(
			request.Accept(http.MethodGet, http.MethodPost),
			limiter.Limiter(rate.NewLimiter(100, 100)),
			recovery.Recovery(),
		)),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	srv.Listen()
	log.Info("Server listening...")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	s := <-ch
	log.Info("Server shutting down by", s.String())
	cf()
	log.Info("Server exiting")
}
