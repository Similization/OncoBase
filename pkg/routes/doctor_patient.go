package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createDoctorPatientRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	doctorPatient := route.Group("/doctor-patient")
	{
		doctorPatient.POST("/")
		doctorPatient.GET("/")
		doctorPatient.GET("/:id")
		doctorPatient.PUT("/:id")
		doctorPatient.DELETE("/:id")
	}
	return doctorPatient
}
