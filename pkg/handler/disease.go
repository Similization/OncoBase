package handler

import (
	"med/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateDisease godoc
// @Summary Create a new disease
// @Description Creates a new disease entry.
// @Tags Disease
// @Accept json
// @Produce json
// @Param input body model.Disease true "Disease data"
// @Success 200 {object} model.Disease "Created disease data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /disease [post]
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

	ctx.JSON(http.StatusOK, createdDisease)
}

// GetDiseaseList godoc
// @Summary Get disease list
// @Description Retrieves a list of diseases.
// @Tags Disease
// @Produce json
// @Success 200 {array} []model.Disease "Disease list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /disease [get]
func (h *Handler) GetDiseaseList(ctx *gin.Context) {
	diseaseList, err := h.services.Disease.GetDiseaseList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, diseaseList)
}

// GetDiseaseById godoc
// @Summary Get disease by ID
// @Description Retrieves a disease entry by its ID.
// @Tags Disease
// @Produce json
// @Param id path string true "Disease ID"
// @Success 200 {object} model.Disease "Disease data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /disease/{id} [get]
func (h *Handler) GetDiseaseById(ctx *gin.Context) {
	id := ctx.Param(userContext)

	disease, err := h.services.Disease.GetDiseaseById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, disease)
}

// UpdateDisease godoc
// @Summary Update disease
// @Description Updates an existing disease entry.
// @Tags Disease
// @Accept json
// @Produce json
// @Param input body model.Disease true "Updated disease data"
// @Success 200 {object} model.Disease "Updated disease data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /disease [put]
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

	ctx.JSON(http.StatusOK, updatedDisease)
}

// DeleteDisease godoc
// @Summary Delete disease
// @Description Deletes a disease entry by its ID.
// @Tags Disease
// @Produce json
// @Param id path string true "Disease ID"
// @Success 200 {string} string "Disease ID"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /disease/{id} [delete]
func (h *Handler) DeleteDisease(ctx *gin.Context) {
	id := ctx.Param(userContext)

	err := h.services.Disease.DeleteDisease(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)
}
