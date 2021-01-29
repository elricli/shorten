package server

import (
	"github.com/drrrMikado/shorten/pkg/middleware"
)

type Option func(option *option)

type option struct {
	staticPath string
	address    string
	middleware middleware.Middleware
}

func Address(addr string) Option {
	return func(o *option) {
		o.address = addr
	}
}

func StaticPath(path string) Option {
	return func(o *option) {
		o.staticPath = path
	}
}

func Middleware(m middleware.Middleware) Option {
	return func(o *option) {
		o.middleware = m
	}
}
