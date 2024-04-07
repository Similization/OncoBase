package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createDiseaseRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	disease := route.Group("/disease")
	{
		disease.POST("/", handlers.CreateDisease)
		disease.GET("/", handlers.GetDiseaseList)
		disease.GET("/:id", handlers.GetDiseaseById)
		disease.PUT("/:id", handlers.UpdateDisease)
		disease.DELETE("/:id", handlers.DeleteDisease)
	}
	return disease
}
