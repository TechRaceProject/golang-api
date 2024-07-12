package login

import (
	"api/src/models"
	"api/src/services"
	"api/src/tests"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCannotLoginIfUserDoesntExist(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	services.SetConnection(databaseConnection)

	router := tests.GetTestRouter()

	databaseConnection.AutoMigrate(&models.User{})

	var users []models.User
	databaseConnection.Find(&users)
	assert.Empty(t, users, "Expected the users table to be empty")

	body := []byte(`{"email":"test@example.com","password":"password"}`)

	request, _ := http.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	requestRecorder := httptest.NewRecorder()
	router.ServeHTTP(requestRecorder, request)

	assert.Equal(t, http.StatusUnprocessableEntity, requestRecorder.Code)
}
