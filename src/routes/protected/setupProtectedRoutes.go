package protected

import (
	controllers "api/src/controllers"
	race_controller "api/src/controllers/race"
	"api/src/controllers/vehicle"
	"api/src/controllers/vehicle_state"

	"api/src/middleware"

	"github.com/gin-gonic/gin"
)

func SetupProtectedRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.Use(middleware.AuthMiddleware())

	routerGroup.GET("/protected", controllers.Welcome)

	/*userGroup := routerGroup.Group("/user")
	// USER ROUTES
	{
		userGroup.GET("/:userId", user_controller.UserHandler)
		userGroup.PUT("/:userId", user_controller.UpdateUserHandler)
		userGroup.GET("/", user_controller.GetAllUserHandler)
		userGroup.DELETE("/:userId", user_controller.DeleteUserHandler)
	}*/

	// RACE ROUTES
	raceGroup := routerGroup.Group("/races")
	{
		raceGroup.GET("/", race_controller.GetAllRaceHandler)
		raceGroup.GET("/:raceId", race_controller.GetSingleRaceHandler)
		raceGroup.POST("/", race_controller.CreateRaceHandler)
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
