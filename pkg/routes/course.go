package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createCourseRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	course := route.Group("/course")
	{
		course.POST("/", handlers.CreateCourse)
		course.GET("/", handlers.GetCourseList)
		course.GET("/:id", handlers.GetCourseById)
		course.PUT("/:id", handlers.UpdateCourse)
		course.DELETE("/:id", handlers.DeleteCourse)
	}
	return course
}
