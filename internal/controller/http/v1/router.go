package v1

import (
	"celebut-api/configs"
	"celebut-api/internal/controller/http/v1/handlers"
	"celebut-api/internal/controller/http/v1/handlers/auth"
	"celebut-api/internal/controller/http/v1/handlers/business"
	"celebut-api/internal/mappers"
	"celebut-api/internal/middleware"
	"celebut-api/internal/repo/postgres/accounts"
	"celebut-api/internal/repo/postgres/celebrations"
	"celebut-api/internal/services/mailer"
	"celebut-api/internal/services/token"
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
func NewAppRouter(handler *gin.Engine, l logger.Interface, pg *postgres.Postgres, cfg *configs.Config) *gin.RouterGroup {

	// mailer
	mailerService := mailer.NewMailerService()
	// repo
	industryRepo := accounts.NewIndustryRepo(pg)
	clientRepo := accounts.NewClientRepo(pg)
	userRepo := accounts.NewUserRepo(pg)
	registerOTPRepo := accounts.NewRegisterOTPRepo(pg)

	celebrationsRepo := celebrations.NewCelebrationRepo(pg)

	otpService := token.NewOTPService(cfg.OTP.Secret)

	// Use cases
	clientUseCase := usecase.NewClientUseCase(clientRepo)
	industryUseCase := usecase.NewIndustryUseCase(industryRepo)
	userUseCase := usecase.NewUserUseCase(userRepo, registerOTPRepo, otpService)
	celebrationsUseCase := usecase.NewCelebrationUseCase(celebrationsRepo)

	mapper := &mappers.DtoUserMapper{}
	celebrationMapper := &mappers.DtoCelebrationMapper{}
	tokenService := token.NewTokenService(cfg.Token.Secret)

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

	routes := handler.Group("/v1")

	// Middleware
	routes.Use(
		middleware.ClientAuthorization(clientUseCase, l),
		middleware.Authorization(*userUseCase, tokenService, l))

	// Routes
	{
		auth.NewClientAuthRoute(routes, clientUseCase, l)
		auth.NewRegisterRoutes(routes, userUseCase, l, tokenService, mapper, mailerService)
		auth.NewLoginRoute(routes, userUseCase, l, mapper, tokenService)
		business.NewIndustryRoutes(routes, industryUseCase, l)
		handlers.NewCelebrationsRoute(routes, celebrationsUseCase, l, celebrationMapper)
	}

	return routes
}
