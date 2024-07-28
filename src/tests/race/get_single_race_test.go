package race

import (
	"api/src/models"
	"api/src/tests"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_get_single_race_successfully(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	databaseConnection.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.Race{})

	vehicle := models.Vehicle{
		Name: "Toyota",
	}
	databaseConnection.Create(&vehicle)

	race := models.Race{
		Duration:          100,
		ElapsedTime:       90,
		Laps:              3,
		RaceType:          "VS",
		AverageSpeed:      120,
		TotalFaults:       1,
		EffectiveDuration: 85,
		UserID:            1,
		VehicleID:         vehicle.ID,
	}
	databaseConnection.Create(&race)

	requestURL := fmt.Sprintf("/api/race/%d", race.ID)
	requestRecorder, _ := tests.PerformAuthenticatedRequest(http.MethodGet, requestURL, nil)

	assert.Equal(t, http.StatusOK, requestRecorder.Code)

	var response map[string]interface{}
	err := json.Unmarshal(requestRecorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	data, ok := response["data"].(map[string]interface{})
	assert.True(t, ok)
	assert.Equal(t, float64(race.ID), data["ID"])
}

func Test_get_single_race_not_found(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	databaseConnection.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.Race{})

	invalidRaceID := 999999

	requestURL := fmt.Sprintf("/api/race/%d", invalidRaceID)
	requestRecorder, _ := tests.PerformAuthenticatedRequest(http.MethodGet, requestURL, nil)

	assert.Equal(t, http.StatusNotFound, requestRecorder.Code)
}
