package signup

import (
	"api/internal/models"
	"api/src/tests"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_cannot_signup_if_password_is_not_provided_test(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	databaseConnection.AutoMigrate(&models.User{})

	user := map[string]string{
		"email": "test123@test.com",
	}

	body, _ := json.Marshal(user)

	requestRecorder, _ := tests.PerformUnAuthenticatedRequest(http.MethodPost, "/api/signup", body)

	assert.Equal(t, http.StatusUnprocessableEntity, requestRecorder.Code)
}
