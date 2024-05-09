package handler

import (
	"med/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateUnitMeasure godoc
// @Summary Create unit measure
// @Description Creates a new unit measure entry.
// @Tags UnitMeasure
// @Accept json
// @Produce json
// @Param input body model.UnitMeasure true "Unit measure data"
// @Success 200 {object} string "Created unit measure data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /unit-measure [post]
func (h *Handler) CreateUnitMeasure(ctx *gin.Context) {
	var unitMeasure model.UnitMeasure

	if err := ctx.BindJSON(&unitMeasure); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.UnitMeasure.CreateUnitMeasure(unitMeasure)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "created")
}

// GetUnitMeasureList godoc
// @Summary Get unit measure list
// @Description Retrieves a list of unit measure entries.
// @Tags UnitMeasure
// @Produce json
// @Success 200 {object} []model.UnitMeasure "Unit measure list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /unit-measure [get]
func (h *Handler) GetUnitMeasureList(ctx *gin.Context) {
	unitMeasureList, err := h.services.UnitMeasure.GetUnitMeasureList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, unitMeasureList)
}

// GetUnitMeasureById godoc
// @Summary Get unit measure by ID
// @Description Retrieves a unit measure entry by ID.
// @Tags UnitMeasure
// @Produce json
// @Param id path string true "Unit measure ID"
// @Success 200 {object} model.UnitMeasure "Unit measure data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /unit-measure/{id} [get]
func (h *Handler) GetUnitMeasureById(ctx *gin.Context) {
	id := ctx.Param(userContext)

	unitMeasure, err := h.services.UnitMeasure.GetUnitMeasureById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, unitMeasure)
}

// UpdateUnitMeasure godoc
// @Summary Update unit measure
// @Description Updates an existing unit measure entry.
// @Tags UnitMeasure
// @Accept json
// @Produce json
// @Param input body model.UnitMeasure true "Unit measure data"
// @Success 200 {object} string "Updated unit measure data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /unit-measure [put]
func (h *Handler) UpdateUnitMeasure(ctx *gin.Context) {
	var unitMeasure model.UnitMeasure

	if err := ctx.BindJSON(&unitMeasure); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.UnitMeasure.UpdateUnitMeasure(unitMeasure)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "updated")
}

// DeleteUnitMeasure godoc
// @Summary Delete unit measure
// @Description Deletes a unit measure entry by ID.
// @Tags UnitMeasure
// @Produce json
// @Param id path string true "Unit measure ID"
// @Success 200 {object} string "Deleted unit measure data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /unit-measure/{id} [delete]
func (h *Handler) DeleteUnitMeasure(ctx *gin.Context) {
	id := ctx.Param(userContext)

	err := h.services.UnitMeasure.DeleteUnitMeasure(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "deleted")
}
