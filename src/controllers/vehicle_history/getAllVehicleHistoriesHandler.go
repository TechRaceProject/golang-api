package vehicle_history

import (
	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
)

func GetAllVehicleHistoriesHandler(c *gin.Context) {
	connection := services.GetConnection()

	var vehicleHistories []models.VehicleHistory

	connection.Preload("Vehicle").Find(&vehicleHistories)

	services.SetOK(c, "Vehicules histories successfully retrieved", vehicleHistories)
}
