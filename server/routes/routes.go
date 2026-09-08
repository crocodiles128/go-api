package routes

import (
	controllers "github.com/crocodiles128/go-api/controller"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/api/v1")
	{
		planes := main.Group("/planes")
		{
			planes.GET("/", controllers.ShowPlanes)

		}
		return router
	}
}
