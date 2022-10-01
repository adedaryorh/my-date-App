package app

import (
	"celebut-api/configs"
	"celebut-api/internal/controller/http/v1"
	"celebut-api/pkg/httpserver"
	"celebut-api/pkg/logger"
	"celebut-api/pkg/postgres"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

// Run creates objects via constructors.
func Run(cfg *configs.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.NewClientRepo: %w", err))
	}
	defer pg.Close()

	initialiseRepositories(pg)

	// HTTP Server
	handler := gin.New()

	_ = v1.NewAppRouter(handler, l, pg, cfg)
	httpServer := httpserver.New(handler)

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
