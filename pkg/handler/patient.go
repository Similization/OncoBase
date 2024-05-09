package handler

import (
	"med/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePatient godoc
// @Summary Create patient
// @Description Creates a new patient.
// @Tags Patient
// @Accept json
// @Produce json
// @Param input body model.Patient true "Patient data"
// @Success 200 {object} IdResponse "Created patient data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patients [post]
func (h *Handler) CreatePatient(ctx *gin.Context) {
	var patient model.Patient

	if err := ctx.BindJSON(&patient); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Patient.CreatePatient(patient)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, IdResponse{Id: id})
}

// GetPatientList godoc
// @Summary Get patient list
// @Description Retrieves a list of patients.
// @Tags Patient
// @Produce json
// @Success 200 {array} []model.Patient "Patient list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patients [get]
func (h *Handler) GetPatientList(ctx *gin.Context) {
	patientList, err := h.services.Patient.GetPatientList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, patientList)
}

// GetPatientById godoc
// @Summary Get patient by ID
// @Description Retrieves a patient by ID.
// @Tags Patient
// @Produce json
// @Param id path string true "Patient ID"
// @Success 200 {object} model.Patient "Patient data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patients/{id} [get]
func (h *Handler) GetPatientById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param(userContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	patient, err := h.services.Patient.GetPatientById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, patient)
}

// UpdatePatient godoc
// @Summary Update patient
// @Description Updates an existing patient.
// @Tags Patient
// @Accept json
// @Produce json
// @Param input body model.Patient true "Patient data"
// @Success 200 {object} string "Updated patient data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patients [put]
func (h *Handler) UpdatePatient(ctx *gin.Context) {
	var patient model.Patient

	if err := ctx.BindJSON(&patient); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Patient.UpdatePatient(patient)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "updated")
}

// DeletePatient godoc
// @Summary Delete patient
// @Description Deletes a patient by ID.
// @Tags Patient
// @Produce json
// @Param id path string true "Patient ID"
// @Success 200 {string} string "Deleted patient data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patients/{id} [delete]
func (h *Handler) DeletePatient(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param(userContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Patient.DeletePatient(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "deleted")
}
