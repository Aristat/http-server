// +build wireinject

package router

import (
	"github.com/aristat/http-go-kit/internal/app/handlers"
	"github.com/aristat/http-go-kit/internal/app/logger"
	"github.com/aristat/http-go-kit/internal/app/tracer"
	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

// Build
func Build() (*chi.Mux, func(), error) {
	panic(wire.Build(
		ProviderProductionSet,
		handlers.ProviderProductionSet,
		tracer.ProviderProductionSet,
		logger.ProviderProductionSet,
	))
}
