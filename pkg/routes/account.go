package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

func createAccountRoutes[G Group](route G, handlers *handler.Handler) *gin.RouterGroup {
	account := route.Group("/account", handlers.UserIdentity)
	{
		account.GET("/settings", handlers.AccountHandler.Settings)
		account.GET("/blood-count", handlers.PatientIdentity, handlers.AccountHandler.BloodCount)
		account.GET("/doctors", handlers.PatientIdentity, handlers.AccountHandler.Doctors)
		account.GET("/patients-data", handlers.DoctorIdentity, handlers.AccountHandler.PatientData)
		account.GET("/console", handlers.AdminIdentity, handlers.AccountHandler.Console)
	}
	return account
}
