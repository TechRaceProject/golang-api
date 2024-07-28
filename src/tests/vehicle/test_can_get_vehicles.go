package tests

import (
	"api/src/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"api/src/tests"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCanGetVehicles(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := tests.GetTestDBConnection()
	db.AutoMigrate(&models.Vehicle{})
	tests.setupTestVehicle(db)

	recorder := httptest.NewRecorder()
	router := tests.GetTestRouter()
	request, _ := http.NewRequest(http.MethodGet, "/api/vehicles", nil)
	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var response map[string]interface{}
	json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.Greater(t, len(response["data"].([]interface{})), 0)
}
