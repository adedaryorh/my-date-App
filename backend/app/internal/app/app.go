package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"backend.app/configs"
	v1 "backend.app/internal/controller/http/v1"
	"backend.app/pkg/httpserver"
	"backend.app/pkg/logger"
	"backend.app/pkg/postgres"
	"github.com/gin-gonic/gin"
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

	// TODO: Use config
	handler.MaxMultipartMemory = 8 << 20

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
