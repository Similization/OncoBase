package handler

import (
	"med/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetDrugList(ctx *gin.Context) {
	drugList, err := h.services.Drug.GetDrugList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Get drug list": drugList,
	})
}

func (h *Handler) GetDrugById(ctx *gin.Context) {
	id := ctx.Param("id")

	drug, err := h.services.Drug.GetDrugById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Get drug": drug,
	})
}

func (h *Handler) CreateDrug(ctx *gin.Context) {
	var drug model.Drug

	if err := ctx.BindJSON(&drug); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	createdDrug, err := h.services.Drug.CreateDrug(drug)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Created drug": createdDrug,
	})
}

func (h *Handler) UpdateDrug(ctx *gin.Context) {
	var drug model.Drug

	if err := ctx.BindJSON(&drug); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	updatedDrug, err := h.services.Drug.UpdateDrug(drug)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Updated drug": updatedDrug,
	})
}

func (h *Handler) DeleteDrug(ctx *gin.Context) {
	id := ctx.Param(userContext)

	err := h.services.Drug.DeleteDrug(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Drug deleted": id,
	})
}
