package server

import (
	"encoding/json"
	"errors"
	"github.com/drrrMikado/shorten/pkg/validator"
	"net/http"
	"net/url"
	"strings"
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
	shortUrl, err := s.Svc.ShortUrl.Shorten(ctx, rawurl)
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

func (s *Server) defaultHandler(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path
	switch path {
	case "/":
		http.ServeFile(w, r, s.staticPath+"/html/index.html")
	default:
		shortUrl, err := s.Svc.ShortUrl.Get(r.Context(), strings.Trim(path, "/"))
		if err != nil || shortUrl.LongUrl == "" {
			return ErrLinkNotExist
		}
		http.Redirect(w, r, shortUrl.LongUrl, http.StatusMovedPermanently)
	}
	return nil
}

func errResp(w http.ResponseWriter, err error) error {
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"errcode": 1,
		"errmsg":  err.Error(),
	})
}
