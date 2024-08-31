package signup

import (
	"api/src/models"
	"api/src/tests"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_can_signup_if_email_and_password_are_provided_test(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Initialisez la connexion à la base de données de test
	databaseConnection := tests.GetTestDBConnection()

	// Effectuez les migrations pour toutes les tables nécessaires
	err := databaseConnection.AutoMigrate(
		&models.User{},
		&models.Vehicle{},
		&models.VehicleState{},
		&models.PrimaryLedColor{}, // Ajoutez toutes les tables nécessaires ici
		&models.SecondaryLedColor{},
		&models.BuzzerVariable{},
		&models.HeadAngle{},
		&models.VehicleBattery{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Préparez les données nécessaires pour le test
	databaseConnection.Create(&models.Vehicle{
		Name: "A vehicle is required to enable any user to signup",
	})

	// Données utilisateur pour le test
	user := map[string]string{
		"email":    "test@example.com",
		"password": "password",
	}
	body, _ := json.Marshal(user)

	// Effectuez une requête POST non authentifiée pour le signup
	requestRecorder, _ := tests.PerformUnAuthenticatedRequest(http.MethodPost, "/api/signup", body)

	// Vérifiez le statut de la réponse
	assert.Equal(t, http.StatusCreated, requestRecorder.Code, "Expected status code 201, got %d", requestRecorder.Code)

	// Vérifiez le contenu de la réponse
	var response map[string]interface{}
	err = json.Unmarshal(requestRecorder.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}
	assert.Equal(t, "User created", response["message"], "Expected message 'User created', got '%s'", response["message"])

	// Vérifiez que l'utilisateur a été créé dans la base de données
	var createdUser models.User
	result := databaseConnection.Where("email = ?", user["email"]).First(&createdUser)
	if result.Error != nil {
		t.Fatalf("Expected user with email %s to be created, but got error: %v", user["email"], result.Error)
	}
	assert.Equal(t, user["email"], createdUser.Email, "Expected email to be '%s', got '%s'", user["email"], createdUser.Email)

	// Nettoyez la base de données en supprimant l'utilisateur créé
	databaseConnection.Unscoped().Delete(&createdUser)
}
