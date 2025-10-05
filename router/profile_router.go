package router

import (
	"github.com/amril10/rest-api-go/config"
	"github.com/amril10/rest-api-go/handler"
	"github.com/amril10/rest-api-go/middleware"
	"github.com/amril10/rest-api-go/repository"
	"github.com/amril10/rest-api-go/service"
	"github.com/gin-gonic/gin"
)

func ProfileRouter(api *gin.RouterGroup) {
	profileRepository := repository.NewProfileRepository(config.DB)
	profileService := service.NewProfileService(profileRepository)
	profileHandler := handler.NewProfileHandler(profileService)

	auth := api.Group("/profile")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/", profileHandler.GetProfile)
		auth.PATCH("/", profileHandler.UpdateProfile)
	}
}
