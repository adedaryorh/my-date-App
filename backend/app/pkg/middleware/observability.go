package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/propagation"
)

// ObservabilityMiddleware adds tracing and metrics to HTTP requests
func ObservabilityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()
		if path == "" {
			path = c.Request.URL.Path
		}

		// Extract method
		method := c.Request.Method

		// Start span
		ctx, span := otel.Tracer("http-server").Start(
			c.Request.Context(),
			fmt.Sprintf("HTTP %s", method),
			otel.WithAttributes(
				semconv.HTTPMethodKey.String(method),
				semconv.HTTPURLKey.String(c.Request.URL.Path),
			),
		)
		defer span.End()

		// Propagate context to downstream services
		c.Request = c.Request.WithContext(ctx)

		// Process request
		c.Next()

		// Calculate duration
		duration := time.Since(start)
		status := c.Writer.Status()

		// Record metrics
		httpRequestsDuration.WithLabelValues(method, path, strconv.Itoa(status)).Observe(duration.Seconds())
		httpRequestsTotal.WithLabelValues(method, path, strconv.Itoa(status)).Inc()
		httpRequestSize.WithLabelValues(method, path).Observe(float64(c.Request.ContentLength))
		// Note: Response size requires custom response writer - simplified here

		// Record span attributes
		span.SetAttributes(
			attribute.Int("http.status_code", status),
			attribute.String("http.route", path),
		)

		if status >= 400 {
			span.SetStatus(codes.Error, fmt.Sprintf("HTTP %d", status))
		} else {
			span.SetStatus(codes.Ok, "")
		}
	}
}