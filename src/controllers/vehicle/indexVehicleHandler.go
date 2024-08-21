package vehicle

import (
	"api/src/models"
	"api/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func IndexVehicleHandler(c *gin.Context) {
	connection := services.GetConnection()

	var vehicles []models.Vehicle

	connection.Preload(clause.Associations).Find(&vehicles)

	services.SetOK(c, "Vehicles successfully retrieved", vehicles)
}
