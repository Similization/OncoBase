package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createDrugRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	drug := route.Group("/drug")
	{
		drug.POST("/", handlers.CreateDrug)
		drug.GET("/", handlers.GetDrugList)
		drug.GET("/:id", handlers.GetDrugById)
		drug.PUT("/:id", handlers.UpdateDrug)
		drug.DELETE("/:id", handlers.DeleteDrug)
	}
	return drug
}
