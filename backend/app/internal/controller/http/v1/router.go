package v1

import (
	"fmt"

	"backend.app/configs"
	"backend.app/docs"
	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/pkg/middleware"

	"github.com/gin-gonic/gin"

	// IMPORTANT: swagger docs
	_ "backend.app/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Routes struct {
	Handler handlers.Operations
	Server  *gin.Engine
	MW      *middleware.Middleware
}

// NewAppRouter -.
func NewAppRouter(server *gin.Engine, handler handlers.Operations, mw *middleware.Middleware, cfg *configs.Config) Routes {
	// Options
	server.Use(gin.Logger())
	server.Use(gin.Recovery())

	docs.SwaggerInfo.Title = "Celebut API"
	docs.SwaggerInfo.Description = "Celebut Backend REST endpoints"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// Swagger
	url := ginSwagger.URL(fmt.Sprintf("%s/swagger/doc.json", cfg.AppHost))
	server.GET("/swagger/*any", func(c *gin.Context) {
		if c.Param("any") == "/" || c.Param("any") == "" {
			c.Redirect(http.StatusTemporaryRedirect, "/swagger/index.html")
		} else {
			ginSwagger.WrapHandler(swaggerFiles.Handler, url)(c)
		}
	})

	// K8s probe
	server.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	// handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Global rate limiting middleware - applied to all routes
	// Generous limits for general API usage
	server.Use(mw.RateLimiter.RateLimit(middleware.RateLimiterConfig{
		Requests: 100,                    // 100 requests
		Window:   time.Minute,            // per minute
		KeyFunc:  middleware.KeyFuncs.IP, // Limit by IP address
	}))

	return Routes{
		Handler: handler,
		Server:  server,
		MW:      mw,
	}
}

// RegisterMetrics registers the Prometheus metrics endpoint
func (ro *Routes) RegisterMetrics(server *gin.Engine) {
	server.GET("/metrics", gin.WrapH(promhttp.Handler()))
}

func (ro Routes) RegisterRoutes(server *gin.Engine, handler handlers.Operations) {
	version := server.Group("/v1")

	AuthRoutes(version, handler, ro.MW)
	WalletRoutes(version, handler, ro.MW)
	ProfileRoutes(version, handler, ro.MW)
	WebSocketRoutes(version, handler, ro.MW)
	SettingsRoutes(version, handler, ro.MW)
	BlockRoutes(version, handler, ro.MW)
	FollowerRoutes(version, handler, ro.MW)
	CelebrationRoutes(version, handler, ro.MW)
	NotificationRoutes(version, handler, ro.MW) // Add notification routes
	AIRoutes(version, handler, ro.MW)           // Add AI routes
	AdminRoutes(version, handler, ro.MW)
}

func CheckRoutes(r *gin.Engine) {
	r.GET("/v1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Secured API version 1.0",
		})
	})
}
