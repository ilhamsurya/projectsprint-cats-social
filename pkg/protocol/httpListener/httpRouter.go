package httpListener

import (
	"github.com/gin-gonic/gin"
)

type HttpRouterImpl struct {
	handlers *HttpHandlerImpl
}

func NewHttpRoute(
	handlers *HttpHandlerImpl,
) *HttpRouterImpl {
	return &HttpRouterImpl{
		handlers: handlers,
	}
}

func (h *HttpRouterImpl) Router() *gin.Engine {
	server := h.handlers.Router()
	return server
}
