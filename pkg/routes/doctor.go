package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createDoctorRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	doctor := route.Group("/doctor")
	{
		doctor.POST("/")
		doctor.GET("/")
		doctor.GET("/:id")
		doctor.PUT("/:id")
		doctor.DELETE("/:id")
	}
	return doctor
}
