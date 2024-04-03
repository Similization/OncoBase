package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountHandlers struct {
}

func (h *AccountHandlers) Settings(ctx *gin.Context) {
	id, _ := ctx.Get(userContext)
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *AccountHandlers) BloodCount(ctx *gin.Context) {

}

func (h *AccountHandlers) Doctors(ctx *gin.Context) {

}

func (h *AccountHandlers) PatientData(ctx *gin.Context) {

}

func (h *AccountHandlers) Console(ctx *gin.Context) {

}
