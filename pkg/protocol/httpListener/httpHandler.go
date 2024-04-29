package httpListener

import (
	"net/http"
	"projectsphere/cats-social/pkg/middleware/auth"
	"projectsphere/cats-social/pkg/middleware/logger"
	"projectsphere/cats-social/pkg/protocol/msg"
	"projectsphere/cats-social/pkg/utils/config"

	"github.com/gin-gonic/gin"
)

type HttpHandlerImpl struct {
}

func NewHttpHandler() *HttpHandlerImpl {
	return &HttpHandlerImpl{}
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

	// basePath.POST("/login", h.UserLogin)

	AddUserRouter(
		basePath,
	)

	return server
}

func AddUserRouter(
	r *gin.RouterGroup,

) {
	user := r.Group("/user")

	user.Use(auth.JwtAuthUserMiddleware())
	{

	}
}
