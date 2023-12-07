package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createAccountRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	account := route.Group("/account", handlers.UserIdentity)
	{
		account.GET("/settings", handlers.Settings)
		account.GET("/blood-count", handlers.BloodCount)
		account.GET("/doctors", handlers.Doctors)
		// account.GET("/patients-data")
		// account.GET("/console")
	}
	return account
}
