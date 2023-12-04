package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createAuthRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	auth := route.Group("/auth")
	{
		auth.POST("/login", handlers.LogIn)
		auth.POST("/registry", handlers.Registry)
		auth.POST("/logout", handlers.LogOut)
	}
	return auth
}
