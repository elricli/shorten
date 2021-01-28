package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
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
	server, cf, err := InitServer(staticPath)
	if err != nil {
		log.Fatalln(err)
	}
	server.Serve()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill, os.Interrupt)
	for {
		s := <-ch
		log.Println("get a signal", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Println("exit")
			cf()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
