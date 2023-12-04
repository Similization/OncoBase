package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createPatientsRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	patients := route.Group("/patients")
	{
		patients.POST("/", handlers.CreatePatient)
		patients.GET("/", handlers.GetPatientList)
		patients.GET("/:id", handlers.GetPatientById)
		patients.PUT("/:id", handlers.UpdatePatient)
		patients.DELETE("/:id", handlers.DeletePatient)
	}
	return patients
}
