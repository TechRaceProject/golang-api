package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "api/src/services"
    "api/src/models"
)

func GetAllSensorData(c *gin.Context) {
    var sensorData []models.SensorData
    if err := services.GetConnection().Find(&sensorData).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, sensorData)
}
