package shorten

import (
	"encoding/json"
	"errors"
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
	staticPath  string
}

// ServerConfig contains everything needed by a server.
type ServerConfig struct {
	RedisClient *redis.Client
	BloomFilter *bloomfilter.BloomFilter
	StaticPath  string
}

// NewServer creates as new Server with the given dependencies.
func NewServer(cfg *config.Config, scfg ServerConfig) *Server {
	return &Server{
		cfg:         cfg,
		redisClient: scfg.RedisClient,
		bloomFilter: scfg.BloomFilter,
		staticPath:  scfg.StaticPath,
	}
}

// Install registers server routes.
func (s *Server) Install() http.Handler {
	r := mux.NewRouter()
	// API router
	r.HandleFunc("/{key:[0-9a-zA-Z]{7}}", s.errorWrap(s.redirect)).
		Methods(http.MethodGet)
	r.HandleFunc("/api/shorten", s.shorten).
		Methods(http.MethodPost)
	// frontend router
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, s.staticPath+"/html/index.html")
	})
	r.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, s.staticPath+"/img/favicon.ico")
	})
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(s.staticPath))))

	return r
}

func (s *Server) shorten(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	_ = r.ParseForm()
	url := r.Form.Get("url")
	key, err := Insert(ctx, url, s.cfg.Domain, s.redisClient, s.bloomFilter)
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
