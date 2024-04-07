package handler

import (
	"med/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetPatientList(ctx *gin.Context) {
	patientList, err := h.services.Patient.GetPatientList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Get patient list": patientList,
	})
}

func (h *Handler) GetPatientById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	patient, err := h.services.Patient.GetPatientById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Get patient": patient,
	})
}

func (h *Handler) CreatePatient(ctx *gin.Context) {
	var patient model.Patient

	if err := ctx.BindJSON(&patient); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	createdPatient, err := h.services.Patient.CreatePatient(patient)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Created patient": createdPatient,
	})
}

func (h *Handler) UpdatePatient(ctx *gin.Context) {
	var patient model.Patient

	if err := ctx.BindJSON(&patient); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	updatedPatient, err := h.services.Patient.UpdatePatient(patient)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Updated patient": updatedPatient,
	})
}

func (h *Handler) DeletePatient(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Patient.DeletePatient(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Patient deleted": id,
	})
}
