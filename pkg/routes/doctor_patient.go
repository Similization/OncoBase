package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createDoctorPatientRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	doctorPatient := route.Group("/doctor-patient")
	{
		doctorPatient.POST("/", handlers.CreateDoctorPatient)
		doctorPatient.GET("/:doctor_id", handlers.GetDoctorPatientList)
		doctorPatient.DELETE("/:doctor_id/:patient_id", handlers.DeleteDoctorPatient)
	}
	return doctorPatient
}
