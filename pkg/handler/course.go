package handler

import (
	"med/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCourse godoc
// @Summary Create new course
// @Description Creates a new course.
// @Tags Courses
// @Accept json
// @Produce json
// @Param course body model.Course true "Course data"
// @Success 200 {object} string "Created course data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /courses [post]
func (h *Handler) CreateCourse(ctx *gin.Context) {
	var course model.Course

	if err := ctx.BindJSON(&course); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Course.CreateCourse(course)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "created")
}

// GetCourseList godoc
// @Summary Get list of courses
// @Description Returns a list of courses.
// @Tags Courses
// @Produce json
// @Success 200 {object} []model.Course "List of courses"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /courses [get]
func (h *Handler) GetCourseList(ctx *gin.Context) {
	courseList, err := h.services.Course.GetCourseList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, courseList)
}

// GetCourseById godoc
// @Summary Get course by ID
// @Description Returns a course by its ID.
// @Tags Courses
// @Produce json
// @Param id path string true "Course ID"
// @Success 200 {object} model.Course "Course details"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /courses/{id} [get]
func (h *Handler) GetCourseById(ctx *gin.Context) {
	id := ctx.Param(userContext)

	course, err := h.services.Course.GetCourseById(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, course)
}

// UpdateCourse godoc
// @Summary Update course
// @Description Updates an existing course.
// @Tags Courses
// @Accept json
// @Produce json
// @Param course body model.Course true "Updated course data"
// @Success 200 {object} string "Updated course data"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /courses [put]
func (h *Handler) UpdateCourse(ctx *gin.Context) {
	var course model.Course

	if err := ctx.BindJSON(&course); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Course.UpdateCourse(course)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "updated")
}

// DeleteCourse godoc
// @Summary Delete course
// @Description Deletes an existing course by its ID.
// @Tags Courses
// @Produce json
// @Param id path string true "Course ID"
// @Success 200 {object} string "Deleted course data"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /courses/{id} [delete]
func (h *Handler) DeleteCourse(ctx *gin.Context) {
	id := ctx.Param(userContext)

	err := h.services.Course.DeleteCourse(id)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, "deleted")
}
