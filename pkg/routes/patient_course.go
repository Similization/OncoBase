package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createPatientCourseRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	patientCourse := route.Group("/patient-course")
	{
		patientCourse.POST("/", handlers.CreatePatientCourse)
		patientCourse.GET("/", handlers.GetPatientCourseList)
		patientCourse.GET("/:id", handlers.GetPatientCourseById)
		patientCourse.PUT("/:id", handlers.UpdatePatientCourse)
		patientCourse.DELETE("/:id", handlers.DeletePatientCourse)
	}
	return patientCourse
}
