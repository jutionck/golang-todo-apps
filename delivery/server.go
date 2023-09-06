package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jutionck/golang-todo-apps/config"
	"github.com/jutionck/golang-todo-apps/delivery/controller"
	"github.com/jutionck/golang-todo-apps/delivery/middleware"
	"github.com/jutionck/golang-todo-apps/docs"
	"github.com/jutionck/golang-todo-apps/manager"
	"github.com/jutionck/golang-todo-apps/usecase"
	"github.com/jutionck/golang-todo-apps/utils/service"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	uc            manager.UseCaseManager
	authService   usecase.AuthenticationUseCase
	engine        *gin.Engine
	loggerService service.LoggerService
	jwtService    service.JwtService
	host          string
}

func (s *Server) setupControllers() {
	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	rg := s.engine.Group("/api/v1")
	controller.NewUserController(s.uc.UserUseCase(), rg, authMiddleware).Route()
	controller.NewTodoController(s.uc.TodoUseCase(), rg, authMiddleware).Route()
	controller.NewAuthController(rg, s.authService).Route()
	controller.NewInitController(rg, s.uc.UserUseCase()).Route()
}

func (s *Server) swagDocs() {
	docs.SwaggerInfo.Title = "Todo App"
	docs.SwaggerInfo.Version = "v1"
	docs.SwaggerInfo.BasePath = "/api/v1"
	s.engine.GET("/api/v1/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func (s *Server) Run() {
	s.engine.Use(middleware.NewLogMiddleware(s.loggerService).Logger())
	s.setupControllers()
	s.swagDocs()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running %s", err.Error()))
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	// manager
	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		panic(err)
	}
	ur := manager.NewRepoManager(infraManager)
	uUc := manager.NewUseCaseManager(ur)
	engine := gin.Default()
	loggerService := service.NewLoggerService(cfg.FileConfig)
	jwtService := service.NewJwtService(cfg.TokenConfig)
	authUseCase := usecase.NewAuthenticationUseCase(uUc.UserUseCase(), jwtService)
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		uc:            uUc,
		engine:        engine,
		loggerService: loggerService,
		jwtService:    jwtService,
		authService:   authUseCase,
		host:          host,
	}
}
