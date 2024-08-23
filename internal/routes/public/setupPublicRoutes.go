package public

import (
	handlers "api/internal/handler"
	services "api/pkg/httputils"

	"github.com/gin-gonic/gin"
)

func SetupPublicRoutes(routerGroup *gin.RouterGroup) {
	routerGroup.POST("/signup", handlers.Signup)

	routerGroup.POST("/login", handlers.Login)

	routerGroup.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	//todo: passer cette route en authenticated
	routerGroup.GET("/sse", services.SSEHandler)
}
