package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"

	"github.com/prometheus/client_golang/prometheus"

	"backend.app/configs"
	"backend.app/internal/app"
)

func init() {
	// Initialize OpenTelemetry
	if err := initTracing(); err != nil {
		log.Fatalf("Failed to initialize tracer: %v", err)
	}

	// Initialize Prometheus metrics
	initMetrics()
}

func main() {
	// Configuration
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(cfg)
}

func initTracing() error {
	ctx := context.Background()

	// OTLP exporter (sends to Jaeger/Tempo)
	exp, err := otlptrace.New(ctx,
		otlptrace.WithEndpoint(os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")), // e.g., "jaeger:4317"
		otlptrace.WithInsecure(),
	)
	if err != nil {
		return fmt.Errorf("creating OTLP exporter: %w", err)
	}

	// Tracer provider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("celebut-backend"),
			attribute.String("environment", os.Getenv("ENVIRONMENT")),
		)),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	// Graceful shutdown
	go func() {
		<-ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()
		if err := tp.Shutdown(ctx); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	return nil
}

func initMetrics() {
	// HTTP request duration histogram
	prometheus.MustRegister(httpRequestsDuration)
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(httpRequestSize)
	prometheus.MustRegister(httpResponseSize)
}

// Metrics definitions
var (
	httpRequestsDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_requests_duration_seconds",
			Help:    "Duration of HTTP requests.",
			Buckets: prometheus.ExponentialBuckets(0.001, 2, 12), // 1ms to ~2s
		},
		[]string{"method", "path", "status"},
	)
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "path", "status"},
	)
	httpRequestSize = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_request_size_bytes",
			Help: "Size of HTTP requests.",
		},
		[]string{"method", "path"},
	)
	httpResponseSize = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_response_size_bytes",
			Help: "Size of HTTP responses.",
		},
		[]string{"method", "path"},
	)
)
