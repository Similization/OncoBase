package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createBloodCountRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	bloodCount := route.Group("/blood-count")
	{
		bloodCount.POST("/", handlers.CreateBloodCount)
		bloodCount.GET("/", handlers.GetBloodCountList)
		bloodCount.GET("/:id", handlers.GetBloodCountById)
		bloodCount.PUT("/:id", handlers.UpdateBloodCount)
		bloodCount.DELETE("/:id", handlers.DeleteBloodCount)
	}
	return bloodCount
}
