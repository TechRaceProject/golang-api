package public

import (
	"api/src/controllers"
	"api/src/services"

	"github.com/gin-gonic/gin"
)

func SetupPublicRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/signup", controllers.Signup)

	routerGroup.POST("/login", controllers.Login)

	routerGroup.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	//todo: passer cette route en authenticated
	routerGroup.GET("/sse", services.SSEHandler)
}
