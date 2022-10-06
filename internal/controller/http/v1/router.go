package v1

import (
	"celebut-api/configs"
	"celebut-api/docs"
	"celebut-api/internal/controller/http/v1/handlers"
	"celebut-api/internal/controller/http/v1/handlers/auth"
	"celebut-api/internal/controller/http/v1/handlers/business"
	"celebut-api/internal/mappers"
	"celebut-api/internal/middleware"
	"celebut-api/internal/repo/postgres/accounts"
	"celebut-api/internal/repo/postgres/celebrations"
	"celebut-api/internal/services/mailer"
	"celebut-api/internal/services/otp_generator"
	"celebut-api/internal/services/token"
	"celebut-api/internal/services/user_otp"
	"celebut-api/internal/usecase"
	"celebut-api/internal/usecase/otp"
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

	docs.SwaggerInfo.Title = "Celebut API"
	docs.SwaggerInfo.Description = "Celebut Backend REST endpoints"
	docs.SwaggerInfo.Version = "1.0"

	docs.SwaggerInfo.Host = "localhost:8083"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// mailer
	mailerService := mailer.NewMailerService()
	// repo
	industryRepo := accounts.NewIndustryRepo(pg)
	clientRepo := accounts.NewClientRepo(pg)
	userRepo := accounts.NewUserRepo(pg)
	userOTPRepo := accounts.NewUserOTPRepo(pg)

	celebrationsRepo := celebrations.NewCelebrationRepo(pg)

	var otpGeneratorService otp_generator.Generator

	if cfg.Env == "production" {
		otpGeneratorService = otp_generator.NewOTPGeneratorService(cfg.OTP.Secret)
	} else {
		otpGeneratorService = otp_generator.NewLocalOTPGeneratorService()
	}

	userOtpService := user_otp.NewUserOTPService(userOTPRepo, otpGeneratorService)

	// Use cases
	clientUseCase := usecase.NewClientUseCase(clientRepo)
	industryUseCase := usecase.NewIndustryUseCase(industryRepo)
	userUseCase := usecase.NewUserUseCase(userRepo, userOTPRepo, otpGeneratorService)
	celebrationsUseCase := usecase.NewCelebrationUseCase(celebrationsRepo)

	otpUseCase := otp.NewUserOTPUseCase(userOtpService, mailerService)

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
		handlers.NewOTPRoute(routes, userUseCase, otpUseCase, l)
	}

	return routes
}
