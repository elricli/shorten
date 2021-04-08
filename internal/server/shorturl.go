package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/drrrMikado/shorten/pkg/validator"
)

var (
	// ErrLinkNotExist .
	ErrLinkNotExist = errors.New("sorry, the link you accessed doesn't exist on our service. Please keep in mind that our shortened links are case sensitive and may contain letters and numbers")
)

func (s *Server) shorten(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_ = r.ParseForm()
	rawurl := r.Form.Get("url")
	if !validator.IsURL(rawurl) {
		_ = errResp(w, errors.New(rawurl+" is not valid url"))
		return
	} else if u, err := url.Parse(rawurl); err != nil {
		_ = errResp(w, err)
		return
	} else if u.Scheme == "" {
		u.Scheme = "http"
		rawurl = u.String()
	}
	expireParam := r.Form.Get("expire")
	var expire time.Time
	if expireParam != "" {
		expireSec, err := strconv.Atoi(expireParam)
		if err != nil {
			_ = errResp(w, err)
			return
		}
		expire = time.Now().Add(time.Duration(expireSec) * time.Second)
	}
	shortUrl, err := s.svc.ShortUrl.Shorten(ctx, rawurl, expire)
	if err != nil {
		_ = errResp(w, err)
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"errcode": 0,
		"errmsg":  "",
		"data":    shortUrl.Key,
	})
	return
}

func errResp(w http.ResponseWriter, err error) error {
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"errcode": 1,
		"errmsg":  err.Error(),
	})
}
