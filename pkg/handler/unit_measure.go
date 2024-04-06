package handler

import (
	"med/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUnitMeasureList(ctx *gin.Context) {
	unitMeasureList, err := h.services.UnitMeasure.GetUnitMeasureList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Get unit measure list": unitMeasureList,
	})
}

func (h *Handler) GetUnitMeasureById(ctx *gin.Context) {
	id := ctx.Param("id")

	unitMeasure, err := h.services.UnitMeasure.GetUnitMeasureById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Get unit measure": unitMeasure,
	})
}

func (h *Handler) CreateUnitMeasure(ctx *gin.Context) {
	var unitMeasure model.UnitMeasure

	if err := ctx.BindJSON(&unitMeasure); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	createdUnitMeasure, err := h.services.UnitMeasure.CreateUnitMeasure(unitMeasure)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Created unit measure": createdUnitMeasure,
	})
}

func (h *Handler) UpdateUnitMeasure(ctx *gin.Context) {
	var unitMeasure model.UnitMeasure

	if err := ctx.BindJSON(&unitMeasure); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	updatedUnitMeasure, err := h.services.UnitMeasure.UpdateUnitMeasure(unitMeasure)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Updated unit measure": updatedUnitMeasure,
	})
}

func (h *Handler) DeleteUnitMeasure(ctx *gin.Context) {
	id := ctx.Param(userContext)

	err := h.services.UnitMeasure.DeleteUnitMeasure(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"Unit measure deleted": id,
	})
}
