package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createDoctorPatientRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	doctorPatient := route.Group("/doctor-patient")
	{
		doctorPatient.POST("/", handlers.CreateDoctorPatient)
		doctorPatient.GET("/doctor/:doctor_id", handlers.GetDoctorPatientListByDoctor)
		doctorPatient.GET("/patient/:patient_id", handlers.GetDoctorPatientListByPatient)
		doctorPatient.DELETE("/:doctor_id/:patient_id", handlers.DeleteDoctorPatient)
	}
	return doctorPatient
}
