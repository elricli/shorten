//+build wireinject

package main

import (
	"github.com/drrrMikado/shorten/internal/repo"
	"github.com/drrrMikado/shorten/internal/server"
	"github.com/drrrMikado/shorten/internal/service"
	"github.com/drrrMikado/shorten/pkg/rate"
	"github.com/google/wire"
)

//go:generate wire
func InitServer(staticPath string, limiter *rate.Limiter) (*server.Server, func(), error) {
	panic(wire.Build(server.NewServer, service.Set, repo.Set))
}
