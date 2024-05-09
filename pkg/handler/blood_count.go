package handler

import (
	"med/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBloodCount godoc
// @Summary Create a new blood count
// @Description Creates a new blood count entry.
// @Tags BloodCount
// @Accept json
// @Produce json
// @Param input body model.BloodCount true "Blood count data"
// @Success 200 {object} string "Created blood count data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /blood-count [post]
func (h *Handler) CreateBloodCount(ctx *gin.Context) {
	var bloodCount model.BloodCount

	if err := ctx.BindJSON(&bloodCount); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.BloodCount.CreateBloodCount(bloodCount)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "created")
}

// GetBloodCountList godoc
// @Summary Get blood count list
// @Description Retrieves a list of blood counts.
// @Tags BloodCount
// @Produce json
// @Success 200 {array} []model.BloodCount "Blood count list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /blood-count [get]
func (h *Handler) GetBloodCountList(ctx *gin.Context) {
	bloodCountList, err := h.services.BloodCount.GetBloodCountList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, bloodCountList)
}

// GetBloodCountById godoc
// @Summary Get blood count by ID
// @Description Retrieves a blood count entry by its ID.
// @Tags BloodCount
// @Produce json
// @Param id path string true "Blood count ID"
// @Success 200 {object} model.BloodCount "Blood count data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /blood-count/{id} [get]
func (h *Handler) GetBloodCountById(ctx *gin.Context) {
	id := ctx.Param(userContext)

	bloodCount, err := h.services.BloodCount.GetBloodCountById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, bloodCount)
}

// UpdateBloodCount godoc
// @Summary Update blood count
// @Description Updates an existing blood count entry.
// @Tags BloodCount
// @Accept json
// @Produce json
// @Param input body model.BloodCount true "Updated blood count data"
// @Success 200 {object} string "Updated blood count data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /blood-count [put]
func (h *Handler) UpdateBloodCount(ctx *gin.Context) {
	var bloodCount model.BloodCount

	if err := ctx.BindJSON(&bloodCount); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.BloodCount.UpdateBloodCount(bloodCount)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "updated")
}

// DeleteBloodCount godoc
// @Summary Delete blood count
// @Description Deletes a blood count entry by its ID.
// @Tags BloodCount
// @Produce json
// @Param id path string true "Blood count ID"
// @Success 200 {string} string "Deleted blood count data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /blood-count/{id} [delete]
func (h *Handler) DeleteBloodCount(ctx *gin.Context) {
	id := ctx.Param(userContext)

	err := h.services.BloodCount.DeleteBloodCount(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "deleted")
}
