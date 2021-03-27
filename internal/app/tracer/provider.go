package tracer

import (
	"github.com/go-kit/kit/log"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"

	"github.com/google/wire"
)

func Provider(logger log.Logger) (opentracing.Tracer, func(), error) {
	jaegerConfig := config.Configuration{
		Disabled:    false,
		ServiceName: "http-server",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           false,
			LocalAgentHostPort: "localhost:6831",
		},
	}

	// TODO: use logger interface instead jaeger.StdLogger
	tracer, closer, err := jaegerConfig.NewTracer(config.Logger(jaeger.StdLogger))

	cleanup := func() {
		if e := closer.Close(); e != nil {
			logger.Log(e.Error())
		}
	}

	return tracer, cleanup, err
}

var (
	// ProviderProductionSet wire set
	ProviderProductionSet = wire.NewSet(Provider)
)
