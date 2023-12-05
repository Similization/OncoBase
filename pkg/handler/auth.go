package handler

import (
	server "med"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LogIn(ctx *gin.Context) {
}

func (h *Handler) Registry(ctx *gin.Context) {
	var input server.User

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) LogOut(ctx *gin.Context) {
	// logout user
}

func (h *Handler) ResetPassword(ctx *gin.Context) {
	// send message to mail
}
