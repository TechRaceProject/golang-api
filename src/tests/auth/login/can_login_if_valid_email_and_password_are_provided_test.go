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

func Test_can_login_if_valid_email_and_password_are_provided(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	err := databaseConnection.AutoMigrate(
		&models.User{},
		&models.Vehicle{},
		&models.Race{},
		&models.VehicleState{},
		&models.PrimaryLedColor{},
		&models.BuzzerVariable{},
		&models.HeadAngle{},
		&models.VehicleBattery{},
	)

	if err != nil {
		t.Error("An error occurred while migrating the database: ", err)
	}

	hashedPassword, _ := services.HashPassword("password")

	username := "username"

	user := models.User{
		Email:    "test@example.com",
		Password: string(hashedPassword),
		Username: &username,
	}

	databaseConnection.Create(&user)

	body, _ := json.Marshal(map[string]string{
		"email":    "test@example.com",
		"password": "password",
	})

	requestRecorder, _ := tests.PerformUnAuthenticatedRequest(http.MethodPost, "/api/login", body)

	assert.Equal(t, http.StatusOK, requestRecorder.Code)

	databaseConnection.Unscoped().Delete(&user)
}
