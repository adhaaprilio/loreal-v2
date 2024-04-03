package router

import (
	"backend/config"
	"backend/handler"
	"backend/repository"
	"backend/service"

	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	authRepository := repository.NewAuthRepository(config.ConnectDatabase())
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	api.POST("/user/register", authHandler.Register)
}
