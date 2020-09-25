package shorten

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"time"

	"github.com/drrrMikado/shorten/internal/bloomfilter"
	"github.com/drrrMikado/shorten/internal/fastrand"
	"github.com/drrrMikado/shorten/internal/validator"
	"github.com/go-redis/redis/v8"
)

var (
	redisHashTableKey = "shorten_hash_table_key"
)

// Shorten url info.
type Shorten struct {
	Key      string `json:"key,omitempty"`
	Link     string `json:"link,omitempty"`
	LongURL  string `json:"long_url,omitempty"`
	CreateAt int64  `json:"create_at,omitempty"`
}

// Insert url.
func Insert(ctx context.Context, rawurl, domain string, redisClient *redis.Client, bf *bloomfilter.BloomFilter) (string, error) {
	if redisClient == nil {
		return "", errors.New("client is not available")
	} else if bf == nil {
		return "", errors.New("bf is not available")
	} else if !validator.IsURL(rawurl) {
		return "", errors.New(rawurl + " is not valid url")
	} else if u, err := url.Parse(rawurl); err != nil {
		return "", err
	} else if u.Scheme == "" {
		u.Scheme = "http"
		rawurl = u.String()
	}
	var key string
	for {
		key = fastrand.String(7)
		if !bf.MightContain([]byte(key)) {
			break
		}
		log.Println("key might contain, key:", key)
	}
	if key == "" {
		return "", errors.New("key is empty")
	}
	shorten := &Shorten{
		Key:      key,
		Link:     domain + "/" + key,
		LongURL:  rawurl,
		CreateAt: time.Now().UnixNano(),
	}
	b, err := json.Marshal(shorten)
	if err != nil {
		return "", err
	}
	if err = redisClient.HSet(ctx, redisHashTableKey, key, b).Err(); err != nil {
		return "", err
	}
	go bf.Insert([]byte(key))
	return shorten.Link, nil
}

// Get url by key.
func Get(ctx context.Context, key string, redisClient *redis.Client, bf *bloomfilter.BloomFilter) (longurl string, err error) {
	if !bf.MightContain([]byte(key)) {
		return "", errors.New("key not exist")
	}
	str, err := redisClient.HGet(ctx, redisHashTableKey, key).Result()
	if err != nil {
		return
	}
	shorten := &Shorten{}
	if err = json.Unmarshal([]byte(str), shorten); err != nil {
		return
	}
	return shorten.LongURL, nil
}
