package race

import (
	"api/src/models"
	"api/src/models/attributes"
	"api/src/tests"
	"fmt"
	"net/http"
	"testing"
	"time"

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

	var startTime attributes.CustomTime
	startTime.Time = time.Now()

	race := models.Race{
		VehicleID:         vehicle.ID,
		StartTime:         startTime,
		CollisionDuration: 3,
		DistanceCovered:   100,
		AverageSpeed:      10,
		OutOfParcours:     0,
		UserID:            1,
		Type:              "manual",
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
