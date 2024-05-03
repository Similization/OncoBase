package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createProcedureBloodCountRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	procedureBloodCount := route.Group("/procedure-blood-count")
	{
		procedureBloodCount.POST("/")
		procedureBloodCount.GET("/")
		procedureBloodCount.GET("/:id")
		procedureBloodCount.PUT("/:id")
		procedureBloodCount.DELETE("/:id")
	}
	return procedureBloodCount
}
