package server

import (
	"github.com/drrrMikado/shorten/pkg/middleware"
	"github.com/go-playground/validator/v10"
)

type Option func(option *option)

type option struct {
	network    string
	address    string
	middleware middleware.Middleware
	validator  *validator.Validate
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
		if addr == "" {
			addr = _defaultAddr
		}
		o.address = addr
	}
}

// Validator with validator.
func Validator(v *validator.Validate) Option {
	return func(o *option) {
		if v != nil {
			o.validator = v
		}
	}
}

// Middleware with server middleware option.
func Middleware(m middleware.Middleware) Option {
	return func(o *option) {
		o.middleware = m
	}
}
