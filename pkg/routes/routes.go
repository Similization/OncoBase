package route

import (
	"med/pkg/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// _ "github.com/Similization/OncoBase/docs"
)

type Group interface {
	Group(string, ...gin.HandlerFunc) *gin.RouterGroup
}

func InitRoutes(handlers *handler.Handler) *gin.Engine {
	router := gin.Default()
	createAuthRoutes(router, handlers)

	account := createAccountRoutes(router, handlers)

	createBloodCountRoutes(router, handlers)
	createBloodCountValueRoutes(router, handlers)

	createCourseRoutes(router, handlers)

	createDiagnosisRoutes(router, handlers)
	createDiseaseRoutes(router, handlers)
	createDoctorRoutes(router, handlers)
	createDoctorPatientRoutes(router, handlers)
	createDrugRoutes(router, handlers)

	createPatientsRoutes(account, handlers)
	createPatientCourseRoutes(router, handlers)
	createPatientDiseaseRoutes(router, handlers)
	createProcedureBloodCountRoutes(router, handlers)

	createUnitMeasureRoutes(router, handlers)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
