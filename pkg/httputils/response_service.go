package httputils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int         `json:"-"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

// SetCreated sets 201 Created response
func SetCreated(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{
		StatusCode: http.StatusCreated,
		Message:    message,
		Data:       data,
	})
}

// SetOK sets 200 OK response
func SetOK(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		StatusCode: http.StatusOK,
		Message:    message,
		Data:       data,
	})
}
