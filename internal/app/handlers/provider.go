package handlers

import (
	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

func Provider(logger *logrus.Logger) (*ServerHandler, func(), error) {
	serverHandler := NewServerHandler(logger)
	return serverHandler, func() {}, nil
}

var (
	// ProviderProductionSet wire set
	ProviderProductionSet = wire.NewSet(Provider)
)
