package v1

import (
	"context"
	"log"

	"backend.app/configs"
	"backend.app/docs"
	"backend.app/internal/controller/http/v1/handlers"
	"backend.app/internal/controller/http/v1/handlers/auth"
	"backend.app/internal/controller/http/v1/handlers/business"
	postsroutes "backend.app/internal/controller/http/v1/handlers/posts"
	"backend.app/internal/controller/http/v1/handlers/relationships"
	"backend.app/internal/mappers"
	"backend.app/internal/middleware"
	"backend.app/internal/repo/postgres/accounts"
	"backend.app/internal/repo/postgres/posts"
	relationships3 "backend.app/internal/repo/postgres/relationships"
	"backend.app/internal/services/file"
	"backend.app/internal/services/mailer"
	"backend.app/internal/services/otp_generator"
	postsservice "backend.app/internal/services/posts"
	relationships2 "backend.app/internal/services/relationships"
	"backend.app/internal/services/token"
	"backend.app/internal/services/user_otp"
	"backend.app/internal/services/users"
	"backend.app/internal/usecase"
	"backend.app/internal/usecase/otp"
	"backend.app/pkg/logger"
	"backend.app/pkg/postgres"
	"github.com/aws/aws-sdk-go-v2/config"
	credentials "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/gin-gonic/gin"

	// IMPORTANT: swagger docs
	_ "backend.app/docs"

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
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description `Bearer token` for authenticated requests
// @securityDefinitions.apikey Auth-Token
// @in header
// @name x-auth-token
// @description Type token for unauthenticated requests
func NewAppRouter(handler *gin.Engine, l logger.Interface, pg *postgres.Postgres, cfg *configs.Config) *gin.RouterGroup {

	docs.SwaggerInfo.Title = "Celebut API"
	docs.SwaggerInfo.Description = "Celebut Backend REST endpoints"
	docs.SwaggerInfo.Version = "1.0"

	if cfg.Env == "local" {
		docs.SwaggerInfo.Host = "localhost:8083"
	} else {
		docs.SwaggerInfo.Host = "staging-api.celebut.com"
	}

	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	creds := credentials.NewStaticCredentialsProvider(cfg.AWS.AccessKey, cfg.AWS.Secret, "")
	awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithCredentialsProvider(creds), config.WithRegion(cfg.AWS.Region))
	if err != nil {
		log.Printf("error: %v", err)

		panic(err)
	}

	awsS3Client := file.NewS3Client(awsCfg, cfg.AWS.Region)

	// repo
	industryRepo := accounts.NewIndustryRepo(pg)
	clientRepo := accounts.NewClientRepo(pg)
	userRepo := accounts.NewUserRepo(pg)
	userOTPRepo := accounts.NewUserOTPRepo(pg)
	postsRepo := posts.NewPostsRepo(pg)
	postMediaRepo := posts.NewPostMediaRepo(pg)
	relationshipsRepo := relationships3.NewRelationshipsRepo(pg)
	reactionsRepo := posts.NewUserReactionRepo(pg)

	var otpGeneratorService otp_generator.Generator

	if cfg.Env == "production" {
		otpGeneratorService = otp_generator.NewOTPGeneratorService(cfg.OTP.Secret)
	} else {
		otpGeneratorService = otp_generator.NewLocalOTPGeneratorService()
	}

	// mailer
	mailerService := mailer.NewMailerService()
	userOtpService := user_otp.NewUserOTPService(userOTPRepo, otpGeneratorService)

	// Use cases
	clientUseCase := usecase.NewClientUseCase(clientRepo)
	industryUseCase := usecase.NewIndustryUseCase(industryRepo)
	userUseCase := usecase.NewUserUseCase(userRepo, userOTPRepo, otpGeneratorService)
	postsUseCase := usecase.NewPostUseCase(postsRepo, postMediaRepo, reactionsRepo, userRepo)
	otpUseCase := otp.NewUserOTPUseCase(userOtpService, mailerService)
	relationshipUseCase := usecase.NewRelationshipUseCase(relationshipsRepo)
	reactionsUseCase := usecase.NewUserReactionUseCase(reactionsRepo)

	// Mappers
	userMapper := &mappers.DtoUserMapper{}
	userReactionMapper := &mappers.DtoUserReactionMapper{}
	postsMapper := mappers.NewDtoPostMapper(userReactionMapper)
	industriesMapper := &mappers.DtoIndustryMapper{}

	// Services
	userService := users.NewUserService(userRepo, awsS3Client, userUseCase, userMapper)
	tokenService := token.NewTokenService(cfg.Token.Secret)
	postsService := postsservice.NewPostService(postsUseCase, reactionsUseCase, relationshipUseCase, awsS3Client, postsMapper)
	followerService := relationships2.NewFollowerService(relationshipUseCase, userUseCase, userMapper)
	lovedOneService := relationships2.NewLovedOneService(relationshipUseCase, userUseCase, userMapper)

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
		auth.NewRegisterRoutes(routes, userUseCase, l, tokenService, userMapper, mailerService, awsS3Client)
		auth.NewLoginRoute(routes, userUseCase, l, userMapper, tokenService)
		business.NewIndustryRoutes(routes, industryUseCase, l, industriesMapper)
		postsroutes.NewPostsRoute(routes, userService, postsService, l, postsMapper)
		postsroutes.NewFeedRoute(routes, userService, postsService, l, postsMapper)
		postsroutes.NewCommentsRoute(routes, userService, postsService, l, postsMapper)
		postsroutes.NewReactionsRoute(routes, userService, postsService, l, postsMapper)
		handlers.NewOTPRoute(routes, userUseCase, otpUseCase, l)
		handlers.NewUserRoutes(routes, userService, l, userMapper)
		relationships.NewCustomersRoute(routes, followerService, l)
		relationships.NewLovedOnesRoute(routes, lovedOneService, l)
	}

	return routes
}
