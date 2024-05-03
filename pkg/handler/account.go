package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandlers struct {
}

// Settings godoc
// @Summary Get account settings
// @Description Retrieves the account settings for the authenticated user.
// @Tags Account
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} int "Account settings"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Router /account/settings [get]
func (h *AccountHandlers) Settings(ctx *gin.Context) {
	id, _ := ctx.Get(userContext)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// BloodCount godoc
// @Summary Get blood count
// @Description Retrieves blood count data.
// @Tags Account
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} array "Blood count data"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Router /account/blood-count [get]
func (h *AccountHandlers) BloodCount(ctx *gin.Context) {

}

// Doctors godoc
// @Summary Get doctors
// @Description Retrieves doctor information.
// @Tags Account
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} array "Doctor list"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Router /account/doctors [get]
func (h *AccountHandlers) Doctors(ctx *gin.Context) {

}

// PatientData godoc
// @Summary Get patient data
// @Description Retrieves patient data.
// @Tags Account
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} array "Patient data"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Router /account/patient-data [get]
func (h *AccountHandlers) PatientData(ctx *gin.Context) {

}

// Console godoc
// @Summary Get console
// @Description Retrieves console data.
// @Tags Account
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} array "Console data"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Router /account/console [get]
func (h *AccountHandlers) Console(ctx *gin.Context) {

}
