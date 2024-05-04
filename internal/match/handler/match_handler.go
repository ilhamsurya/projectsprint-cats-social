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
	// Check if the header is missing
	if c.GetHeader("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, msg.Unauthorization("No authorization header provided"))
		return
	}

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

func (h MatchHandler) RejectMatchRequest(c *gin.Context) {
	// Check if the header is missing
	if c.GetHeader("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, msg.Unauthorization("No authorization header provided"))
		return
	}

	// Check if the request body is empty
	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, msg.BadRequest("Request body is empty"))
		return
	}

	// Parse JSON payload
	payload := new(entity.ProcessMatchRequest)
	err := c.ShouldBindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, msg.BadRequest(err.Error()))
		return
	}

	// Check for null values in payload fields
	if payload.MatchId == 0 {
		c.JSON(http.StatusBadRequest, msg.BadRequest("JSON payload contains null values"))
		return
	}

	userID, err := auth.GetUserIdInsideCtx(c)
	if err != nil {
		fmt.Println(err)
	}

	err = h.matchSvc.RejectMatchRequest(c.Request.Context(), *payload, int(userID))
	if err != nil {
		respError := msg.UnwrapRespError(err)
		c.JSON(respError.Code, respError)
		return
	}

	c.JSON(http.StatusOK, msg.ReturnResult("successfully reject the cat match request", nil))
}

func (h MatchHandler) ApproveMatchRequest(c *gin.Context) {
	// Check if the header is missing
	if c.GetHeader("Authorization") == "" {
		c.JSON(http.StatusUnauthorized, msg.Unauthorization("No authorization header provided"))
		return
	}

	// Check if the request body is empty
	if c.Request.Body == nil {
		c.JSON(http.StatusBadRequest, msg.BadRequest("Request body is empty"))
		return
	}

	// Parse JSON payload
	payload := new(entity.ProcessMatchRequest)
	err := c.ShouldBindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, msg.BadRequest(err.Error()))
		return
	}

	// Check for null values in payload fields
	if payload.MatchId == 0 {
		c.JSON(http.StatusBadRequest, msg.BadRequest("JSON payload contains null values"))
		return
	}

	userID, err := auth.GetUserIdInsideCtx(c)
	if err != nil {
		fmt.Println(err)
	}

	err = h.matchSvc.RejectMatchRequest(c.Request.Context(), *payload, int(userID))
	if err != nil {
		respError := msg.UnwrapRespError(err)
		c.JSON(respError.Code, respError)
		return
	}

	c.JSON(http.StatusOK, msg.ReturnResult("successfully approve the cat match request", nil))
}
