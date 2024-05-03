package handler

import (
	"med/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateDoctorPatient godoc
// @Summary Create doctor-patient relationship
// @Description Creates a relationship between a doctor and a patient.
// @Tags DoctorPatient
// @Accept json
// @Produce json
// @Param input body model.DoctorPatient true "Doctor-patient relationship data"
// @Success 200 {object} model.DoctorPatient "Created doctor-patient relationship data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /doctor-patient [post]
func (h *Handler) CreateDoctorPatient(ctx *gin.Context) {
	var doctorPatient model.DoctorPatient

	if err := ctx.BindJSON(&doctorPatient); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	createdDoctorPatient, err := h.services.DoctorPatient.CreateDoctorPatient(doctorPatient)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, createdDoctorPatient)
}

// GetDoctorPatientList godoc
// @Summary Get doctor-patient relationship list by doctor ID
// @Description Retrieves a list of patient IDs associated with the given doctor ID.
// @Tags DoctorPatient
// @Produce json
// @Param doctor_id path string true "Doctor ID"
// @Success 200 {array} []model.DoctorPatient "Patient ID list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /doctor-patient/{doctor_id} [get]
func (h *Handler) GetDoctorPatientList(ctx *gin.Context) {
	doctorId, err := strconv.Atoi(ctx.Param(doctorContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	doctorPatientList, err := h.services.DoctorPatient.GetDoctorPatientList(doctorId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, doctorPatientList)
}

// DeleteDoctorPatient godoc
// @Summary Delete doctor-patient relationship
// @Description Deletes the relationship between a doctor and a patient.
// @Tags DoctorPatient
// @Produce json
// @Param doctor_id path string true "Doctor ID"
// @Param patient_id path string true "Patient ID"
// @Success 200 {object} int "Deleted doctor-patient relationship data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /doctor-patient/{doctor_id}/{patient_id} [delete]
func (h *Handler) DeleteDoctorPatient(ctx *gin.Context) {
	doctorId, err := strconv.Atoi(ctx.Param(doctorContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	patientId, err := strconv.Atoi(ctx.Param(patientContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.DoctorPatient.DeleteDoctorPatient(doctorId, patientId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"doctor_id":  doctorId,
		"patient_id": patientId,
	})
}
