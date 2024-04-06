package handler

import (
	"med/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetDiagnosisList(ctx *gin.Context) {
	diagnosisList, err := h.services.Diagnosis.GetDiagnosisList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Get diagnosis list": diagnosisList,
	})
}

func (h *Handler) GetDiagnosisById(ctx *gin.Context) {
	id := ctx.Param("id")

	diagnosis, err := h.services.Diagnosis.GetDiagnosisById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Get diagnosis": diagnosis,
	})
}

func (h *Handler) CreateDiagnosis(ctx *gin.Context) {
	var diagnosis model.Diagnosis

	if err := ctx.BindJSON(&diagnosis); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	createdDiagnosis, err := h.services.Diagnosis.CreateDiagnosis(diagnosis)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Created diagnosis": createdDiagnosis,
	})
}

func (h *Handler) UpdateDiagnosis(ctx *gin.Context) {
	var diagnosis model.Diagnosis

	if err := ctx.BindJSON(&diagnosis); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	updatedDiagnosis, err := h.services.Diagnosis.UpdateDiagnosis(diagnosis)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Updated diagnosis": updatedDiagnosis,
	})
}

func (h *Handler) DeleteDiagnosis(ctx *gin.Context) {
	id := ctx.Param(userContext)

	err := h.services.Diagnosis.DeleteDiagnosis(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Diagnosis deleted": id,
	})
}
