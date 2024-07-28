package tests

import (
	"api/src/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"api/src/tests"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCanUpdateVehicle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := tests.GetTestDBConnection()
	db.AutoMigrate(&models.Vehicle{})
	vehicle := tests.setupTestVehicle(db)

	updatedData := map[string]interface{}{
		"vehicle_name": "Updated Vehicle",
		"battery_life": 95.0,
	}
	body, _ := json.Marshal(updatedData)
	recorder := httptest.NewRecorder()
	router := tests.GetTestRouter()
	request, _ := http.NewRequest(http.MethodPatch, "/api/vehicles/"+strconv.Itoa(int(vehicle.ID)), bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string]interface{}
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Equal(t, "Updated Vehicle", response["data"].(map[string]interface{})["attributes"].(map[string]interface{})["vehicle_name"])
}
