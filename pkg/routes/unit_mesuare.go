package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createUnitMeasureRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	unitMeasure := route.Group("/unit-measure")
	{
		unitMeasure.POST("/", handlers.CreateUnitMeasure)
		unitMeasure.GET("/", handlers.GetUnitMeasureList)
		unitMeasure.GET("/:id", handlers.GetUnitMeasureById)
		unitMeasure.PUT("/:id", handlers.UpdateUnitMeasure)
		unitMeasure.DELETE("/:id", handlers.DeleteUnitMeasure)
	}
	return unitMeasure
}
