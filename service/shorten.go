package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"time"

	"github.com/drrrMikado/shorten/internal/fastrand"
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
	if s.rdb == nil {
		return "", errors.New("client is not available")
		// } else if s.bf == nil {
		// 	return "", errors.New("bf is not available")
	} else if !validator.IsURL(rawurl) {
		return "", errors.New(rawurl + " is not valid url")
	} else if u, err := url.Parse(rawurl); err != nil {
		return "", err
	} else if u.Scheme == "" {
		u.Scheme = "http"
		rawurl = u.String()
	}
	query := `SELECT key FROM shorten_url WHERE key = $1`
	var key string
	for {
		key = fastrand.String(7)
		err := s.db.QueryRowContext(ctx, query, key).Scan()
		if err != nil && err != sql.ErrNoRows {
			return "", err
		} else if err == sql.ErrNoRows {
			break
		}
		// if !s.bf.MightContain([]byte(key)) {
		// 	break
		// }
		log.Println("key might contain, key:", key)
	}
	if key == "" {
		return "", errors.New("key is empty")
	}
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
	insert := `INSERT INTO shorten_url VALUES($1, $2, $3, $4);`
	if _, err := s.db.ExecContext(
		ctx, insert, shorten.Key, shorten.ShortURL, shorten.LongURL, shorten.CreateAt); err != nil {
		return "", err
	}
	// go s.bf.Insert([]byte(key))
	return shorten.ShortURL, nil
}

// Get url by key.
func (s *Service) Get(ctx context.Context, key string) (v string, err error) {
	// if !s.bf.MightContain([]byte(key)) {
	// 	return "", errors.New("key not exist")
	// }
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
	query := `SELECT long_url FROM shorten_url WHERE key = $1`
	var longURL string
	if err := s.db.QueryRowContext(ctx, query, key).Scan(&longURL); err != nil {
		return "", nil
	}
	return longURL, nil
}
