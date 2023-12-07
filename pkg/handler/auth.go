package handler

import (
	"med/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LogIn(ctx *gin.Context) {
	var input model.AuthUser

	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) Registry(ctx *gin.Context) {
	var user model.User

	if err := ctx.BindJSON(&user); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	email, err := h.services.Authorization.CreateUser(user)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"email": email,
	})
}

func (h *Handler) LogOut(ctx *gin.Context) {
	// logout user
}

func (h *Handler) ResetPassword(ctx *gin.Context) {
	// send message to mail
}
