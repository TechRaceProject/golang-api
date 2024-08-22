package race

import (
	"api/src/models"
	"api/src/tests"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_create_race_successfully(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	// Ensure the necessary migrations are run
	databaseConnection.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.Race{})
	// Create a mock user
	user := models.User{
		Email:    "testuser@example.com",
		Password: "securepassword",
	}
	databaseConnection.Create(&user)
	// Create a mock vehicle
	vehicle := models.Vehicle{
		Name: "Toyota",
	}
	databaseConnection.Create(&vehicle)

	// Define start and end times for the race
	startTime := time.Now()
	endTime := startTime.Add(time.Hour)

	// Prepare the JSON body for the POST request
	createBody, _ := json.Marshal(map[string]interface{}{
		"name":                 "testuser",
		"start_time":           startTime.Format(time.RFC3339),
		"end_time":             endTime.Format(time.RFC3339),
		"number_of_collisions": 5,
		"distance_travelled":   150,
		"average_speed":        130,
		"out_of_parcours":      1,
		"user_id":              user.ID,
		"vehicle_id":           vehicle.ID,
		"type":                 "manual",
		"status":               "not_started",
	})

	// Perform the authenticated request
	requestURL := fmt.Sprintf("/api/users/%d/races", user.ID) // Use /users/:userId/races
	requestRecorder, _ := tests.PerformAuthenticatedRequest(http.MethodPost, requestURL, createBody)

	assert.Equal(t, http.StatusCreated, requestRecorder.Code)
}
