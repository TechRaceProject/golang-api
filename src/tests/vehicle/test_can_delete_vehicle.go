package tests

import (
	"api/src/models"
	"api/src/tests"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCanDeleteVehicle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := tests.GetTestDBConnection()
	db.AutoMigrate(&models.Vehicle{})
	vehicle := tests.SetupTestVehicle(db)

	recorder := httptest.NewRecorder()
	router := tests.GetTestRouter()
	request, _ := http.NewRequest(http.MethodDelete, "/api/vehicles/"+strconv.Itoa(int(vehicle.ID)), nil)
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string]interface{}
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Equal(t, "Vehicule deleted successfully", response["meta"].(map[string]interface{})["message"])
}
