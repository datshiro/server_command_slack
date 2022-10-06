package server

import (
	"fmt"
	"net/http"
	"server/internal/infras/server/config"
	"server/internal/infras/server/handlers"
	"server/internal/infras/server/middlewares"

	appContext "server/internal/infras/server/context"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func NewServer(cfg config.Config) Server {
	return &server{
		cfg: cfg,
		gin: gin.Default(),
		ctx: appContext.NewContext(),
	}
}

type Server interface {
	ConfigMiddlewares()
	Start()
	RegisterHandlers()
}
type server struct {
	cfg config.Config
	ctx appContext.Context
	gin *gin.Engine
}

func (s *server) ConfigMiddlewares() {
	s.gin.Use(middlewares.CORSMiddleware())
	s.gin.Use(func(c *gin.Context) {
		s.ctx.GinInject(c)
		c.Next()
	})
}

func (s *server) Start() {
	s.gin.Run(fmt.Sprintf("%s:%s", s.cfg.Host, s.cfg.Port))
}

func (s *server) RegisterHandlers() {
	s.gin.GET("/health", alive)
	s.gin.GET("/live", alive)
	s.gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	handlers.RegisterHandler(s.gin, s.cfg.APIPath)
}
func alive(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
}
