package handler

import (
	"med/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetDiseaseList(ctx *gin.Context) {
	diseaseList, err := h.services.Disease.GetDiseaseList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Get disease list": diseaseList,
	})
}

func (h *Handler) GetDiseaseById(ctx *gin.Context) {
	id := ctx.Param("id")

	disease, err := h.services.Disease.GetDiseaseById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Get disease": disease,
	})
}

func (h *Handler) CreateDisease(ctx *gin.Context) {
	var disease model.Disease

	if err := ctx.BindJSON(&disease); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	createdDisease, err := h.services.Disease.CreateDisease(disease)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Created disease": createdDisease,
	})
}

func (h *Handler) UpdateDisease(ctx *gin.Context) {
	var disease model.Disease

	if err := ctx.BindJSON(&disease); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	updatedDisease, err := h.services.Disease.UpdateDisease(disease)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Updated disease": updatedDisease,
	})
}

func (h *Handler) DeleteDisease(ctx *gin.Context) {
	id := ctx.Param(userContext)

	err := h.services.Disease.DeleteDisease(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Disease deleted": id,
	})
}
