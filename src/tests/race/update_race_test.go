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

func Test_update_race_successfully(t *testing.T) {
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
		VehicleID:          vehicle.ID,
		StartTime:          startTime,
		EndTime:            nil,
		NumberOfCollisions: 3,
		DistanceTravelled:  100,
		AverageSpeed:       120,
		OutOfParcours:      0,
		UserID:             1,
	}
	databaseConnection.Create(&race)

	updateBody, _ := json.Marshal(map[string]interface{}{
		"end_time": startTime.Add(time.Minute).Format(time.RFC3339), // Updated end time
		"status":   "not_started",
	})

	requestURL := fmt.Sprintf("/api/races/%d", race.ID)
	requestRecorder, _ := tests.PerformAuthenticatedRequest(http.MethodPatch, requestURL, updateBody)

	assert.Equal(t, http.StatusOK, requestRecorder.Code)

	databaseConnection.Unscoped().Delete(&vehicle)
	databaseConnection.Unscoped().Delete(&race)
}
