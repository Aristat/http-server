package logger

import (
	"os"

	"github.com/google/wire"
	"github.com/sirupsen/logrus"
)

// TODO: use logger interface
func Provider() (*logrus.Logger, func(), error) {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	return log, func() {}, nil
}

var (
	// ProviderProductionSet wire set
	ProviderProductionSet = wire.NewSet(Provider)
)
