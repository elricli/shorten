package service

import (
	"context"
	"errors"
	"log"
	"net/url"

	"github.com/drrrMikado/shorten/internal/fastrand"
	"github.com/drrrMikado/shorten/internal/validator"
)

var (
	redisHashKey = "shorten_hash_table_key"
)

// Insert url.
func (s *Service) Insert(ctx context.Context, rawurl string) (key string, err error) {
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
	if err = s.rdb.HSet(ctx, redisHashKey, key, rawurl).Err(); err != nil {
		log.Println("HSet err:", err)
		return
	}
	go s.bf.Insert([]byte(key))
	return
}

// Get url by key.
func (s *Service) Get(ctx context.Context, key string) (v string, err error) {
	if !s.bf.MightContain([]byte(key)) {
		return "", errors.New("key not exist")
	}
	v, err = s.rdb.HGet(ctx, redisHashKey, key).Result()
	if err != nil {
		log.Println("HGet field:", key, " err:", err)
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
