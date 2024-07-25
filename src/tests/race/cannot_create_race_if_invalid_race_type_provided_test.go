package race

import (
	"api/src/models"
	"api/src/tests"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_cannot_create_race_if_invalid_race_type_provided(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	databaseConnection.AutoMigrate(&models.Vehicle{}, &models.Race{})

	vehicle := models.Vehicle{
		ID:   1,
		Name: "Toyota",
	}
	databaseConnection.Create(&vehicle)

	body, _ := json.Marshal(map[string]interface{}{
		"duration":           120,
		"elapsed_time":       110,
		"laps":               5,
		"race_type":          "INVALID RACE TYPE",
		"average_speed":      150,
		"total_faults":       2,
		"effective_duration": 118,
		"user_id":            1,
		"vehicle_id":         vehicle.ID,
	})

	requestRecorder, _ := tests.PerformAuthenticatedRequest(http.MethodPost, "/api/races/", body)

	responseBody := requestRecorder.Body.String()

	assert.Equal(t, http.StatusUnprocessableEntity, requestRecorder.Code)

	expectedErrorMessage := "CreateRaceValidator.RaceType"
	assert.Contains(t, responseBody, expectedErrorMessage)
}
