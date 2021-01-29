package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/drrrMikado/shorten/internal/server"
	"github.com/drrrMikado/shorten/pkg/middleware"
	"github.com/drrrMikado/shorten/pkg/middleware/limiter"
	"github.com/drrrMikado/shorten/pkg/middleware/recovery"
	"github.com/drrrMikado/shorten/pkg/middleware/request"
	"github.com/drrrMikado/shorten/pkg/rate"
)

var (
	staticPath string
)

func init() {
	flag.StringVar(&staticPath, "static", "public/static", "static file path")
	log.SetFlags(log.Ldate | log.Lshortfile | log.Lmicroseconds)
}

func main() {
	flag.Parse()
	srv, cf, err := InitServer(
		server.Address(":8080"),
		server.StaticPath(staticPath),
		server.Middleware(middleware.Chain(
			request.Accept(http.MethodGet, http.MethodPost),
			limiter.Limiter(rate.NewLimiter(100, 100)),
			recovery.Recovery(),
		)),
	)
	if err != nil {
		log.Fatalln(err)
	}
	srv.Serve()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill, os.Interrupt)
	for {
		s := <-ch
		log.Println("get a signal", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("Server shutting down...")
			cf()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
