package protected

import (
	handlers "api/internal/handler"
	race_controller "api/internal/handler/race"
	"api/internal/handler/vehicle"
	"api/internal/handler/vehicle_state"

	"api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middleware.AuthMiddleware())

	routerGroup.GET("/protected", handlers.Welcome)

	userGroup := routerGroup.Group("/users")
	// USER ROUTES
	{
		userGroup.POST("/:userId/races", race_controller.CreateRaceHandler)
		userGroup.GET("/:userId/races", race_controller.GetAllUsersRaceHandler)

	}

	// RACE ROUTES
	raceGroup := routerGroup.Group("/races")
	{
		raceGroup.GET("/", race_controller.GetAllRaceHandler)
		raceGroup.PATCH("/:raceId", race_controller.UpdateRaceHandler)
		raceGroup.DELETE("/:raceId", race_controller.DeleteRaceHandler)
	}

	// VEHICLE ROUTES
	vehicleGroup := routerGroup.Group("/vehicles")
	{
		vehicleGroup.GET("/", vehicle.IndexVehicleHandler)
		vehicleGroup.GET("/:id", vehicle.GetVehicleHandler)
	}

	// VEHICLE STATE ROUTES
	vehicleStateGroup := routerGroup.Group("/vehicle-states")
	{
		vehicleStateGroup.PATCH("/:id", vehicle_state.UpdateVehicleStateHandler)
		vehicleStateGroup.GET("/:id", vehicle_state.GetASingleVehiculeStateHandler)
	}

}
