package handler

import (
	"med/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateDrug godoc
// @Summary Create drug
// @Description Creates a new drug.
// @Tags Drug
// @Accept json
// @Produce json
// @Param input body model.Drug true "Drug data"
// @Success 200 {object} model.Drug "Created drug data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /drugs [post]
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

	ctx.JSON(http.StatusOK, createdDrug)
}

// GetDrugList godoc
// @Summary Get drug list
// @Description Retrieves a list of drugs.
// @Tags Drug
// @Produce json
// @Success 200 {array} []model.Drug "Drug list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /drugs [get]
func (h *Handler) GetDrugList(ctx *gin.Context) {
	drugList, err := h.services.Drug.GetDrugList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, drugList)
}

// GetDrugById godoc
// @Summary Get drug by ID
// @Description Retrieves a drug by ID.
// @Tags Drug
// @Produce json
// @Param id path string true "Drug ID"
// @Success 200 {object} model.Drug "Drug data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /drugs/{id} [get]
func (h *Handler) GetDrugById(ctx *gin.Context) {
	id := ctx.Param("id")

	drug, err := h.services.Drug.GetDrugById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, drug)
}

// UpdateDrug godoc
// @Summary Update drug
// @Description Updates an existing drug.
// @Tags Drug
// @Accept json
// @Produce json
// @Param input body model.Drug true "Drug data"
// @Success 200 {object} model.Drug "Updated drug data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /drugs [put]
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

	ctx.JSON(http.StatusOK, updatedDrug)
}

// DeleteDrug godoc
// @Summary Delete drug
// @Description Deletes a drug by ID.
// @Tags Drug
// @Produce json
// @Param id path string true "Drug ID"
// @Success 200 {string} string "Drug ID deleted"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /drugs/{id} [delete]
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
