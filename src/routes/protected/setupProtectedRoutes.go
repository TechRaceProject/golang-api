package protected

import (
	controllers "api/src/controllers"
	race_controller "api/src/controllers/race"

	//user_controller "api/src/controllers/user"
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

	/*userGroup := routerGroup.Group("/user")
	// USER ROUTES
	{
		userGroup.GET("/:userId", user_controller.UserHandler)
		userGroup.PATCH("/:userId", user_controller.UpdateUserHandler)
		userGroup.GET("/", user_controller.GetAllUserHandler)
		userGroup.DELETE("/:userId", user_controller.DeleteUserHandler)
	}*/

	raceGroup := routerGroup.Group("/race")

	// RACE ROUTES
	{
		raceGroup.GET("/", race_controller.GetAllRaceHandler)
		raceGroup.GET("/:raceId", race_controller.GetSingleRaceHandler)
		raceGroup.POST("/", race_controller.CreateRaceHandler)
		raceGroup.PATCH("/:raceId", race_controller.UpdateRaceHandler)
		raceGroup.DELETE("/:raceId", race_controller.DeleteRaceHandler)
	}

}
