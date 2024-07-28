package tests

import (
	"api/src/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"api/src/tests"

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
	recorder := httptest.NewRecorder()
	router := tests.GetTestRouter()
	request, _ := http.NewRequest(http.MethodPost, "/api/vehicles", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	var response map[string]interface{}
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Equal(t, "New Vehicle", response["data"].(map[string]interface{})["attributes"].(map[string]interface{})["vehicle_name"])
}
