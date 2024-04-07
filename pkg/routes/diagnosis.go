package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createDiagnosisRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	diagnosis := route.Group("/diagnosis")
	{
		diagnosis.POST("/", handlers.CreateDiagnosis)
		diagnosis.GET("/", handlers.GetDiagnosisList)
		diagnosis.GET("/:id", handlers.GetDiagnosisById)
		diagnosis.PUT("/:id", handlers.UpdateDiagnosis)
		diagnosis.DELETE("/:id", handlers.DeleteDiagnosis)
	}
	return diagnosis
}
