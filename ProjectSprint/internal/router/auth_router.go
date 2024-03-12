package router

import (
	"ProjectSprint/internal/config"
	"ProjectSprint/internal/handlers"
	"ProjectSprint/internal/repository"
	"ProjectSprint/internal/services"

	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	authRepository := repository.NewAuthRepository(config.DB)
	authService := services.NewAuthService(authRepository)
	authHandler := handlers.NewAuthHandler(authService)

	api.POST("/user/register", authHandler.Register)
	api.POST("/user/login", authHandler.Login)
}
