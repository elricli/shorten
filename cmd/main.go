package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/drrrMikado/shorten/internal/config"
	"github.com/drrrMikado/shorten/server"
	"github.com/drrrMikado/shorten/service"
	_ "github.com/lib/pq"
)

var (
	staticPath string
	configFile string
)

func init() {
	flag.StringVar(&staticPath, "static", "content/static", "static file path")
	flag.StringVar(&configFile, "config", "config.yml", "config file path")
	log.SetFlags(log.Ldate | log.Lshortfile | log.Lmicroseconds)
}

func main() {
	flag.Parse()
	ctx := context.Background()
	cfg, err := config.Init(configFile)
	if err != nil {
		log.Fatalln(err)
	}
	_ = cfg.Dump(os.Stdout)
	svc, err := service.New(ctx, cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer svc.Close()
	server.HTTPServe(staticPath, svc)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill, os.Interrupt)
	for {
		s := <-ch
		log.Println("get a signal", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("exit")
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
