package v1

import (
	"celebut-api/configs"
	"celebut-api/internal/controller/http/v1/handlers"
	"celebut-api/internal/controller/http/v1/handlers/auth"
	"celebut-api/internal/controller/http/v1/handlers/business"
	"celebut-api/internal/mappers"
	"celebut-api/internal/middleware"
	postgres2 "celebut-api/internal/repo/postgres"
	"celebut-api/internal/services"
	"celebut-api/internal/usecase"
	"celebut-api/pkg/logger"
	"celebut-api/pkg/postgres"
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
func NewAppRouter(handler *gin.Engine, l logger.Interface, pg *postgres.Postgres, cfg *configs.Config) {

	// repo
	industryRepo := postgres2.NewIndustryRepo(pg)
	clientRepo := postgres2.NewClientRepo(pg)
	userRepo := postgres2.NewUserRepo(pg)
	celebrationsRepo := postgres2.NewCelebrationRepo(pg)

	// Use cases
	clientUseCase := usecase.NewClientUseCase(clientRepo)
	industryUseCase := usecase.NewIndustryUseCase(industryRepo)
	userUseCase := usecase.NewUserUseCase(userRepo)
	celebrationsUseCase := usecase.NewCelebrationUseCase(celebrationsRepo)

	mapper := &mappers.DtoUserMapper{}
	celebrationMapper := &mappers.DtoCelebrationMapper{}
	tokenService := services.NewTokenService(cfg.Token.Secret)

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

	clientAuthMiddleware := middleware.ClientAuthorization(clientUseCase, l)

	authRoutes := handler.Group("/v1")
	authRoutes.Use(clientAuthMiddleware)
	{
		auth.NewRegisterRoute(authRoutes, userUseCase, l, tokenService, mapper)
		auth.NewLoginRoute(authRoutes, userUseCase, l, mapper, tokenService)
		business.NewIndustryRoutes(authRoutes, industryUseCase, l)
	}

	appRoutes := handler.Group("/v1")
	appRoutes.Use(clientAuthMiddleware)
	{
		handlers.NewCelebrationsRoute(authRoutes, celebrationsUseCase, l, celebrationMapper)
	}

}
