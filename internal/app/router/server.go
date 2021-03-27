package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
)

func NewChiServer(logger log.Logger, tracer opentracing.Tracer) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)

	r.Use(loggerMiddleware(logger))
	r.Use(tracerMiddleware(tracer))

	return r
}
