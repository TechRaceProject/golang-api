package login

import (
	"api/internal/models"
	"api/src/tests"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_cannot_login_if_user_doesnt_exist_test(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	databaseConnection.AutoMigrate(&models.User{})

	var users []models.User
	databaseConnection.Find(&users)
	assert.Empty(t, users, "Expected the users table to be empty")

	body := []byte(`{"email":"test@example.com","password":"password"}`)

	requestRecorder, _ := tests.PerformUnAuthenticatedRequest(http.MethodPost, "/api/login", body)

	assert.Equal(t, http.StatusUnprocessableEntity, requestRecorder.Code)
}
