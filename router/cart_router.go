package router

import (
	"github.com/amril10/rest-api-go/config"
	"github.com/amril10/rest-api-go/handler"
	"github.com/amril10/rest-api-go/middleware"
	"github.com/amril10/rest-api-go/repository"
	"github.com/amril10/rest-api-go/service"
	"github.com/gin-gonic/gin"
)

func CartRouter(api *gin.RouterGroup) {
	cartRepository := repository.NewCartRepository(config.DB)
	cartService := service.NewCartService(cartRepository)
	cartHandler := handler.NewCartHandler(cartService)

	auth := api.Group("/carts")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/", cartHandler.GetAllCart)
		auth.POST("/", cartHandler.CreateOrUpdateCart)
		auth.PATCH("/:id", cartHandler.UpdateCart)
		auth.DELETE("/:id", cartHandler.DeleteCart)
	}
}
