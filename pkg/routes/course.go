package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createCourseRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	course := route.Group("/course")
	{
		course.POST("/")
		course.GET("/")
		course.GET("/:id")
		course.PUT("/:id")
		course.DELETE("/:id")
	}
	return course
}
