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
	"github.com/gin-gonic/gin"
)

// Run creates objects via constructors.
func Run(cfg *configs.Config) {
	l := logger.New(cfg.LogLevel)
	var err error
	// Connect to DB

	db := database.ConnectDB(cfg, l)
	defer db.Postgres.Close()

	// HTTP Server
	server := gin.New()

	// TODO: Use config
	server.MaxMultipartMemory = 8 << 20
	handler := handlers.NewHandler(l, cfg, &db)
	routesWithServer := v1.NewAppRouter(server, handler, cfg)
	routesWithServer.RegisterRoutes(server, handler)

	routesWithServer.Server.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Celebut Application",
		})
	})

	httpServer := httpserver.New(routesWithServer.Server, cfg.Port)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
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
