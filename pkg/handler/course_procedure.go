package handler

import (
	"med/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCourseProcedure godoc
// @Summary Create a new course procedure
// @Description Creates a new course procedure entry.
// @Tags CourseProcedure
// @Accept json
// @Produce json
// @Param input body model.CourseProcedure true "Course procedure data"
// @Success 200 {object} model.CourseProcedure "Created course procedure data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /course-procedure [post]
func (h *Handler) CreateCourseProcedure(ctx *gin.Context) {
	var courseProcedure model.CourseProcedure

	if err := ctx.BindJSON(&courseProcedure); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	createdCourseProcedure, err := h.services.CourseProcedure.CreateCourseProcedure(courseProcedure)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, createdCourseProcedure)
}

// GetCourseProcedureList godoc
// @Summary Get course procedure list
// @Description Retrieves a list of course procedures.
// @Tags CourseProcedure
// @Produce json
// @Success 200 {array} []model.CourseProcedure "Course procedure list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /course-procedure [get]
func (h *Handler) GetCourseProcedureList(ctx *gin.Context) {
	courseProcedureList, err := h.services.CourseProcedure.GetCourseProcedureList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, courseProcedureList)
}

// GetCourseProcedureById godoc
// @Summary Get course procedure by ID
// @Description Retrieves a course procedure entry by its ID.
// @Tags CourseProcedure
// @Produce json
// @Param id path string true "Course procedure ID"
// @Success 200 {object} model.CourseProcedure "Course procedure data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /course-procedure/{id} [get]
func (h *Handler) GetCourseProcedureById(ctx *gin.Context) {
	id := ctx.Param(userContext)

	courseProcedure, err := h.services.CourseProcedure.GetCourseProcedureById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, courseProcedure)
}

// UpdateCourseProcedure godoc
// @Summary Update course procedure
// @Description Updates an existing course procedure entry.
// @Tags CourseProcedure
// @Accept json
// @Produce json
// @Param input body model.CourseProcedure true "Updated course procedure data"
// @Success 200 {object} model.CourseProcedure "Updated course procedure data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /course-procedure [put]
func (h *Handler) UpdateCourseProcedure(ctx *gin.Context) {
	var courseProcedure model.CourseProcedure

	if err := ctx.BindJSON(&courseProcedure); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	updatedCourseProcedure, err := h.services.CourseProcedure.UpdateCourseProcedure(courseProcedure)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, updatedCourseProcedure)
}

// DeleteCourseProcedure godoc
// @Summary Delete course procedure
// @Description Deletes a course procedure entry by its ID.
// @Tags CourseProcedure
// @Produce json
// @Param id path string true "Course procedure ID"
// @Success 200 {string} string "Course procedure ID"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /course-procedure/{id} [delete]
func (h *Handler) DeleteCourseProcedure(ctx *gin.Context) {
	id := ctx.Param(userContext)

	err := h.services.CourseProcedure.DeleteCourseProcedure(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)
}
