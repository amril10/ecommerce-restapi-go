package main

import (
	"fmt"

	"github.com/amril10/rest-api-go/config"
	"github.com/amril10/rest-api-go/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.LoadDB()

	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.AuthRouter(api)
	router.ProfileRouter(api)
	router.CategoryRouter(api)
	router.ProductRouter(api)
	router.CartRouter(api)

	r.Run(fmt.Sprintf(":%v", config.ENV.PORT))
}
