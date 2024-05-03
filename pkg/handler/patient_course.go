package handler

import (
	"med/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatePatientCourse godoc
// @Summary Create patient course
// @Description Creates a new patient course.
// @Tags PatientCourse
// @Accept json
// @Produce json
// @Param input body model.PatientCourse true "Patient course data"
// @Success 200 {object} model.PatientCourse "Created patient course data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patient-courses [post]
func (h *Handler) CreatePatientCourse(ctx *gin.Context) {
	var patientCourse model.PatientCourse

	if err := ctx.BindJSON(&patientCourse); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	createdPatientCourse, err := h.services.PatientCourse.CreatePatientCourse(patientCourse)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, createdPatientCourse)
}

// GetPatientCourseList godoc
// @Summary Get patient course list
// @Description Retrieves a list of patient courses.
// @Tags PatientCourse
// @Produce json
// @Success 200 {array} []model.PatientCourse "Patient course list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patient-courses [get]
func (h *Handler) GetPatientCourseList(ctx *gin.Context) {
	patientCourseList, err := h.services.PatientCourse.GetPatientCourseList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, patientCourseList)
}

// GetPatientCourseById godoc
// @Summary Get patient course by ID
// @Description Retrieves a patient course by ID.
// @Tags PatientCourse
// @Produce json
// @Param id path string true "Patient course ID"
// @Success 200 {object} model.PatientCourse "Patient course data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patient-courses/{id} [get]
func (h *Handler) GetPatientCourseById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param(userContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	patientCourse, err := h.services.PatientCourse.GetPatientCourseById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, patientCourse)
}

// UpdatePatientCourse godoc
// @Summary Update patient course
// @Description Updates an existing patient course.
// @Tags PatientCourse
// @Accept json
// @Produce json
// @Param input body model.PatientCourse true "Patient course data"
// @Success 200 {object} model.PatientCourse "Updated patient course data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patient-courses [put]
func (h *Handler) UpdatePatientCourse(ctx *gin.Context) {
	var patientCourse model.PatientCourse

	if err := ctx.BindJSON(&patientCourse); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	updatedPatientCourse, err := h.services.PatientCourse.UpdatePatientCourse(patientCourse)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, updatedPatientCourse)
}

// DeletePatientCourse godoc
// @Summary Delete patient course
// @Description Deletes a patient course by ID.
// @Tags PatientCourse
// @Produce json
// @Param id path string true "Patient course ID"
// @Success 200 {string} string "Patient course ID deleted"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /patient-courses/{id} [delete]
func (h *Handler) DeletePatientCourse(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param(userContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.PatientCourse.DeletePatientCourse(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)
}
