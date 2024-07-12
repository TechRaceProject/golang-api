package auth_test

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

func TestCanSignupIfEmailAndPasswordAreProvided(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	services.SetConnection(databaseConnection)

	// Créer le serveur de test
	router := tests.GetTestRouter()

	// Créer une requête de test et la table associée
	databaseConnection.AutoMigrate(&models.User{})

	user := map[string]string{
		"email":    "test@example.com",
		"password": "password",
	}
	body, _ := json.Marshal(user)

	request, _ := http.NewRequest(http.MethodPost, "/api/signup", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	// Enregistrer la réponse
	requestRecorder := httptest.NewRecorder()
	router.ServeHTTP(requestRecorder, request)

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
