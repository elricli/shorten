package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

var (
	// ErrLinkNotExist .
	ErrLinkNotExist = errors.New("sorry, the link you accessed doesn't exist on our service. Please keep in mind that our shortened links are case sensitive and may contain letters and numbers")
)

func shorten(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_ = r.ParseForm()
	url := r.Form.Get("url")
	key, err := svc.Insert(ctx, url)
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"errcode": 1,
			"errmsg":  err.Error(),
		})
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"errcode": 0,
		"errmsg":  "",
		"data":    key,
	})
	return
}

func defaultHandler(w http.ResponseWriter, r *http.Request) error {
	path := r.URL.Path
	switch path {
	case "/":
		http.ServeFile(w, r, staticPath+"/html/index.html")
	default:
		url, err := svc.Get(r.Context(), strings.Trim(path, "/"))
		if err != nil || url == "" {
			return ErrLinkNotExist
		}
		http.Redirect(w, r, url, http.StatusMovedPermanently)
	}
	return nil
}
