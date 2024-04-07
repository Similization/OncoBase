package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createPatientsRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	patient := route.Group("/patient")
	{
		patient.POST("/", handlers.CreatePatient)
		patient.GET("/", handlers.GetPatientList)
		patient.GET("/:id", handlers.GetPatientById)
		patient.PUT("/:id", handlers.UpdatePatient)
		patient.DELETE("/:id", handlers.DeletePatient)
	}
	return patient
}
