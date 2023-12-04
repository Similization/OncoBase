package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createBloodCountRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	bloodCount := route.Group("/blood-count")
	{
		bloodCount.POST("/")
		bloodCount.GET("/")
		bloodCount.GET("/:id")
		bloodCount.PUT("/:id")
		bloodCount.DELETE("/:id")
	}
	return bloodCount
}
