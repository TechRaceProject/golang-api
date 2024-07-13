package signup

import (
	"api/src/models"
	"api/src/tests"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_cannot_signup_if_password_is_not_provided_test(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	router := tests.GetTestRouter()

	databaseConnection.AutoMigrate(&models.User{})

	user := map[string]string{
		"email": "test@test.com",
	}
	body, _ := json.Marshal(user)

	request, _ := http.NewRequest(http.MethodPost, "/api/signup", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	requestRecorder := httptest.NewRecorder()
	router.ServeHTTP(requestRecorder, request)

	assert.Equal(t, http.StatusUnprocessableEntity, requestRecorder.Code)
}
