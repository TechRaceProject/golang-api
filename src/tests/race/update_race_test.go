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

func Test_update_race_successfully(t *testing.T) {
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

	body, _ := json.Marshal(map[string]interface{}{
		"duration":           120,
		"elapsed_time":       110,
		"laps":               5,
		"race_type":          "TIME_TRIAL",
		"average_speed":      150,
		"total_faults":       2,
		"effective_duration": 118,
		"user_id":            1,
		"vehicle_id":         vehicle.ID,
	})

	requestURL := fmt.Sprintf("/api/races/%d", race.ID)
	requestRecorder, _ := tests.PerformAuthenticatedRequest(http.MethodPatch, requestURL, body)

	assert.Equal(t, http.StatusOK, requestRecorder.Code)

	databaseConnection.Unscoped().Delete(&vehicle)
	databaseConnection.Unscoped().Delete(&race)
}
