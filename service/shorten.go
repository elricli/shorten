package service

import (
	"context"
	"errors"
	"log"
	"net/url"

	"github.com/drrrMikado/shorten/fastrand"
	"github.com/drrrMikado/shorten/validator"
)

var (
	redisHashKey = "shorten_hash_table_key"
)

// Shorten url.
func (s *Service) Shorten(rawurl string) (key string, err error) {
	if !validator.IsURL(rawurl) {
		err = errors.New(rawurl + " is not valid url")
		return
	}
	for {
		key = fastrand.String(10)
		if !s.bf.MightContain([]byte(key)) {
			break
		}
		log.Println("key might contain, key:", key)
	}
	if err = s.rdb.HSet(context.Background(), redisHashKey, key, rawurl).Err(); err != nil {
		log.Println("hset err:", err)
		return
	}
	s.bf.Insert([]byte(key))
	return
}

// Redirect get url by key.
func (s *Service) Redirect(key string) (v string, err error) {
	if !s.bf.MightContain([]byte(key)) {
		return "", errors.New("key not exist")
	}
	v, err = s.rdb.HGet(context.Background(), redisHashKey, key).Result()
	if err != nil {
		log.Println("hget field:", key, " err:", err)
		return "", err
	}
	var u *url.URL
	u, err = url.Parse(v)
	if err != nil {
		return "", err
	} else if u.Scheme == "" {
		u.Scheme = "http"
		v = u.String()
	}
	return
}
