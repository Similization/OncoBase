package handler

import (
	"med/pkg/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateDoctor godoc
// @Summary Create doctor
// @Description Creates a new doctor.
// @Tags Doctor
// @Accept json
// @Produce json
// @Param input body model.Doctor true "Doctor data"
// @Success 200 {object} model.Doctor "Created doctor data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /doctors [post]
func (h *Handler) CreateDoctor(ctx *gin.Context) {
	var doctor model.Doctor

	if err := ctx.BindJSON(&doctor); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	createdDoctor, err := h.services.Doctor.CreateDoctor(doctor)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, createdDoctor)
}

// GetDoctorList godoc
// @Summary Get doctor list
// @Description Retrieves a list of doctors.
// @Tags Doctor
// @Produce json
// @Success 200 {array} []model.Doctor "Doctor list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /doctors [get]
func (h *Handler) GetDoctorList(ctx *gin.Context) {
	doctorList, err := h.services.Doctor.GetDoctorList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, doctorList)
}

// GetDoctorById godoc
// @Summary Get doctor by ID
// @Description Retrieves a doctor by ID.
// @Tags Doctor
// @Produce json
// @Param id path string true "Doctor ID"
// @Success 200 {object} model.Doctor "Doctor data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /doctors/{id} [get]
func (h *Handler) GetDoctorById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param(userContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	doctor, err := h.services.Doctor.GetDoctorById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, doctor)
}

// UpdateDoctor godoc
// @Summary Update doctor
// @Description Updates an existing doctor.
// @Tags Doctor
// @Accept json
// @Produce json
// @Param input body model.Doctor true "Doctor data"
// @Success 200 {object} model.Doctor "Updated doctor data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /doctors [put]
func (h *Handler) UpdateDoctor(ctx *gin.Context) {
	var doctor model.Doctor

	if err := ctx.BindJSON(&doctor); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	updatedDoctor, err := h.services.Doctor.UpdateDoctor(doctor)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, updatedDoctor)
}

// DeleteDoctor godoc
// @Summary Delete doctor
// @Description Deletes a doctor by ID.
// @Tags Doctor
// @Produce json
// @Param id path string true "Doctor ID"
// @Success 200 {string} string "Doctor ID deleted"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /doctors/{id} [delete]
func (h *Handler) DeleteDoctor(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param(userContext))
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Doctor.DeleteDoctor(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, id)
}
