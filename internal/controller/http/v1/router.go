package v1

import (
	"celebut-api/internal/controller/http/v1/handlers/auth"
	"celebut-api/internal/controller/http/v1/handlers/business"
	"celebut-api/internal/controller/http/v1/middleware"
	"celebut-api/internal/usecase"
	repo "celebut-api/internal/usecase/repo/postgres"
	"celebut-api/pkg/postgres"

	"celebut-api/pkg/logger"

	"github.com/gin-gonic/gin"

	// IMPORTANT: swagger docs
	_ "celebut-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"net/http"
)

// NewAppRouter -.
// Swagger spec:
// @title       Celebut API
// @description Celebut Backend REST endpoints
// @version     1.0
// @host        localhost:8083
// @BasePath    /v1
func NewAppRouter(handler *gin.Engine, l logger.Interface, pg *postgres.Postgres) {

	// Use cases
	clientUseCase := usecase.NewClientUseCase(repo.NewClientRepo(pg))
	industryUseCase := usecase.NewIndustryUseCase(repo.NewIndustryRepo(pg))

	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	//handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	//handler.GET("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("Celebut API"))
	//})

	// Routers

	client := handler.Group("/v1")
	{
		auth.NewClientAuthRoute(client, clientUseCase, l)
	}

	routes := handler.Group("/v1")
	{
		business.NewIndustryRoutes(routes, industryUseCase, l)

	}

	routes.Use(middleware.Authorization(clientUseCase))
}
