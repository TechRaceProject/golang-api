package auth

import (
	"api/src/tests"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_cannot_request_to_a_protected_route_if_not_authenticated(t *testing.T) {
	gin.SetMode(gin.TestMode)

	requestRecorder, _ := tests.PerformUnAuthenticatedRequest(http.MethodGet, "/api/protected", []byte(""))

	assert.Equal(t, http.StatusUnauthorized, requestRecorder.Code)
}
