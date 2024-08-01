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

func TestCanGetVehicles(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := tests.GetTestDBConnection()
	db.AutoMigrate(&models.Vehicle{})
	tests.SetupTestVehicle(db)

	recorder, _ := tests.PerformAuthenticatedRequest(http.MethodGet, "/api/vehicles/", nil)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string]interface{}
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Greater(t, len(response["data"].([]interface{})), 0)
}
