package handler

import (
	"net/http"

	"github.com/amril10/rest-api-go/dto"
	"github.com/amril10/rest-api-go/errorhandler"
	"github.com/amril10/rest-api-go/helper"
	"github.com/amril10/rest-api-go/service"
	"github.com/gin-gonic/gin"
)

type profileHandler struct {
	service service.ProfileService
}

func NewProfileHandler(s service.ProfileService) *profileHandler {
	return &profileHandler{
		service: s,
	}
}

func (h *profileHandler) GetProfile(c *gin.Context) {
	userID := c.MustGet("userID").(int)

	profile, err := h.service.GetProfile(userID)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusOK,
		Message:    "Data profile",
		Data:       profile,
	})

	c.JSON(http.StatusOK, res)
}

func (h *profileHandler) UpdateProfile(c *gin.Context) {
	var profile dto.ProfileUpdateRequest
	if err := c.ShouldBindJSON(&profile); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	userID := c.MustGet("userID").(int)

	result, err := h.service.UpdateProfile(userID, &profile)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusOK,
		Message:    "User name updated successfully",
		Data:       result,
	})

	c.JSON(http.StatusOK, res)
}
