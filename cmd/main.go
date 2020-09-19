package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/drrrMikado/shorten/internal/config"
	"github.com/drrrMikado/shorten/server"
	"github.com/drrrMikado/shorten/service"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile | log.Lmicroseconds)
}

func main() {
	ctx := context.Background()
	cfg, err := config.Init()
	if err != nil {
		log.Fatalln(err)
	}
	svc, err := service.New(ctx, cfg)
	if err != nil {
		log.Fatalln(err)
	}
	server.Serve(svc)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-ch
		log.Println("get a signal", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("exit")
			if err := svc.Close(); err != nil {
				log.Fatalln(err)
			}
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
