package tracer

import (
	"context"
	"time"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"

	logexporter "github.com/shubhamjagdhane/simple-load-balancer/log-exporter"
	"github.com/shubhamjagdhane/simple-load-balancer/logger"
)

type tracerImpl struct {
	tp     trace.TracerProvider
	tr     trace.Tracer
	enable bool
}

var DefaultTracer *tracerImpl

func New(enable bool, projectID string, TracerName string, logger logger.ILogger) *tracerImpl {
	if !enable {
		lExporter := logexporter.New(logger)
		tp := sdktrace.NewTracerProvider(
			// For this example code we use sdktrace.AlwaysSample sampler to sample all traces.
			// In a production application, use sdktrace. ProbabilitySampler with a desired probability.
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithBatcher(lExporter),
		)
		otel.SetTracerProvider(tp)
		tr := otel.Tracer(TracerName)
		DefaultTracer = &tracerImpl{
			tp:     tp,
			tr:     tr,
			enable: enable,
		}
		return DefaultTracer
	}
	// Create Google Cloud Trace exporter to be able to retrieve the collected spans
	exporter, err := texporter.New(texporter.WithProjectID(projectID))
	if err != nil {
		logger.Fatalf("Error creating new exporter: %v", err)
	}

	lExporter := logexporter.New(logger)

	tp := sdktrace.NewTracerProvider(
		// For this example code we use sdktrace.AlwaysSample sampler to sample all traces.
		// In a production application, use sdktrace. ProbabilitySampler with a desired probability.
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithBatcher(lExporter),
	)
	otel.SetTracerProvider(tp)
	tr := otel.Tracer(TracerName)
	DefaultTracer = &tracerImpl{
		tp:     tp,
		tr:     tr,
		enable: enable,
	}
	return DefaultTracer
}

func Shutdown() {
	if DefaultTracer == nil || !DefaultTracer.enable {
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = DefaultTracer.tp.(*sdktrace.TracerProvider).Shutdown(ctx)
}

func StartSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	if DefaultTracer == nil {
		tp := trace.NewNoopTracerProvider()
		DefaultTracer = &tracerImpl{
			tp:     tp,
			tr:     tp.Tracer(""),
			enable: false,
		}
	}
	return DefaultTracer.tr.Start(ctx, name)
}

func SetSpanAttributes(span trace.Span, input map[string]string) {
	for key, value := range input {
		span.SetAttributes(attribute.String(key, value))
	}
}
