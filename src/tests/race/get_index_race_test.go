package race

import (
	"api/src/models"
	"api/src/models/attributes"
	"api/src/tests"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type ApiResponse struct {
	Message string        `json:"message"`
	Data    []models.Race `json:"data"`
}

func Test_get_races_index(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	// Ensure the necessary migrations are run
	databaseConnection.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.Race{})

	// Create a mock user
	user := models.User{
		Email:    "testuser2@example.com",
		Password: "securepassword",
	}
	databaseConnection.Create(&user)

	// Create a mock vehicle
	vehicle := models.Vehicle{
		Name: "Toyota",
	}
	databaseConnection.Create(&vehicle)

	// Create a mock race associated with the user
	var startTime = &attributes.CustomTime{}
	startTime.Time = time.Now()

	race := models.Race{
		VehicleID:         vehicle.ID,
		StartTime:         startTime,
		EndTime:           nil,
		CollisionDuration: 3,
		DistanceCovered:   100,
		AverageSpeed:      10,
		OutOfParcours:     0,
		UserID:            user.ID,
		Type:              "manual",
		Status:            "Not Started",
	}
	databaseConnection.Create(&race)

	// Construct the request URL
	requestURL := fmt.Sprintf("/api/users/%d/races", user.ID)

	// Perform the authenticated request
	requestRecorder, _ := tests.PerformAuthenticatedRequest(http.MethodGet, requestURL, nil)

	// Assert the status code
	assert.Equal(t, http.StatusOK, requestRecorder.Code)

	// Parse and check the response body
	var response ApiResponse
	err := json.Unmarshal(requestRecorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Validate the response data
	races := response.Data
	assert.NotNil(t, races)
	assert.True(t, len(races) > 0)
}
