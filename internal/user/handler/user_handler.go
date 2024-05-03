package handler

import (
	"net/http"
	"projectsphere/cats-social/internal/user/entity"
	"projectsphere/cats-social/internal/user/service"
	"projectsphere/cats-social/pkg/protocol/msg"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userSvc service.UserService
}

func NewUserHandler(userSvc service.UserService) UserHandler {
	return UserHandler{
		userSvc: userSvc,
	}
}

func (h UserHandler) Register(c *gin.Context) {
	payload := new(entity.UserParam)

	err := c.ShouldBindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, msg.BadRequest(err.Error()))
		return
	}

	resp, err := h.userSvc.Register(c.Request.Context(), payload)
	if err != nil {
		respError := msg.UnwrapRespError(err)
		c.JSON(respError.Code, respError)
		return
	}

	c.JSON(http.StatusCreated, msg.ReturnResult("User registered successfully", resp))
}
