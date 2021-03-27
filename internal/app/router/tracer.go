package router

import (
	"net/http"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go"
)

func tracerMiddleware(tracer opentracing.Tracer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := nethttp.Middleware(
			tracer,
			next,
			nethttp.OperationNameFunc(func(r *http.Request) string {
				return "HTTP " + r.Method
			}))
		return fn
	}
}
