package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createPatientCourseRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	patientCourse := route.Group("/patient-course")
	{
		patientCourse.POST("/")
		patientCourse.GET("/")
		patientCourse.GET("/:id")
		patientCourse.PUT("/:id")
		patientCourse.DELETE("/:id")
	}
	return patientCourse
}
