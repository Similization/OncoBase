package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createProcedureRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	procedure := route.Group("/course")
	{
		procedure.POST("/")
		procedure.GET("/")
		procedure.GET("/:id")
		procedure.PUT("/:id")
		procedure.DELETE("/:id")
	}
	return procedure
}
