package router

import (
	"github.com/amril10/rest-api-go/config"
	"github.com/amril10/rest-api-go/handler"
	"github.com/amril10/rest-api-go/middleware"
	"github.com/amril10/rest-api-go/repository"
	"github.com/amril10/rest-api-go/service"
	"github.com/gin-gonic/gin"
)

func CategoryRouter(api *gin.RouterGroup) {
	categoryRepository := repository.NewCategoryRepository(config.DB)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	auth := api.Group("/categories")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/", categoryHandler.GetAllCategory)
		auth.GET("/:slug", categoryHandler.GetCategoryBySlug)
		auth.POST("/", categoryHandler.CreateCategory)
		auth.PUT("/:id", categoryHandler.UpdateCategory)
		auth.DELETE("/:id", categoryHandler.DeleteCategory)
	}
}
