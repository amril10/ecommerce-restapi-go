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

type categoryHandler struct {
	service service.CategoryService
}

func NewCategoryHandler(s service.CategoryService) *categoryHandler {
	return &categoryHandler{
		service: s,
	}
}

func (h *categoryHandler) CreateCategory(c *gin.Context) {
	var category dto.CategoryRequest

	err := c.ShouldBindJSON(&category)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	result, err := h.service.CreateCategory(&category)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusCreated,
		Message:    "Category created successfully",
		Data:       result,
	})

	c.JSON(http.StatusCreated, res)
}

func (h *categoryHandler) GetAllCategory(c *gin.Context) {
	categories, err := h.service.GetAllCategory()
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusOK,
		Message:    "List categories",
		Data:       categories,
	})

	c.JSON(http.StatusOK, res)
}

func (h *categoryHandler) GetCategoryBySlug(c *gin.Context) {
	slugParam := c.Param("slug")

	categories, err := h.service.GetCategoryBySlug(slugParam)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusOK,
		Message:    "Category by slug",
		Data:       categories,
	})

	c.JSON(http.StatusOK, res)
}

func (h *categoryHandler) UpdateCategory(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid ID"})
		return
	}

	var req dto.CategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	result, err := h.service.UpdateCategory(uint(id), &req)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusOK,
		Message:    "Category updated successfully",
		Data:       result,
	})

	c.JSON(http.StatusOK, res)
}

func (h *categoryHandler) DeleteCategory(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid request id"})
		return
	}

	err = h.service.DeleteCategory(uint(id))
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseWithParams{
		StatusCode: http.StatusOK,
		Message:    "Category deleted successfully",
	})

	c.JSON(http.StatusOK, res)
}
