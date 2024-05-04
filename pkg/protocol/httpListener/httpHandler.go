package httpListener

import (
	"net/http"
	catHandler "projectsphere/cats-social/internal/cat/handler"
	matchHandler "projectsphere/cats-social/internal/match/handler"
	userHandler "projectsphere/cats-social/internal/user/handler"
	"projectsphere/cats-social/pkg/middleware/auth"
	"projectsphere/cats-social/pkg/middleware/logger"
	"projectsphere/cats-social/pkg/protocol/msg"
	"projectsphere/cats-social/pkg/utils/config"

	"github.com/gin-gonic/gin"
)

type HttpHandlerImpl struct {
	userHandler  userHandler.UserHandler
	catHandler   catHandler.CatHandler
	matchHandler matchHandler.MatchHandler
	jwtAuth      auth.JWTAuth
}

func NewHttpHandler(
	userHandler userHandler.UserHandler,
	catHandler catHandler.CatHandler,
	matchHandler matchHandler.MatchHandler,
	jwtAuth auth.JWTAuth,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		userHandler:  userHandler,
		catHandler:   catHandler,
		matchHandler: matchHandler,
		jwtAuth:      jwtAuth,
	}
}
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Disposition, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (h *HttpHandlerImpl) Router() *gin.Engine {
	server := gin.New()
	server.Use(gin.Recovery(), logger.Logger(), CORSMiddleware())
	server.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, msg.NotFound(msg.ErrPageNotFound))
	})

	server.Static("/v1/docs", "./dist")
	r := server.Group(config.Get().Application.Group)

	user := r.Group("/user")
	user.POST("/register", h.userHandler.Register)
	user.POST("/login", h.userHandler.Login)

	cat := r.Group("cat") // Adjusted route group
	cat.Use(h.jwtAuth.JwtAuthUserMiddleware())
	{
		cat.PUT("/:id", h.catHandler.Update) // PUT method for updating cat with ID
		cat.POST("", h.catHandler.Create)
		cat.POST("/match", h.matchHandler.Create)
		cat.GET("", h.catHandler.Get)
		cat.DELETE("/:id", h.catHandler.Delete)
	}

	return server
}
