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
	payload := new(entity.CatParam)
	err := c.ShouldBindJSON(payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, msg.BadRequest(err.Error()))
		return
	}

	// Check for null values in payload fields
	if containsNull(payload) {
		c.JSON(http.StatusBadRequest, msg.BadRequest("JSON payload contains null values"))
		return
	}

	// Call service to create cat
	resp, err := h.catSvc.Create(c.Request.Context(), *payload)
	if err != nil {
		respError := msg.UnwrapRespError(err)
		c.JSON(respError.Code, respError)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

// containsNull checks if any field in the CatParam struct is null
func containsNull(param *entity.CatParam) bool {
	if param == nil {
		return false
	}

	// Check each field for null
	if param.Name == "" || param.Race == "" || param.AgeInMonth <= 0 {
		return true
	}
	return false
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
