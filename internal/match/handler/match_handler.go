package handler

import (
	"net/http"
	"projectsphere/cats-social/internal/match/entity"
	"projectsphere/cats-social/internal/match/service"
	"projectsphere/cats-social/pkg/protocol/msg"

	"github.com/gin-gonic/gin"
)

type MatchHandler struct {
	matchSvc service.MatchService
}

func NewMatchHandler(matchSvc service.MatchService) MatchHandler {
	return MatchHandler{
		matchSvc: matchSvc,
	}
}

func (h MatchHandler) Create(c *gin.Context) {
	payload := new(entity.MatchCat)

	err := c.ShouldBindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, msg.BadRequest(err.Error()))
		return
	}

	resp, err := h.matchSvc.Create(c.Request.Context(), *payload)
	if err != nil {
		respError := msg.UnwrapRespError(err)
		c.JSON(respError.Code, respError)
		return
	}

	c.JSON(http.StatusCreated, resp)
}
