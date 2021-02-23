package server

import (
	"github.com/drrrMikado/shorten/pkg/middleware"
)

type Option func(option *option)

type option struct {
	network    string
	address    string
	middleware middleware.Middleware
}

// Network with server network.
func Network(network string) Option {
	return func(o *option) {
		o.network = network
	}
}

// Address with server address.
func Address(addr string) Option {
	return func(o *option) {
		o.address = addr
	}
}

// Middleware with server middleware option.
func Middleware(m middleware.Middleware) Option {
	return func(o *option) {
		o.middleware = m
	}
}
