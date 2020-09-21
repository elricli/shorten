package shorten

import (
	"encoding/json"
	"errors"
	"html/template"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/drrrMikado/shorten/internal/bloomfilter"
	"github.com/drrrMikado/shorten/internal/config"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var (
	// ErrLinkNotExist .
	ErrLinkNotExist = errors.New("sorry, the link you accessed doesn't exist on our service. Please keep in mind that our shortened links are case sensitive and may contain letters and numbers")
)

// Server can be installed to serve.
type Server struct {
	cfg         *config.Config
	redisClient *redis.Client
	bloomFilter *bloomfilter.BloomFilter
}

// ServerConfig contains everything needed by a server.
type ServerConfig struct {
	RedisClient *redis.Client
	BloomFilter *bloomfilter.BloomFilter
}

// TemplateData is template struct.
type TemplateData struct {
	URL    string
	ErrMsg string
	OriURL string
}

// NewServer creates as new Server with the given dependencies.
func NewServer(cfg *config.Config, scfg ServerConfig) *Server {
	return &Server{
		cfg:         cfg,
		redisClient: scfg.RedisClient,
		bloomFilter: scfg.BloomFilter,
	}
}

// Install registers server routes.
func (s *Server) Install() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/", s.errorWrap(s.index))
	r.HandleFunc("/shorten", s.errorWrap(s.shorten)).
		Methods(http.MethodPost, http.MethodGet)
	r.HandleFunc("/{key:[0-9a-zA-Z]{10}}", s.errorWrap(s.redirect)).
		Methods(http.MethodGet)
	r.HandleFunc("/shorten", s.shortenAPI).
		Methods(http.MethodPost).
		PathPrefix("/api/")
	r.PathPrefix("/static/").Handler(http.FileServer(http.Dir("content")))
	return r
}

func (s *Server) shortenAPI(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_ = r.ParseForm()
	url := r.Form.Get("url")
	key, err := Insert(ctx, url, s.redisClient, s.bloomFilter)
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

func (s *Server) shorten(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	_ = r.ParseForm()
	url := r.Form.Get("url")
	t := TemplateData{}
	key, err := Insert(ctx, url, s.redisClient, s.bloomFilter)
	if err != nil {
		t.ErrMsg = err.Error()
	}
	t.URL = key
	t.OriURL = url
	tmpl := template.Must(template.New("index").Parse(indexTemplate))
	return tmpl.Execute(w, t)
}

func (s *Server) redirect(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	key := mux.Vars(r)["key"]
	url, err := Get(ctx, key, s.redisClient, s.bloomFilter)
	if err != nil || url == "" {
		return ErrLinkNotExist
	}
	http.Redirect(w, r, url, http.StatusMovedPermanently)
	return nil
}

func (s *Server) index(w http.ResponseWriter, _ *http.Request) error {
	tmpl := template.Must(template.New("index").Parse(indexTemplate))
	return tmpl.Execute(w, TemplateData{})
}

func (s *Server) errorWrap(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(string(debug.Stack()))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		if err := f(w, r); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
