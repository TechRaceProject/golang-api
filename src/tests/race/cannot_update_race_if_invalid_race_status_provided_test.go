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

func Test_cannot_update_race_if_invalid_race_type_provided(t *testing.T) {

	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	databaseConnection.AutoMigrate(&models.Vehicle{}, &models.Race{})

	vehicle := models.Vehicle{
		Name: "Toyota",
	}
	databaseConnection.Create(&vehicle)

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
		UserID:            1,
	}
	databaseConnection.Create(&race)

	endTime := &attributes.CustomTime{
		Time: startTime.Add(time.Minute),
	}

	updateBody, _ := json.Marshal(map[string]interface{}{
		"end_time": endTime,
		"status":   "invalid",
	})

	requestURL := fmt.Sprintf("/api/races/%d", race.ID)
	requestRecorder, _ := tests.PerformAuthenticatedRequest(http.MethodPatch, requestURL, updateBody)

	responseBody := requestRecorder.Body.String()
	fmt.Println("Response Body:", responseBody)

	assert.Equal(t, http.StatusUnprocessableEntity, requestRecorder.Code)

	expectedErrorMessage := "UpdateRaceValidator.Status"
	assert.Contains(t, responseBody, expectedErrorMessage)

}
