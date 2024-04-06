package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createDiseaseRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	unitMeasure := route.Group("/disease")
	{
		unitMeasure.POST("/", handlers.CreateDisease)
		unitMeasure.GET("/", handlers.GetDiseaseList)
		unitMeasure.GET("/:id", handlers.GetDiseaseById)
		unitMeasure.PUT("/:id", handlers.UpdateDisease)
		unitMeasure.DELETE("/:id", handlers.DeleteDisease)
	}
	return unitMeasure
}
