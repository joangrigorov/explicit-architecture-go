package otel

import (
	"app/config/api"
	"app/internal/infrastructure/framework/observability/otel/middleware"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	b3prop "go.opentelemetry.io/contrib/propagators/b3"
	jaegerprop "go.opentelemetry.io/contrib/propagators/jaeger"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdkTrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
)

func AddOpenTelemetryMiddleware(
	cfg api.Config,
	engine *gin.Engine,
) {
	engine.Use(otelgin.Middleware(cfg.App.Name), middleware.RecordRequestData)
}

func NewTracerProvider(cfg api.Config) (*sdkTrace.TracerProvider, error) {
	headers := map[string]string{
		"content-type": "application/json",
	}

	exp, err := otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			otlptracehttp.WithEndpoint(cfg.Tracing.Endpoint),
			otlptracehttp.WithHeaders(headers),
			otlptracehttp.WithInsecure(),
		),
	)
	if err != nil {
		log.Fatalln(fmt.Errorf("OpenTelemetry exporter err:%w", err))
		return nil, err
	}

	tp := sdkTrace.NewTracerProvider(
		sdkTrace.WithBatcher(
			exp,
			sdkTrace.WithMaxExportBatchSize(sdkTrace.DefaultMaxExportBatchSize),
			sdkTrace.WithBatchTimeout(sdkTrace.DefaultScheduleDelay*time.Millisecond),
			sdkTrace.WithMaxExportBatchSize(sdkTrace.DefaultMaxExportBatchSize),
		),
		sdkTrace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String(cfg.App.Name),
			),
		),
	)

	return tp, nil
}

func RegisterTracer(lc fx.Lifecycle, tp *sdkTrace.TracerProvider) {
	composite := propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		b3prop.New(),
		jaegerprop.Jaeger{},
		propagation.Baggage{},
	)

	otel.SetTextMapPropagator(composite)

	otel.SetTracerProvider(tp)

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Println("Shutting down tracer provider")
			return tp.Shutdown(ctx)
		},
	})
}

func DefaultTracer(cfg api.Config) trace.Tracer {
	return otel.Tracer(cfg.App.Name)
}
