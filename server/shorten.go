package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Shorten http handler..
func Shorten(w http.ResponseWriter, r *http.Request) {
	url := r.Form.Get("url")
	key, err := svc.Shorten(url)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"errcode": 1,
			"errmsg":  err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"errcode": 0,
		"errmsg":  "",
		"data":    key,
	})
	return
}

// Redirect http handler.
func Redirect(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	url, err := svc.Redirect(key)
	if err != nil || url == "" {
		w.Write([]byte(`Sorry, the link you accessed doesn't exist on our service. Please keep in mind that our shortened links are case sensitive and may contain letters and numbers.`))
		return
	}
	http.Redirect(w, r, url, http.StatusMovedPermanently)
	return
}
