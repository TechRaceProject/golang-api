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

func Test_cannot_create_race_if_invalid_race_type_provided(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	databaseConnection.AutoMigrate(&models.Vehicle{}, &models.Race{})

	user := models.User{
		Email:    "testuser98@example.com",
		Password: "securepassword",
	}
	databaseConnection.Create(&user)

	vehicle := models.Vehicle{
		Name: "Toyota",
	}
	databaseConnection.Create(&vehicle)

	startTime := attributes.CustomTime{Time: time.Now()}
	endTime := &attributes.CustomTime{Time: startTime.Add(time.Minute)}

	createBody, _ := json.Marshal(map[string]interface{}{
		"name":               "testuser",
		"start_time":         startTime,
		"end_time":           endTime,
		"collision_duration": 5,
		"distance_covered":   150,
		"average_speed":      130,
		"out_of_parcours":    1,
		"user_id":            user.ID,
		"vehicle_id":         vehicle.ID,
		"type":               "invalid",
		"status":             "not_started",
	})

	requestURL := fmt.Sprintf("/api/users/%d/races", user.ID)
	requestRecorder, _ := tests.PerformAuthenticatedRequest(http.MethodPost, requestURL, createBody)

	responseBody := requestRecorder.Body.String()
	assert.Equal(t, http.StatusUnprocessableEntity, requestRecorder.Code)

	expectedErrorMessage := "CreateRaceValidator.Type"
	assert.Contains(t, responseBody, expectedErrorMessage)
}
