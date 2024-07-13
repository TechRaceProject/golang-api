package login

import (
	"api/src/models"
	"api/src/services"
	"api/src/tests"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_cannot_login_if_invalid_password_is_provided(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

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
		"password": "wrong_password",
	})

	requestRecorder, _ := tests.PerformUnAuthenticatedRequest(http.MethodPost, "/api/login", body)

	assert.Equal(t, http.StatusUnprocessableEntity, requestRecorder.Code)

	databaseConnection.Unscoped().Delete(&user)
}
