package vehicle

import (
	"api/internal/models"
	utils "api/pkg/httputils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func IndexVehicleHandler(c *gin.Context) {
	connection := utils.GetConnection()

	var vehicles []models.Vehicle

	connection.Preload(clause.Associations).Find(&vehicles)

	utils.SetOK(c, "Vehicles successfully retrieved", vehicles)
}
