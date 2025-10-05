package router

import (
	"github.com/amril10/rest-api-go/config"
	"github.com/amril10/rest-api-go/handler"
	"github.com/amril10/rest-api-go/middleware"
	"github.com/amril10/rest-api-go/repository"
	"github.com/amril10/rest-api-go/service"
	"github.com/gin-gonic/gin"
)

func ProductRouter(api *gin.RouterGroup) {
	productRepository := repository.NewProductRepository(config.DB)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	auth := api.Group("/products")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/", productHandler.GetAllProduct)
		auth.GET("/:slug", productHandler.GetProductBySlug)
		auth.POST("/", productHandler.CreateProduct)
		auth.PUT("/:id", productHandler.UpdateProduct)
		auth.DELETE("/:id", productHandler.DeleteProduct)
	}
}
