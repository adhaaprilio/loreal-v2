package handler

import (
	"loreal/dto"
	"loreal/errorHandler"
	"loreal/helper"
	"loreal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest

	if err := c.ShouldBindJSON(&register); err != nil {
		errorHandler.HandleError(c, &errorHandler.BadRequestError{Message: err.Error()})
	}

	if err := h.service.Register(&register); err != nil {
		errorHandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register successfully, please login",
	})

	c.JSON(http.StatusCreated, res)
}

func (h *authHandler) Login(c *gin.Context) {
	var login dto.LoginRequest
	err := c.ShouldBindJSON(&login)

	if err != nil {
		errorHandler.HandleError(c, &errorHandler.BadRequestError{Message: err.Error()})
		return
	}

	result, err := h.service.Login(&login)
	if err != nil {
		errorHandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Berhasil Login",
		Data:       result,
	})

	c.JSON(http.StatusOK, res)

}
