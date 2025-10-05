package router

import (
	"github.com/amril10/rest-api-go/config"
	"github.com/amril10/rest-api-go/handler"
	"github.com/amril10/rest-api-go/repository"
	"github.com/amril10/rest-api-go/service"
	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	authRepository := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)


	api.POST("/auth/register", authHandler.Register)
	api.POST("/auth/login", authHandler.Login)
}