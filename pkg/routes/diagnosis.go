package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createDiagnosisRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	unitMeasure := route.Group("/diagnosis")
	{
		unitMeasure.POST("/", handlers.CreateDiagnosis)
		unitMeasure.GET("/", handlers.GetDiagnosisList)
		unitMeasure.GET("/:id", handlers.GetDiagnosisById)
		unitMeasure.PUT("/:id", handlers.UpdateDiagnosis)
		unitMeasure.DELETE("/:id", handlers.DeleteDiagnosis)
	}
	return unitMeasure
}
