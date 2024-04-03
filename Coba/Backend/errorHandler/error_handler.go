package errorHandler

import (
	"backend/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorHandler struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statuscode"`
}

func (e ErrorHandler) Error() string {
	return e.Message
}

func (e ErrorHandler) Status() int {
	return e.StatusCode
}

func BadRequestError(message string) ErrorHandler {
	return ErrorHandler{Message: message, StatusCode: 400}
}

func UnauthorizedError(message string) ErrorHandler {
	return ErrorHandler{Message: message, StatusCode: 401}
}

func ForbiddenError(message string) ErrorHandler {
	return ErrorHandler{Message: message, StatusCode: 403}
}

func NotFoundError(message string) ErrorHandler {
	return ErrorHandler{Message: message, StatusCode: 404}
}

func ConflictError(message string) ErrorHandler {
	return ErrorHandler{Message: message, StatusCode: 409}
}

func InternalServerError(message string) ErrorHandler {
	return ErrorHandler{Message: message, StatusCode: 500}
}

func HandleError(c *gin.Context, err error) {
	var statusCode int

	switch ErrorHandler.Status(ErrorHandler{}) {
	case 400:
		statusCode = http.StatusBadRequest
	case 404:
		statusCode = http.StatusNotFound
	case 409:
		statusCode = http.StatusConflict
	case 500:
		statusCode = http.StatusInternalServerError
	}

	response := helper.Response(helper.ResponseParams{
		Message:    ErrorHandler.Error(ErrorHandler{}),
		StatusCode: statusCode,
	})

	c.JSON(statusCode, response)
}
