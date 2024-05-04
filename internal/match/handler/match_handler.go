package handler

import (
	"fmt"
	"net/http"
	"projectsphere/cats-social/internal/match/entity"
	"projectsphere/cats-social/internal/match/service"
	"projectsphere/cats-social/pkg/middleware/auth"
	"projectsphere/cats-social/pkg/protocol/msg"
	"strconv"

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

func (h MatchHandler) Delete(c *gin.Context) {
	matchID := c.Param("id")
	userID, err := auth.GetUserIdInsideCtx(c)
	if err != nil {
		fmt.Println(err)
	}

	if matchID == "" {
		c.JSON(http.StatusBadRequest, msg.BadRequest("match ID is required"))
		return
	}

	id, err := strconv.Atoi(matchID)
	if err != nil {
		respError := msg.UnwrapRespError(msg.NotFound("id is not found"))
		c.JSON(respError.Code, respError)
		return
	}

	err = h.matchSvc.Delete(c.Request.Context(), id, int(userID))
	if err != nil {
		respError := msg.UnwrapRespError(err)
		c.JSON(respError.Code, respError)
		return
	}

	c.JSON(http.StatusOK, msg.ReturnResult("successfully delete match request", nil))
}
