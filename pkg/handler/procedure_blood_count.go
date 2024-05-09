package handler

import (
	"med/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateProcedureBloodCount godoc
// @Summary Create procedure blood count
// @Description Creates a new procedure blood count entry.
// @Tags ProcedureBloodCount
// @Accept json
// @Produce json
// @Param input body model.ProcedureBloodCount true "Procedure blood count data"
// @Success 200 {object} string "Created procedure blood count data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /procedure-blood-count [post]
func (h *Handler) CreateProcedureBloodCount(ctx *gin.Context) {
	var procedureBloodCount model.ProcedureBloodCount

	if err := ctx.BindJSON(&procedureBloodCount); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.ProcedureBloodCount.CreateProcedureBloodCount(procedureBloodCount)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "created")
}

// GetProcedureBloodCountList godoc
// @Summary Get procedure blood count list
// @Description Retrieves a list of procedure blood count entries.
// @Tags ProcedureBloodCount
// @Produce json
// @Success 200 {array} []model.ProcedureBloodCount "Procedure blood count list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /procedure-blood-count [get]
func (h *Handler) GetProcedureBloodCountList(ctx *gin.Context) {
	procedureBloodCountList, err := h.services.ProcedureBloodCount.GetProcedureBloodCountList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, procedureBloodCountList)
}

// GetProcedureBloodCountListByProcedure godoc
// @Summary Get procedure blood count list by procedure
// @Description Retrieves a list of procedure blood count entries by procedure ID.
// @Tags ProcedureBloodCount
// @Produce json
// @Param procedure_id path string true "Procedure ID"
// @Success 200 {array} []model.ProcedureBloodCount "Procedure blood count list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /procedure-blood-count/procedures/{procedure_id} [get]
func (h *Handler) GetProcedureBloodCountListByProcedure(ctx *gin.Context) {
	procedureId, err := strconv.Atoi(ctx.Param(procedureContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	procedureBloodCount, err := h.services.ProcedureBloodCount.GetProcedureBloodCountListByProcedure(procedureId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, procedureBloodCount)
}

// GetProcedureBloodCountListByBloodCount godoc
// @Summary Get procedure blood count list by blood count
// @Description Retrieves a list of procedure blood count entries by blood count ID.
// @Tags ProcedureBloodCount
// @Produce json
// @Param blood_count_id path string true "Blood count ID"
// @Success 200 {array} []model.ProcedureBloodCount "Procedure blood count list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /procedure-blood-count/blood-counts/{blood_count_id} [get]
func (h *Handler) GetProcedureBloodCountListByBloodCount(ctx *gin.Context) {
	bloodCountId := ctx.Param(bloodCountContext)

	procedureBloodCount, err := h.services.ProcedureBloodCount.GetProcedureBloodCountListByBloodCount(bloodCountId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, procedureBloodCount)
}

// GetProcedureBloodCountById godoc
// @Summary Get procedure blood count by IDs
// @Description Retrieves a procedure blood count entry by procedure ID and blood count ID.
// @Tags ProcedureBloodCount
// @Produce json
// @Param procedure_id path string true "Procedure ID"
// @Param blood_count_id path string true "Blood count ID"
// @Success 200 {object} model.ProcedureBloodCount "Procedure blood count data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /procedure-blood-count/procedures/{procedure_id}/blood-counts/{blood_count_id} [get]
func (h *Handler) GetProcedureBloodCountById(ctx *gin.Context) {
	procedureId, err := strconv.Atoi(ctx.Param(procedureContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	bloodCountId := ctx.Param(bloodCountContext)

	procedureBloodCount, err := h.services.ProcedureBloodCount.GetProcedureBloodCountById(procedureId, bloodCountId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, procedureBloodCount)
}

// UpdateProcedureBloodCount godoc
// @Summary Update procedure blood count
// @Description Updates an existing procedure blood count entry.
// @Tags ProcedureBloodCount
// @Accept json
// @Produce json
// @Param input body model.ProcedureBloodCount true "Procedure blood count data"
// @Success 200 {object} string "Updated procedure blood count data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /procedure-blood-count [put]
func (h *Handler) UpdateProcedureBloodCount(ctx *gin.Context) {
	var procedureBloodCount model.ProcedureBloodCount

	if err := ctx.BindJSON(&procedureBloodCount); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.ProcedureBloodCount.UpdateProcedureBloodCount(procedureBloodCount)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "updated")
}

// DeleteProcedureBloodCount godoc
// @Summary Delete procedure blood count
// @Description Deletes a procedure blood count entry by procedure ID and blood count ID.
// @Tags ProcedureBloodCount
// @Produce json
// @Param procedure_id path string true "Procedure ID"
// @Param blood_count_id path string true "Blood count ID"
// @Success 200 {string} string "Deleted procedure blood count data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /procedure-blood-count/procedures/{procedure_id}/blood-counts/{blood_count_id} [delete]
func (h *Handler) DeleteProcedureBloodCount(ctx *gin.Context) {
	procedureId, err := strconv.Atoi(ctx.Param(procedureContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	bloodCountId := ctx.Param(bloodCountContext)

	err = h.services.ProcedureBloodCount.DeleteProcedureBloodCount(procedureId, bloodCountId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "deleted")
}
