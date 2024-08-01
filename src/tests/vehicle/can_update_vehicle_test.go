package tests

import (
	"api/src/models"
	"api/src/tests"
	"encoding/json"
	"net/http"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCanUpdateVehicle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := tests.GetTestDBConnection()
	db.AutoMigrate(&models.Vehicle{})
	vehicle := tests.SetupTestVehicle(db)

	updatedData := map[string]interface{}{
		"vehicle_name": "Updated Vehicle",
		"battery_life": 95.0,
	}
	body, _ := json.Marshal(updatedData)

	recorder, _ := tests.PerformAuthenticatedRequest(http.MethodPatch, "/api/vehicles/"+strconv.Itoa(int(vehicle.ID)), body)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string]interface{}
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Equal(t, "Updated Vehicle", response["data"].(map[string]interface{})["attributes"].(map[string]interface{})["vehicle_name"])
}
