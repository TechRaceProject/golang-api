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

func TestCanGetVehicle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := tests.GetTestDBConnection()
	db.AutoMigrate(&models.Vehicle{})
	vehicle := tests.setupTestVehicle(db)

	recorder := httptest.NewRecorder()
	router := tests.GetTestRouter()
	request, _ := http.NewRequest(http.MethodGet, "/api/vehicles/"+strconv.Itoa(int(vehicle.ID)), nil)
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string]interface{}
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Equal(t, "Test Vehicle", response["data"].(map[string]interface{})["attributes"].(map[string]interface{})["vehicle_name"])
}
