package logger

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/google/wire"
)

func Provider() (log.Logger, func(), error) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "caller", log.DefaultCaller)

	return logger, func() {}, nil
}

var (
	// ProviderProductionSet wire set
	ProviderProductionSet = wire.NewSet(Provider)
)
