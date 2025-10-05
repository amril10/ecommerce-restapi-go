package handler

import (
	"net/http"

	"github.com/amril10/rest-api-go/dto"
	"github.com/amril10/rest-api-go/errorhandler"
	"github.com/amril10/rest-api-go/helper"
	"github.com/amril10/rest-api-go/model"
	"github.com/amril10/rest-api-go/service"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest

	if err := c.ShouldBind(&register); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Register(&register); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusCreated,
		Message:    "Register successfully, please login",
	})

	c.JSON(http.StatusCreated, res)
}

func (h *authHandler) Login(c *gin.Context) {
	var login dto.LoginRequest

	err := c.ShouldBindJSON(&login)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	result, err := h.service.Login(&login)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	token, err := helper.GenerateToken(&model.User{ID: result.ID})
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.UnauthorizedError{Message: err.Error()})
		return
	}

	c.SetCookie(
		"token",
		token,
		3600,
		"/",
		"",
		false,
		true,
	)

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusOK,
		Message:    "Login successfully",
		Data:       result,
	})

	c.JSON(http.StatusOK, res)
}
