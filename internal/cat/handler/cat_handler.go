package handler

import (
	"net/http"
	"projectsphere/cats-social/internal/cat/entity"
	"projectsphere/cats-social/internal/cat/service"
	"projectsphere/cats-social/pkg/protocol/msg"

	"github.com/gin-gonic/gin"
)

type CatHandler struct {
	catSvc service.CatService
}

func NewCatHandler(catSvc service.CatService) CatHandler {
	return CatHandler{
		catSvc: catSvc,
	}
}

func (h CatHandler) Create(c *gin.Context) {
	payload := new(entity.CatParam)

	err := c.ShouldBindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, msg.BadRequest(err.Error()))
		return
	}

	resp, err := h.catSvc.Create(c.Request.Context(), *payload)
	if err != nil {
		respError := msg.UnwrapRespError(err)
		c.JSON(respError.Code, respError)
		return
	}

	c.JSON(http.StatusCreated, resp)
}
