package router

import (
	"loreal/config"
	"loreal/handler"
	"loreal/repository"
	"loreal/service"

	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	authRepository := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthSerivce(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	api.POST("/register", authHandler.Register)
	api.POST("/login", authHandler.Login)
}
