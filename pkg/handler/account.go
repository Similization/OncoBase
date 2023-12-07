package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Settings(ctx *gin.Context) {
	email, _ := ctx.Get("email")
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"email": email,
	})
}

func (h *Handler) BloodCount(ctx *gin.Context) {

}

func (h *Handler) Doctors(ctx *gin.Context) {

}

func (h *Handler) Devs(ctx *gin.Context) {

}

func (h *Handler) AnalysisData(ctx *gin.Context) {

}
