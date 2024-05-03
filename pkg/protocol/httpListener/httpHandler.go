package httpListener

import (
	"net/http"
	catHandler "projectsphere/cats-social/internal/cat/handler"
	matchHandler "projectsphere/cats-social/internal/match/handler"
	userHandler "projectsphere/cats-social/internal/user/handler"
	"projectsphere/cats-social/pkg/middleware/logger"
	"projectsphere/cats-social/pkg/protocol/msg"
	"projectsphere/cats-social/pkg/utils/config"

	"github.com/gin-gonic/gin"
)

type HttpHandlerImpl struct {
	userHandler  userHandler.UserHandler
	catHandler   catHandler.CatHandler
	matchHandler matchHandler.MatchHandler
}

func NewHttpHandler(
	userHandler userHandler.UserHandler,
	catHandler catHandler.CatHandler,
	matchHandler matchHandler.MatchHandler,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		userHandler:  userHandler,
		catHandler:   catHandler,
		matchHandler: matchHandler,
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
	basePath := server.Group(config.Get().Application.Group)

	AddUserRouter(
		basePath,
		h.userHandler,
		h.catHandler,
		h.matchHandler,
	)

	return server
}

func AddUserRouter(
	r *gin.RouterGroup,
	userHandler userHandler.UserHandler,
	catHandler catHandler.CatHandler,
	matchHandler matchHandler.MatchHandler,
) {
	user := r.Group("/user")
	user.POST("/register", userHandler.Register)

	cat := r.Group("cat") // Adjusted route group
	{
		cat.PUT("/:id", catHandler.Update) // PUT method for updating cat with ID
		cat.POST("", catHandler.Create)
		cat.POST("/match", matchHandler.Create)
	}

	// user.Use(auth.JwtAuthUserMiddleware())
	// {

	// }
}
