package v1

import (
	"fmt"

	"backend.app/configs"
	"backend.app/docs"
	"backend.app/internal/controller/http/v1/handlers"

	"github.com/gin-gonic/gin"

	// IMPORTANT: swagger docs
	_ "backend.app/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"net/http"
)

type Routes struct {
	Handler handlers.Operations
	Server  *gin.Engine
}

// NewAppRouter -.
func NewAppRouter(server *gin.Engine, handler handlers.Operations, cfg *configs.Config) Routes {
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
	//handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return Routes{Handler: handler, Server: server}
}

func (ro Routes) RegisterRoutes(server *gin.Engine, handler handlers.Operations) {
	version := server.Group("/v1")

	AuthRoutes(version, handler)
	WalletRoutes(version, handler)
	UserRoutes(version, handler)
}
func CheckRoutes(r *gin.Engine) {
	r.GET("/v1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Secured API version 1.0",
		})
	})
}
