package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createPatientDiseaseRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	patientDisease := route.Group("/patient-disease")
	{
		patientDisease.POST("/", handlers.CreatePatientDisease)
		patientDisease.GET("/", handlers.GetPatientDiseaseList)
		patientDisease.GET("/disease/:disease_id", handlers.GetPatientDiseaseListByDisease)
		patientDisease.GET("/patient/:patient_id", handlers.GetPatientDiseaseListByPatient)
		patientDisease.GET("/:patient_id/:disease_id", handlers.GetPatientDiseaseById)
		patientDisease.PUT("/:patient_id/:disease_id", handlers.UpdatePatientDisease)
		patientDisease.DELETE("/:patient_id/:disease_id", handlers.DeletePatientDisease)
	}
	return patientDisease
}
