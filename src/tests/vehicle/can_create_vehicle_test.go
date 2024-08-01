package tests

import (
	"api/src/models"
	"api/src/tests"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCanCreateVehicle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := tests.GetTestDBConnection()
	db.AutoMigrate(&models.Vehicle{})

	vehicle := models.Vehicle{
		Name:          "New Vehicle",
		BatteryLife:   90.0,
		LineSensor1:   true,
		LineSensor2:   false,
		LineSensor3:   true,
		Camera:        true,
		SonarRange:    40.0,
		WheelPower1:   85,
		WheelPower2:   75,
		WheelPower3:   65,
		WheelPower4:   55,
		LedColor:      "blue",
		DisplayPanel:  "OLED",
		SpeakerStatus: true,
		SoundPlaying:  "new sound",
	}

	body, _ := json.Marshal(vehicle)

	requestRecorder, _ := tests.PerformAuthenticatedRequest(http.MethodPost, "/api/vehicles", body)

	assert.Equal(t, http.StatusCreated, requestRecorder.Code)

	var response map[string]interface{}

	json.Unmarshal(requestRecorder.Body.Bytes(), &response)

	assert.Equal(t, "New Vehicle", response["data"].(map[string]interface{})["attributes"].(map[string]interface{})["vehicle_name"])
}
