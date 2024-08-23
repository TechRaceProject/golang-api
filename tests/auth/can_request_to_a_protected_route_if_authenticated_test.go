package auth

import (
	"api/src/tests"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_can_request_to_a_protected_route_if_authenticated_test(t *testing.T) {
	gin.SetMode(gin.TestMode)

	requestRecorder, _ := tests.PerformAuthenticatedRequest(http.MethodGet, "/api/protected", []byte(""))

	assert.Equal(t, http.StatusOK, requestRecorder.Code)
}
