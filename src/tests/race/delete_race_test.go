package race

import (
	"api/src/models"
	"api/src/tests"
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_delete_race_successfully(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	databaseConnection.AutoMigrate(&models.Vehicle{}, &models.Race{})

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

	requestURL := fmt.Sprintf("/api/races/%d", race.ID)
	requestRecorder, _ := tests.PerformAuthenticatedRequest(http.MethodDelete, requestURL, nil)

	assert.Equal(t, http.StatusNoContent, requestRecorder.Code)
}

func Test_delete_race_not_found(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	databaseConnection.AutoMigrate(&models.Vehicle{}, &models.Race{})

	vehicle := models.Vehicle{
		Name: "Toyota",
	}
	databaseConnection.Create(&vehicle)

	nonExistentRaceID := 999
	requestURL := fmt.Sprintf("/api/races/%d", nonExistentRaceID)
	requestRecorder, _ := tests.PerformAuthenticatedRequest(http.MethodDelete, requestURL, nil)

	assert.Equal(t, http.StatusNotFound, requestRecorder.Code)
}
