package handlers

import (
	"ProjectSprint/internal/dto"
	"ProjectSprint/internal/errorHandler"
	"ProjectSprint/internal/helper"
	"ProjectSprint/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service services.AuthService
}

func NewAuthHandler(s services.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		errorHandler.HandleError(c, &errorHandler.Error409{Message: err.Error()})
	}

	// if err := h.service.Register(&register); err != nil {
	// 	errorHandler.HandleError(c, err)
	// 	return
	// }
	result, err := h.service.Register(&register)
	if err != nil {
		errorHandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		// StatusCode: http.StatusCreated,
		Message: "User Registered Succesfully",
		Data:    result,
	})

	c.JSON(http.StatusCreated, res)
}

func (h *authHandler) Login(c *gin.Context) {
	var login dto.LoginRequest
	err := c.ShouldBindJSON(&login)

	if err != nil {
		errorHandler.HandleError(c, &errorHandler.Error400{Message: err.Error()})
		return
	}
	result, err := h.service.Login(&login)
	if err != nil {
		errorHandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		// StatusCode: http.StatusCreated,
		Message: "Login Succesfully",
		Data:    result,
	})

	c.JSON(http.StatusCreated, res)
}
