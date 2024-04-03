package handler

import (
	"backend/entity"
	"backend/errorHandler"
	"backend/service"
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
	var register entity.User

	if err := c.ShouldBindJSON(&register); err != nil {
		errorHandler.ConflictError(err.Error())
	}

	result, err := h.service.Register(&register)
	if err != nil {
		errorHandler.HandleError(c, err)
		return
	}
	res := &entity.ResponseParams{
		Message: "User Registered Succesfully",
		Data:    result,
	}

	c.JSON(http.StatusCreated, res)
}
