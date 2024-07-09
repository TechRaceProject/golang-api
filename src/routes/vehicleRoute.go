package routes

import (
	"api/src/controllers"
	"github.com/gin-gonic/gin"
)

func VehicleRoutes(router *gin.Engine) {
	vehicleGroup := router.Group("/api")
	{
		vehicleGroup.GET("/vehicle/:id", controllers.GetVehicle)
		vehicleGroup.GET("/vehicles", controllers.GetVehicles)
		vehicleGroup.POST("/vehicle", controllers.CreateVehicle)
		vehicleGroup.PATCH("/vehicle/:id", controllers.UpdateVehicle)
		vehicleGroup.DELETE("/vehicle/:id", controllers.DeleteVehicle)
	}
}