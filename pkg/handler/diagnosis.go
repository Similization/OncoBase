package handler

import (
	"med/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateDiagnosis godoc
// @Summary Create diagnosis
// @Description Creates a new diagnosis
// @Tags Diagnoses
// @Accept json
// @Produce json
// @Param diagnosis body model.Diagnosis true "Diagnosis object"
// @Success 200 {object} string "Created diagnosis data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /diagnoses [post]
func (h *Handler) CreateDiagnosis(ctx *gin.Context) {
	var diagnosis model.Diagnosis

	if err := ctx.BindJSON(&diagnosis); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Diagnosis.CreateDiagnosis(diagnosis)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "created")
}

// GetDiagnosisList godoc
// @Summary Get diagnosis list
// @Description Retrieves a list of diagnoses
// @Tags Diagnoses
// @Produce json
// @Success 200 {object} []model.Diagnosis "List of diagnoses"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /diagnoses [get]
func (h *Handler) GetDiagnosisList(ctx *gin.Context) {
	diagnosisList, err := h.services.Diagnosis.GetDiagnosisList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, diagnosisList)
}

// GetDiagnosisById godoc
// @Summary Get diagnosis by ID
// @Description Retrieves a diagnosis by its ID
// @Tags Diagnoses
// @Produce json
// @Param id path string true "Diagnosis ID"
// @Success 200 {object} model.Diagnosis "Diagnosis"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /diagnoses/{id} [get]
func (h *Handler) GetDiagnosisById(ctx *gin.Context) {
	id := ctx.Param(userContext)

	diagnosis, err := h.services.Diagnosis.GetDiagnosisById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, diagnosis)
}

// UpdateDiagnosis godoc
// @Summary Update diagnosis
// @Description Updates an existing diagnosis
// @Tags Diagnoses
// @Accept json
// @Produce json
// @Param id path string true "Diagnosis ID"
// @Param diagnosis body model.Diagnosis true "Diagnosis object"
// @Success 200 {object} string "Updated diagnosis data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /diagnoses/{id} [put]
func (h *Handler) UpdateDiagnosis(ctx *gin.Context) {
	var diagnosis model.Diagnosis

	if err := ctx.BindJSON(&diagnosis); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Diagnosis.UpdateDiagnosis(diagnosis)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "updated")
}

// DeleteDiagnosis godoc
// @Summary Delete diagnosis
// @Description Deletes an existing diagnosis by its ID
// @Tags Diagnoses
// @Produce json
// @Param id path string true "Diagnosis ID"
// @Success 200 {object} string "Deleted diagnosis data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /diagnoses/{id} [delete]
func (h *Handler) DeleteDiagnosis(ctx *gin.Context) {
	id := ctx.Param(userContext)

	err := h.services.Diagnosis.DeleteDiagnosis(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "deleted")
}
