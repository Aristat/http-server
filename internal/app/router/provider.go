package router

import (
	"github.com/aristat/http-go-kit/internal/app/api"
	"github.com/aristat/http-go-kit/internal/app/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-kit/kit/log"
	"github.com/google/wire"
	"github.com/opentracing/opentracing-go"
)

func Provider(logger log.Logger, tracer opentracing.Tracer, serverHandler *handlers.ServerHandler) (*chi.Mux, func(), error) {
	router := NewChiServer(logger, tracer)
	api.HandlerFromMux(serverHandler, router)

	return router, func() {}, nil
}

var (
	// ProviderProductionSet wire set
	ProviderProductionSet = wire.NewSet(Provider)
)
