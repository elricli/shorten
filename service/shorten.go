package service

import (
	"context"
	"errors"
	"log"

	"github.com/drrrMikado/shorten/fastrand"
)

var (
	redisHashKey = "shorten_hash_table_key"
)

// Shorten url.
func (s *Service) Shorten(url string) (key string, err error) {
	for {
		key = fastrand.String(10)
		if !s.bf.MightContain([]byte(key)) {
			break
		}
		log.Println("key might contain, key:", key)
	}
	if err = s.rdb.HSet(context.Background(), redisHashKey, key, url).Err(); err != nil {
		log.Println("hset err:", err)
		return
	}
	s.bf.Insert([]byte(key))
	return
}

// Redirect by key.
func (s *Service) Redirect(key string) (url string, err error) {
	if !s.bf.MightContain([]byte(key)) {
		return "", errors.New("key not exist")
	}
	url, err = s.rdb.HGet(context.Background(), redisHashKey, key).Result()
	if err != nil {
		log.Println("hget field:", key, " err:", err)
		return "", err
	}
	url = "http://" + url
	return
}
