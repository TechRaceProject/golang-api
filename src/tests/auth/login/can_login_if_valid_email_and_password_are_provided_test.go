package login

import (
	"api/src/models"
	"api/src/services"
	"api/src/tests"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_can_login_if_valid_email_and_password_are_provided(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	router := tests.GetTestRouter()

	databaseConnection.AutoMigrate(&models.User{})

	hashedPassword, _ := services.HashPassword("password")

	user := models.User{
		Email:    "test@example.com",
		Password: string(hashedPassword),
		Username: "username",
	}

	databaseConnection.Create(&user)

	body, _ := json.Marshal(map[string]string{
		"email":    "test@example.com",
		"password": "password",
	})

	request, _ := http.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	requestRecorder := httptest.NewRecorder()
	router.ServeHTTP(requestRecorder, request)

	assert.Equal(t, http.StatusOK, requestRecorder.Code)

	databaseConnection.Unscoped().Delete(&user)
}
