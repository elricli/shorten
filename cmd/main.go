package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/drrrMikado/shorten"
	"github.com/drrrMikado/shorten/conf"
	"github.com/drrrMikado/shorten/database/redis"
	"github.com/gorilla/mux"
)

var (
	config string
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile | log.Lmicroseconds)
	flag.StringVar(&config, "config", "conf/conf.yml", "config file path")
	flag.Parse()
}

func main() {
	c := conf.Init(config)
	redis.Init(c.Redis)
	client := redis.Client
	if err := client.Ping(context.Background()).Err(); err != nil {
		log.Fatalln(err)
	}
	if err := client.Set(context.Background(), "thisisformgoclient", "123", 60*time.Minute).Err(); err != nil {
		log.Fatalln(err)
	}
	// fmt.Println(hash.Murmur3_32([]byte("www.baidu.com")))
	// fmt.Println(hash.Murmur3_32([]byte("www.baidu.com")))
	// ------
	// http.HandleFunc("/shorten", shorten.Shorten)
	r := mux.NewRouter()
	r.HandleFunc("/shorten", shorten.Shorten)
	r.HandleFunc("/{key}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://www.bing.com", http.StatusMovedPermanently)
	})
	http.ListenAndServe(":8182", r)
}
