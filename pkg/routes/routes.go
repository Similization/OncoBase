package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
)

type Group interface {
	Group(string, ...gin.HandlerFunc) *gin.RouterGroup
}

func InitRoutes(handlers *handler.Handler) *gin.Engine {
	router := gin.Default()
	createAuthRoutes(router, handlers)

	account := createAccountRoutes(router, handlers)

	createPatientsRoutes(account, handlers)
	// createBloodCountRoutes(patients, handlers)

	// course := createCourseRoutes(patients, handlers)
	// createProcedureRoutes(course, handlers)
	createUnitMeasureRoutes(router, handlers)
	createDiagnosisRoutes(router, handlers)
	createDiseaseRoutes(router, handlers)
	return router
}
