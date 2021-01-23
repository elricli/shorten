package service

import (
	"context"
	"errors"
	"net/url"

	"github.com/drrrMikado/shorten/ent"
	"github.com/drrrMikado/shorten/ent/shorturl"
	"github.com/drrrMikado/shorten/internal/encode"
	"github.com/drrrMikado/shorten/internal/validator"
)

var (
	redisHashTableKey = "shorten_hash_table_key"
)

// Insert url.
func (s *Service) Insert(ctx context.Context, rawurl string) (string, error) {
	if !validator.IsURL(rawurl) {
		return "", errors.New(rawurl + " is not valid url")
	} else if u, err := url.Parse(rawurl); err != nil {
		return "", err
	} else if u.Scheme == "" {
		u.Scheme = "http"
		rawurl = u.String()
	}
	nextID, err := s.idWorker.NextID()
	if err != nil {
		return "", nil
	}
	key := encode.ToBase62(uint64(nextID))
	shortURL := s.c.Domain + "/" + key
	if err = s.rdb.HSet(ctx, redisHashTableKey, key, shortURL).Err(); err != nil {
		return "", err
	}
	// insert to db.
	if _, err = s.entCli.ShortUrl.Create().
		SetKey(key).
		SetShortURL(shortURL).
		SetLongURL(rawurl).
		Save(ctx); err != nil {
		return "", err
	}
	return shortURL, nil
}

// Get url by key.
func (s *Service) Get(ctx context.Context, key string) (string, error) {
	// get from cache
	if longURL, err := s.rdb.HGet(ctx, redisHashTableKey, key).Result(); err != nil {
		return "", err
	} else if longURL != "" {
		return longURL, nil
	}
	// from db
	if shortURL, err := s.entCli.ShortUrl.
		Query().
		Where(shorturl.KeyEQ(key)).
		First(ctx); err != nil && !ent.IsNotFound(err) {
		return "", err
	} else {
		return shortURL.LongURL, nil
	}
}
