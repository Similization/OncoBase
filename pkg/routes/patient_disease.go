package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createPatientDiseaseRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	patientDisease := route.Group("/patient-disease")
	{
		patientDisease.POST("/")
		patientDisease.GET("/")
		patientDisease.GET("/:id")
		patientDisease.PUT("/:id")
		patientDisease.DELETE("/:id")
	}
	return patientDisease
}
