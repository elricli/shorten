module github.com/drrrMikado/shorten

go 1.15

replace github.com/drrrMikado/shorten => ./

require (
	github.com/go-redis/redis/v8 v8.0.0-beta.12
	github.com/gorilla/mux v1.8.0
	gopkg.in/yaml.v2 v2.3.0
)
