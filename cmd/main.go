package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/drrrMikado/shorten/conf"
	"github.com/drrrMikado/shorten/server"
	"github.com/drrrMikado/shorten/service"
)

var (
	config string
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile | log.Lmicroseconds)
	flag.StringVar(&config, "config", "conf/config.yml", "config file path")
}

func main() {
	flag.Parse()
	ctx := context.Background()
	c := conf.Init(config)
	svc := service.New(ctx, c)
	server.Serve(svc)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-ch
		log.Println("get a signal", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("exit")
			svc.Close()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
