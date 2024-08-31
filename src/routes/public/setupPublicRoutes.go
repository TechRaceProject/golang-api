package public

import (
	"api/src/config"
	"api/src/controllers"
	"api/src/middleware"
	"api/src/services"

	"github.com/gin-gonic/gin"
)

func SetupPublicRoutes(routerGroup *gin.RouterGroup, cfg *config.Config) {
	routerGroup.POST("/signup", controllers.Signup, middleware.SendWelcomeEmailMiddleware(cfg))
	routerGroup.POST("/login", controllers.Login)

	routerGroup.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	routerGroup.GET("/sse", services.SSEHandler)
}
