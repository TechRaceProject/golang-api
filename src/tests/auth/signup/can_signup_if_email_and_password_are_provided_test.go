package signup

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

func Test_can_signup_if_email_and_password_are_provided_test(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	services.AutoMigrateModels(databaseConnection)

	databaseConnection.Create(&models.Vehicle{
		Name: "a vehicule is required to enable any user to signup",
	})

	user := map[string]string{
		"email":    "test@example.com",
		"password": "password",
	}
	body, _ := json.Marshal(user)

	// On effecture une requête unauthenticated
	requestRecorder, _ := tests.PerformUnAuthenticatedRequest(http.MethodPost, "/api/signup", body)

	// Vérifier le statut de la réponse
	assert.Equal(t, http.StatusCreated, requestRecorder.Code)

	// Vérifier le contenu de la réponse
	var response map[string]string
	json.Unmarshal(requestRecorder.Body.Bytes(), &response)
	assert.Equal(t, "User created", response["message"])

	// Vérifier que l'utilisateur a été créé dans la base de données
	var createdUser models.User
	databaseConnection.Where("email = ?", user["email"]).First(&createdUser)
	assert.Equal(t, user["email"], createdUser.Email)

	// On retire l'utilisateur de la base de données
	databaseConnection.Unscoped().Delete(&createdUser)
}
