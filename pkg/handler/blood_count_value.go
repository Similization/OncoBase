package handler

import (
	"med/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BloodCountValueResponse struct {
	DiseaseId    string `json:"disease_id"`
	BloodCountId string `json:"blood_count_id"`
}

// CreateBloodCountValue godoc
// @Summary Create blood count value
// @Description Creates a new blood count value entry.
// @Tags BloodCountValue
// @Accept json
// @Produce json
// @Param input body model.BloodCountValue true "Blood count value data"
// @Success 200 {object} model.BloodCountValue "Created blood count value"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /blood-count-value [post]
func (h *Handler) CreateBloodCountValue(ctx *gin.Context) {
	var bloodCountValue model.BloodCountValue

	if err := ctx.BindJSON(&bloodCountValue); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	createdBloodCountValue, err := h.services.BloodCountValue.CreateBloodCountValue(bloodCountValue)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, createdBloodCountValue)
}

// GetBloodCountValueList godoc
// @Summary Get blood count value list
// @Description Retrieves a list of all blood count values.
// @Tags BloodCountValue
// @Produce json
// @Success 200 {array} []model.BloodCountValue "Blood count value list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /blood-count-value [get]
func (h *Handler) GetBloodCountValueList(ctx *gin.Context) {
	bloodCountValueList, err := h.services.BloodCountValue.GetBloodCountValueList()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, bloodCountValueList)
}

// GetBloodCountValueListByDisease godoc
// @Summary Get blood count value list by disease
// @Description Retrieves a list of blood count values associated with a specific disease.
// @Tags BloodCountValue
// @Param disease_id path string true "Disease ID"
// @Produce json
// @Success 200 {array} []model.BloodCountValue "Blood count value list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /blood-count-value/disease/{disease_id} [get]
func (h *Handler) GetBloodCountValueListByDisease(ctx *gin.Context) {
	diseaseId := ctx.Param(diseaseContext)

	bloodCountValue, err := h.services.BloodCountValue.GetBloodCountValueListByDisease(diseaseId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, bloodCountValue)
}

// GetBloodCountValueListByBloodCount godoc
// @Summary Get blood count value list by blood count
// @Description Retrieves a list of blood count values associated with a specific blood count.
// @Tags BloodCountValue
// @Param blood_count_id path string true "Blood count ID"
// @Produce json
// @Success 200 {array} []model.BloodCountValue "Blood count value list"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /blood-count-value/blood-count/{blood_count_id} [get]
func (h *Handler) GetBloodCountValueListByBloodCount(ctx *gin.Context) {
	bloodCountId := ctx.Param(bloodCountContext)

	bloodCountValue, err := h.services.BloodCountValue.GetBloodCountValueListByBloodCount(bloodCountId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, bloodCountValue)
}

// GetBloodCountValueById godoc
// @Summary Get blood count value by ID
// @Description Retrieves a blood count value by its ID.
// @Tags BloodCountValue
// @Param disease_id path string true "Disease ID"
// @Param blood_count_id path string true "Blood count ID"
// @Produce json
// @Success 200 {object} model.BloodCountValue "Blood count value"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /blood-count-value/{disease_id}/{blood_count_id} [get]
func (h *Handler) GetBloodCountValueById(ctx *gin.Context) {
	diseaseId := ctx.Param(diseaseContext)
	bloodCountId := ctx.Param(bloodCountContext)

	bloodCountValue, err := h.services.BloodCountValue.GetBloodCountValueById(diseaseId, bloodCountId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, bloodCountValue)
}

// UpdateBloodCountValue godoc
// @Summary Update blood count value
// @Description Updates an existing blood count value entry.
// @Tags BloodCountValue
// @Accept json
// @Produce json
// @Param input body model.BloodCountValue true "Blood count value data"
// @Success 200 {object} model.BloodCountValue "Updated blood count value"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /blood-count-value [put]
func (h *Handler) UpdateBloodCountValue(ctx *gin.Context) {
	var bloodCountValue model.BloodCountValue

	if err := ctx.BindJSON(&bloodCountValue); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	updatedBloodCountValue, err := h.services.BloodCountValue.UpdateBloodCountValue(bloodCountValue)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, updatedBloodCountValue)
}

// DeleteBloodCountValue godoc
// @Summary Delete blood count value
// @Description Deletes an existing blood count value entry.
// @Tags BloodCountValue
// @Param disease_id path string true "Disease ID"
// @Param blood_count_id path string true "Blood count ID"
// @Produce json
// @Success 200 {object} BloodCountValueResponse "Deleted blood count value"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /blood-count-value/{disease_id}/{blood_count_id} [delete]
func (h *Handler) DeleteBloodCountValue(ctx *gin.Context) {
	diseaseId := ctx.Param(diseaseContext)
	bloodCountId := ctx.Param(bloodCountContext)

	err := h.services.BloodCountValue.DeleteBloodCountValue(diseaseId, bloodCountId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, BloodCountValueResponse{DiseaseId: diseaseId, BloodCountId: bloodCountId})
}
