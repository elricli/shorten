package domain

import (
	"github.com/drrrMikado/shorten/internal/domain/alias"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	alias.ProviderSet,
)
