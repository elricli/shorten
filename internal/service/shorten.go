package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/url"
	"time"

	"github.com/drrrMikado/shorten/internal/encode"
	"github.com/drrrMikado/shorten/internal/validator"
)

var (
	redisHashTableKey = "shorten_hash_table_key"
)

// Shorten url info.
type Shorten struct {
	Key      string    `json:"key,omitempty"`
	ShortURL string    `json:"short_url,omitempty"`
	LongURL  string    `json:"long_url,omitempty"`
	CreateAt time.Time `json:"create_at,omitempty"`
}

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
	shorten := &Shorten{
		Key:      key,
		ShortURL: s.c.Domain + "/" + key,
		LongURL:  rawurl,
		CreateAt: time.Now(),
	}
	b, err := json.Marshal(shorten)
	if err != nil {
		return "", err
	}
	if err = s.rdb.HSet(ctx, redisHashTableKey, key, b).Err(); err != nil {
		return "", err
	}
	if _, err := s.db.ExecContext(ctx,
		`INSERT INTO shorten_url VALUES($1, $2, $3, $4);`,
		shorten.Key, shorten.ShortURL, shorten.LongURL, shorten.CreateAt); err != nil {
		return "", err
	}
	return shorten.ShortURL, nil
}

// Get url by key.
func (s *Service) Get(ctx context.Context, key string) (longURL string, err error) {
	str, err := s.rdb.HGet(ctx, redisHashTableKey, key).Result()
	if err != nil {
		return
	}
	if str != "" {
		shorten := &Shorten{}
		if err = json.Unmarshal([]byte(str), shorten); err != nil {
			return
		}
		return shorten.LongURL, nil
	}
	if err = s.db.QueryRowContext(ctx,
		`SELECT long_url FROM shorten_url WHERE key = $1 LIMIT 1;`,
		key).Scan(&longURL); err != nil && err != sql.ErrNoRows {
		return
	}
	return longURL, nil
}
