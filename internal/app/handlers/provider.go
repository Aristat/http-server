package handlers

import (
	"github.com/go-kit/kit/log"
	"github.com/google/wire"
)

func Provider(logger log.Logger) (*ServerHandler, func(), error) {
	serverHandler := NewServerHandler(logger)
	return serverHandler, func() {}, nil
}

var (
	// ProviderProductionSet wire set
	ProviderProductionSet = wire.NewSet(Provider)
)
