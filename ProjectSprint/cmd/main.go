package main

import (
	"ProjectSprint/internal/config"
	"ProjectSprint/internal/router"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := gin.Default()
	api := r.Group("/v1")
	router.AuthRouter(api)
	r.Run(fmt.Sprintf(":%v", 8000))
}
