package handler

import (
	"med/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PatientDiseaseResponse struct {
	PatientId int `json:"patient_id"`
	DiseaseId int `json:"disease_id"`
}

// CreatePatientDisease godoc
// @Summary Create patient disease
// @Description Creates a new patient disease.
// @Tags PatientDisease
// @Accept json
// @Produce json
// @Param input body model.PatientDisease true "Patient disease data"
// @Success 200 {object} model.PatientDisease "Created patient disease data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patient-diseases [post]
func (h *Handler) CreatePatientDisease(ctx *gin.Context) {
	var patientDisease model.PatientDisease

	if err := ctx.BindJSON(&patientDisease); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	createdPatientDisease, err := h.services.PatientDisease.CreatePatientDisease(patientDisease)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, createdPatientDisease)
}

// GetPatientDiseaseList godoc
// @Summary Get patient disease list
// @Description Retrieves a list of patient diseases.
// @Tags PatientDisease
// @Produce json
// @Success 200 {array} []model.PatientDisease "Patient disease list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patient-diseases [get]
func (h *Handler) GetPatientDiseaseList(ctx *gin.Context) {
	patientDiseaseList, err := h.services.PatientDisease.GetPatientDiseaseList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, patientDiseaseList)
}

// GetPatientDiseaseListByPatient godoc
// @Summary Get patient disease list by patient
// @Description Retrieves a list of patient diseases by patient ID.
// @Tags PatientDisease
// @Produce json
// @Param patient_id path string true "Patient ID"
// @Success 200 {array} []model.PatientDisease "Patient disease list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patient-diseases/patient/{patient_id} [get]
func (h *Handler) GetPatientDiseaseListByPatient(ctx *gin.Context) {
	patientId, err := strconv.Atoi(ctx.Param(patientContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	patientDisease, err := h.services.PatientDisease.GetPatientDiseaseListByPatient(patientId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, patientDisease)
}

// GetPatientDiseaseListByDisease godoc
// @Summary Get patient disease list by disease
// @Description Retrieves a list of patient diseases by disease ID.
// @Tags PatientDisease
// @Produce json
// @Param disease_id path string true "Disease ID"
// @Success 200 {array} []model.PatientDisease "Patient disease list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patient-diseases/disease/{disease_id} [get]
func (h *Handler) GetPatientDiseaseListByDisease(ctx *gin.Context) {
	diseaseId, err := strconv.Atoi(ctx.Param(diseaseContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	patientDisease, err := h.services.PatientDisease.GetPatientDiseaseListByDisease(diseaseId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, patientDisease)
}

// GetPatientDiseaseById godoc
// @Summary Get patient disease by ID
// @Description Retrieves a patient disease by patient and disease ID.
// @Tags PatientDisease
// @Produce json
// @Param patient_id path string true "Patient ID"
// @Param disease_id path string true "Disease ID"
// @Success 200 {object} model.PatientDisease "Patient disease data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patient-diseases/{patient_id}/{disease_id} [get]
func (h *Handler) GetPatientDiseaseById(ctx *gin.Context) {
	patientId, err := strconv.Atoi(ctx.Param(patientContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	diseaseId, err := strconv.Atoi(ctx.Param(diseaseContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	patientDisease, err := h.services.PatientDisease.GetPatientDiseaseById(patientId, diseaseId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, patientDisease)
}

// UpdatePatientDisease godoc
// @Summary Update patient disease
// @Description Updates an existing patient disease.
// @Tags PatientDisease
// @Accept json
// @Produce json
// @Param input body model.PatientDisease true "Patient disease data"
// @Success 200 {object} model.PatientDisease "Updated patient disease data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patient-diseases [put]
func (h *Handler) UpdatePatientDisease(ctx *gin.Context) {
	var patientDisease model.PatientDisease

	if err := ctx.BindJSON(&patientDisease); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	updatedPatientDisease, err := h.services.PatientDisease.UpdatePatientDisease(patientDisease)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, updatedPatientDisease)
}

// DeletePatientDisease godoc
// @Summary Delete patient disease
// @Description Deletes a patient disease by patient and disease ID.
// @Tags PatientDisease
// @Produce json
// @Param patient_id path string true "Patient ID"
// @Param disease_id path string true "Disease ID"
// @Success 200 {string} PatientDiseaseResponse "Patient disease ID deleted"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patient-diseases/{patient_id}/{disease_id} [delete]
func (h *Handler) DeletePatientDisease(ctx *gin.Context) {
	patientId, err := strconv.Atoi(ctx.Param(patientContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	diseaseId, err := strconv.Atoi(ctx.Param(diseaseContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.PatientDisease.DeletePatientDisease(patientId, diseaseId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, PatientDiseaseResponse{PatientId: patientId, DiseaseId: diseaseId})
}
