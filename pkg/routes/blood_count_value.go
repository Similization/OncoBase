package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createBloodCountValueRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	bloodCountValue := route.Group("/blood-count-value")
	{
		bloodCountValue.POST("/")
		bloodCountValue.GET("/")
		bloodCountValue.GET("/:id")
		bloodCountValue.PUT("/:id")
		bloodCountValue.DELETE("/:id")
	}
	return bloodCountValue
}
