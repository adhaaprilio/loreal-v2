package errorHandler

import (
	"ProjectSprint/internal/dto"
	"ProjectSprint/internal/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *Error400:
		statusCode = http.StatusBadRequest
	case *Error404:
		statusCode = http.StatusNotFound
	case *Error409:
		statusCode = http.StatusConflict
	case *Error500:
		statusCode = http.StatusInternalServerError
	}

	response := helper.Response(dto.ResponseParams{
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	c.JSON(statusCode, response)
}
