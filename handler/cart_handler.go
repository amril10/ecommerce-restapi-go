package handler

import (
	"net/http"
	"strconv"

	"github.com/amril10/rest-api-go/dto"
	"github.com/amril10/rest-api-go/errorhandler"
	"github.com/amril10/rest-api-go/helper"
	"github.com/amril10/rest-api-go/service"
	"github.com/gin-gonic/gin"
)

type cartHandler struct {
	service service.CartService
}

func NewCartHandler(s service.CartService) *cartHandler {
	return &cartHandler{
		service: s,
	}
}

func (h *cartHandler) GetAllCart(c *gin.Context) {
	carts, err := h.service.GetAllCart()
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusOK,
		Message:    "List of Cart",
		Data:       carts,
	})

	c.JSON(http.StatusOK, res)
}

func (h *cartHandler) CreateOrUpdateCart(c *gin.Context) {
	var cart dto.CartRequest
	err := c.ShouldBindJSON(&cart)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	userID := c.MustGet("userID").(int)

	result, err := h.service.CreateOrUpdateCart(&cart, userID)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	message := "Cart updated successfully"
	status := http.StatusOK
	if result.CreatedAt.Equal(result.UpdatedAt) {
		message = "Cart created successfully"
		status = http.StatusCreated
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: status,
		Message:    message,
		Data:       result,
	})

	c.JSON(status, res)
}

func (h *cartHandler) UpdateCart(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid ID"})
		return
	}

	var cart dto.UpdateCartRequest
	if err := c.ShouldBindJSON(&cart); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	userID := c.MustGet("userID").(int)

	result, err := h.service.UpdateCart(id, userID, &cart)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusOK,
		Message:    "cart updated successfully",
		Data:       result,
	})

	c.JSON(http.StatusOK, res)
}

func (h *cartHandler) DeleteCart(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid ID"})
		return
	}

	err = h.service.DeleteCart(id)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusOK,
		Message:    "Cart deleted successfully",
	})

	c.JSON(http.StatusOK, res)
}
