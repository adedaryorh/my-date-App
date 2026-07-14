package app

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"backend.app/configs"
	"backend.app/database"
	v1 "backend.app/internal/controller/http/v1"
	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/pkg/httpserver"
	"backend.app/pkg/logger"
	"backend.app/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Run creates objects via constructors.
func Run(cfg *configs.Config) {
	l := logger.New(cfg.LogLevel)
	var err error

	if cfg.AppEnv != "dev" {
		runMigrations(cfg, l)
	}

	// Connect to DB
	db := database.ConnectDB(cfg, l)
	defer db.Postgres.Close()

	// HTTP Server
	server := gin.New()

	// TODO: Use config
	server.MaxMultipartMemory = 8 << 20

	// Add observability middleware (tracing + metrics)
	server.Use(middleware.ObservabilityMiddleware())

	handler := handlers.NewHandler(l, cfg, &db)
	mw := middleware.NewMiddleware(&db, cfg, l) // Create middleware instance
	routesWithServer := v1.NewAppRouter(server, handler, mw, cfg)
	routesWithServer.RegisterRoutes(server, handler)

	// Add health check endpoint
	routesWithServer.Server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Celebut Application",
		})
	})

	// Add metrics endpoint for Prometheus
	routesWithServer.Server.GET("/metrics", gin.WrapH(promhttp.Handler()))

	httpServer := httpserver.New(routesWithServer.Server, cfg.Port)

	// Waiting signal
	interrupt := make(chan os.Signal, 10)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
