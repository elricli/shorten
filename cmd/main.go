package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/drrrMikado/shorten/internal/bloomfilter"
	"github.com/drrrMikado/shorten/internal/config"
	"github.com/drrrMikado/shorten/internal/database"
	"github.com/drrrMikado/shorten/internal/shorten"
)

var (
	staticPath string
)

func init() {
	flag.StringVar(&staticPath, "static", "content/static", "static file path")
	log.SetFlags(log.Ldate | log.Lshortfile | log.Lmicroseconds)
}

func main() {
	flag.Parse()
	ctx := context.Background()
	cfg, err := config.Init()
	if err != nil {
		log.Fatalln(err)
	}
	_ = cfg.Dump(os.Stdout)
	redisClient, err := database.NewRedisClient(ctx, cfg.Redis)
	if err != nil {
		log.Fatalf("unable to new redis client: %v\n", err)
	}
	bf, err := bloomfilter.New(cfg.BloomFilter.ExpectedInsertions, cfg.BloomFilter.FPP, cfg.BloomFilter.HashSeed)
	if err != nil {
		log.Fatalf("bloomfilter.New: %v\n", err)
	}
	scfg := shorten.ServerConfig{
		RedisClient: redisClient,
		BloomFilter: bf,
		StaticPath:  staticPath,
	}
	server := shorten.NewServer(cfg, scfg)
	handler := server.Install()
	go log.Fatalln(http.ListenAndServe(":80", handler))
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-ch
		log.Println("get a signal", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("exit")
			_ = redisClient.Close()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
