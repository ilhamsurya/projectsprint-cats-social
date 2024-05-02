package handler

import (
	"fmt"
	"net/http"
	"projectsphere/cats-social/internal/cat/entity"
	"projectsphere/cats-social/internal/cat/service"
	"projectsphere/cats-social/pkg/protocol/msg"
	"strconv"

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

func (h CatHandler) Update(c *gin.Context) {
	payload := new(entity.CatParam)
	err := c.ShouldBindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, msg.BadRequest(err.Error()))
		return
	}
	catID := c.Param("id")

	if catID == "" {
		c.JSON(http.StatusBadRequest, msg.BadRequest("cat ID is required"))
		return
	}

	id, err := strconv.Atoi(catID)
	if err != nil {
		c.JSON(http.StatusNotFound, msg.BadRequest("id is not found"))
		return
	}

	resp, err := h.catSvc.Update(c.Request.Context(), id, *payload)
	fmt.Print(err)
	if err != nil {
		respError := msg.UnwrapRespError(err)
		c.JSON(respError.Code, respError)
		return
	}

	c.JSON(http.StatusOK, resp)
}
