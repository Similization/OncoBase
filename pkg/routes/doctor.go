package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createDoctorRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	doctor := route.Group("/doctor")
	{
		doctor.POST("/", handlers.CreateDoctor)
		doctor.GET("/", handlers.GetDoctorList)
		doctor.GET("/:id", handlers.GetDoctorById)
		doctor.PUT("/:id", handlers.UpdateDoctor)
		doctor.DELETE("/:id", handlers.DeleteDoctor)
	}
	return doctor
}
