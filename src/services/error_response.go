package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
	Data interface{} `json:"data,omitempty"`
}

// SetErrorResponse sets a JSON error response with HTTP status code
func SetErrorResponse(c *gin.Context, statusCode int, errors []string, data interface{}) {
	var errResponse ErrorResponse
	errResponse.Errors = make([]struct {
		Message string `json:"message"`
	}, len(errors))

	for i, errMsg := range errors {
		errResponse.Errors[i].Message = errMsg
	}

	errResponse.Data = data

	c.JSON(statusCode, errResponse)
}

// SetUnauthorized sets 401 Unauthorized response
func SetUnauthorized(c *gin.Context, message string) {
	SetErrorResponse(c, http.StatusUnauthorized, []string{message}, nil)
}

// SetNotFound sends a 404 Not Found response with a custom message
func SetNotFound(c *gin.Context, message string) {
	SetErrorResponse(c, http.StatusNotFound, []string{message}, nil)
}

// SetUnprocessableEntity sets 422 Unprocessable Entity response
func SetUnprocessableEntity(c *gin.Context, message string) {
	SetErrorResponse(c, http.StatusUnprocessableEntity, []string{message}, nil)
}

// SetInternalServerError sends a 500 Internal Server Error response with a custom message
func SetInternalServerError(c *gin.Context, message string) {
	SetErrorResponse(c, http.StatusInternalServerError, []string{message}, nil)
}

// SetNoContent sends a 204 No Content response
func SetNoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
