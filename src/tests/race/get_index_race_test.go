package race

import (
	"api/src/models"
	"api/src/tests"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type ApiResponse struct {
	Message string        `json:"message"`
	Data    []models.Race `json:"data"`
}

func Test_get_races_index(t *testing.T) {
	gin.SetMode(gin.TestMode)

	databaseConnection := tests.GetTestDBConnection()

	databaseConnection.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.Race{})

	requestRecorder, _ := tests.PerformAuthenticatedRequest(http.MethodGet, "/api/race/", nil)

	assert.Equal(t, http.StatusOK, requestRecorder.Code)

	var response ApiResponse
	err := json.Unmarshal(requestRecorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	races := response.Data
	assert.NotNil(t, races)
	assert.True(t, len(races) >= 0)
}
