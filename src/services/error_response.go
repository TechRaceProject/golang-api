package services

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

func SetValidationErrorResponse(c *gin.Context, err error) {
	errorMessages := ExtractValidationErrors(err)

	SetErrorResponse(c, http.StatusUnprocessableEntity, errorMessages, nil)
}

func SetJsonBindingErrorResponse(c *gin.Context, err error) {
	if err.Error() == "EOF" {
		SetUnprocessableEntity(c, "Invalid request body")

		return
	}

	SetValidationErrorResponse(c, err)
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

func ExtractValidationErrors(err error) []string {
	var validationErrors validator.ValidationErrors
	var errorMessages []string

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			errorMessage := fieldError.Namespace() + " " + fieldError.Tag()

			errorMessages = append(errorMessages, errorMessage)
		}

		return errorMessages
	}

	// If the error is not a validation error, return the error message as-is
	errorMessages = append(errorMessages, err.Error())

	return errorMessages
}
