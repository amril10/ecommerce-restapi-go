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

type productHandler struct {
	service service.ProductService
}

func NewProductHandler(s service.ProductService) *productHandler {
	return &productHandler{
		service: s,
	}
}

func (h *productHandler) GetAllProduct(c *gin.Context) {
	products, err := h.service.GetAllProduct()
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusOK,
		Message:    "List product",
		Data:       products,
	})

	c.JSON(http.StatusOK, res)
}

func (h *productHandler) GetProductBySlug(c *gin.Context) {
	slugParam := c.Param("slug")

	product, err := h.service.GetProductBySlug(slugParam)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusOK,
		Message:    "Product data by slug",
		Data:       product,
	})

	c.JSON(http.StatusOK, res)
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var products dto.ProductRequest
	err := c.ShouldBindJSON(&products)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	result, err := h.service.CreateProduct(&products)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusCreated,
		Message:    "Product created successfully",
		Data:       result,
	})

	c.JSON(http.StatusCreated, res)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid ID"})
		return
	}

	var req dto.ProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	result, err := h.service.UpdateProduct(id, &req)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusOK,
		Message:    "Updated product successfully",
		Data:       result,
	})

	c.JSON(http.StatusOK, res)
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid ID"})
		return
	}

	err = h.service.DeleteProduct(id)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusOK,
		Message:    "Deleted product successfully",
	})

	c.JSON(http.StatusOK, res)
}
