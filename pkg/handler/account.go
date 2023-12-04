package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Settings(c *gin.Context) {
	c.String(http.StatusOK, "The available groups are [...]")
}

func (h *Handler) BloodCount(c *gin.Context) {

}

func (h *Handler) Doctors(c *gin.Context) {

}

func (h *Handler) Devs(c *gin.Context) {

}

func (h *Handler) AnalysisData(c *gin.Context) {

}
