package main

import (
	"fmt"
	"loreal/config"
	"loreal/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := gin.Default()
	r.Use(cors.Default())
	api := r.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.AuthRouter(api)

	r.Run(fmt.Sprintf(":%v", config.ENV.PORT)) // listen and serve on 0.0.0.0:8080
}
