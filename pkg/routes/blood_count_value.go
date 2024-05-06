package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createBloodCountValueRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	bloodCountValue := route.Group("/blood-count-value")
	{
		bloodCountValue.POST("/", handlers.CreateBloodCountValue)
		bloodCountValue.GET("/", handlers.GetBloodCountList)
		bloodCountValue.GET("/disease/:disease_id", handlers.GetBloodCountValueListByDisease)
		bloodCountValue.GET("/blood-count/:blood_count_id", handlers.GetBloodCountValueListByBloodCount)
		bloodCountValue.GET("/:disease_id/:blood_count_id", handlers.GetBloodCountValueById)
		bloodCountValue.PUT("/:disease_id/:blood_count_id", handlers.UpdateBloodCountValue)
		bloodCountValue.DELETE("/:disease_id/:blood_count_id", handlers.DeleteBloodCountValue)
	}
	return bloodCountValue
}
