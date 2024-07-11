package protected

import (
	"api/src/controllers"
	"api/src/middleware"

	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middleware.AuthMiddleware())

	routerGroup.GET("/protected", controllers.Welcome)

	vehicleGroup := routerGroup.Group("/vehicles")
	{
		vehicleGroup.GET("/:id", controllers.GetVehicle)
		vehicleGroup.GET("/", controllers.GetVehicles)
		vehicleGroup.POST("/", controllers.CreateVehicle)
		vehicleGroup.PATCH("/:id", controllers.UpdateVehicle)
		vehicleGroup.DELETE("/:id", controllers.DeleteVehicle)

	}
}
