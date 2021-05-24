//+build wireinject

package main

import (
	"github.com/drrrMikado/shorten/internal/domain"
	"github.com/drrrMikado/shorten/internal/repo"
	"github.com/drrrMikado/shorten/internal/server"
	"github.com/drrrMikado/shorten/internal/service"
	"github.com/drrrMikado/shorten/pkg/snowflake"
	"github.com/google/wire"
	"go.uber.org/zap"
)

//go:generate wire
func Init(logger *zap.SugaredLogger, serviceIDWorker *snowflake.IDWorker, repoCfg repo.Config, opts ...server.Option) (*server.Server, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, repo.ProviderSet, domain.ProviderSet))
}
