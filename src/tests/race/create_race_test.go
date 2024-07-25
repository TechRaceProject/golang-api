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

func Test_create_race_successfully(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	databaseConnection.AutoMigrate(&models.Vehicle{}, &models.Race{})

	vehicle := models.Vehicle{
		Name: "Toyota",
	}
	databaseConnection.Create(&vehicle)

	body, _ := json.Marshal(map[string]interface{}{
		"duration":           120,
		"elapsed_time":       110,
		"laps":               5,
		"race_type":          "VS",
		"average_speed":      150,
		"total_faults":       2,
		"effective_duration": 118,
		"user_id":            1,
		"vehicle_id":         vehicle.ID,
	})

	requestRecorder, _ := tests.PerformAuthenticatedRequest(http.MethodPost, "/api/races/", body)

	assert.Equal(t, http.StatusCreated, requestRecorder.Code)
}
