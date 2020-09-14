package shorten

import (
	"log"
	"net/http"

	"github.com/drrrMikado/shorten/bloomfilter"
	"github.com/drrrMikado/shorten/fastrand"
)

var (
	keyLen       = 10
	expectedSize = 10000000
	fpp          = 0.00001
	bf           = bloomfilter.New(expectedSize, fpp)
)

// Shorten .
func Shorten(w http.ResponseWriter, r *http.Request) {
	url := r.PostForm.Get("url")
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte("{\"errcode\":0, \"errmessage\":\"\",\"data\":\"" + getUKey(url) + "\"}"))
}

func getUKey(url string) string {
	for {
		key := fastrand.String(uint32(keyLen))
		if !bf.MightContain([]byte(key)) {
			return key
		}
		log.Println("key might contain, key:", key)
	}
}
