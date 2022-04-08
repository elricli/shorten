package server

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

var (
	// ErrLinkNotExist .
	ErrLinkNotExist = errors.New("sorry, the link you accessed doesn't exist on our service. Please keep in mind that our shortened links are case sensitive and may contain letters and numbers")
)

type response struct {
	Code    int
	Message string
	Detail  interface{}
}

func (s *Server) shorten(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_ = r.ParseForm()
	rawurl := r.Form.Get("url")
	if err := s.opt.validator.Var(rawurl, "required,url"); err != nil {
		errResp(w, err)
		return
	}
	expireParam := r.Form.Get("expire")
	if err := s.opt.validator.Var(expireParam, "gte=0"); err != nil {
		errResp(w, err)
		return
	}
	var expire time.Time
	if expireSec, _ := strconv.Atoi(expireParam); expireSec > 0 {
		expire = time.Now().Add(time.Duration(expireSec) * time.Second)
	}
	alias, err := s.svc.Shorten(ctx, rawurl, expire)
	if err != nil {
		errResp(w, err)
		return
	}
	successResp(w, map[string]string{"Key": alias.Key})
}

func errResp(w http.ResponseWriter, err error) {
	_ = json.NewEncoder(w).Encode(response{
		Code:    1,
		Message: err.Error(),
	})
}

func successResp(w http.ResponseWriter, details interface{}) {
	_ = json.NewEncoder(w).Encode(response{
		Code:    0,
		Message: "",
		Detail:  details,
	})
}
