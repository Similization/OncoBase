package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createProcedureBloodCountRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	procedureBloodCount := route.Group("/procedure-blood-count")
	{
		procedureBloodCount.POST("/", handlers.CreateProcedureBloodCount)
		procedureBloodCount.GET("/", handlers.GetProcedureBloodCountList)
		procedureBloodCount.GET("/procedure/:procedure_id", handlers.GetProcedureBloodCountListByProcedure)
		procedureBloodCount.GET("/blood-count/:blood_count_id", handlers.GetProcedureBloodCountListByBloodCount)
		procedureBloodCount.GET("/:procedure_id/:blood_count_id", handlers.GetProcedureBloodCountById)
		procedureBloodCount.PUT("/:procedure_id/:blood_count_id", handlers.UpdateProcedureBloodCount)
		procedureBloodCount.DELETE("/:procedure_id/:blood_count_id", handlers.DeleteProcedureBloodCount)
	}
	return procedureBloodCount
}
