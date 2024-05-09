package handler

import (
	"med/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LogIn godoc
// @Summary Log in user
// @Description Logs in the user and returns an authentication token.
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body model.AuthUser true "User credentials"
// @Success 200 {object} string "Authentication token"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /auth/login [post]
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

// Registry godoc
// @Summary Register new user
// @Description Registers a new user and returns the user's email.
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body model.User true "User data"
// @Success 200 {object} string "User email"
// @Failure 400 {object} ErrorResponse "Bad request"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /auth/registry [post]
func (h *Handler) Registry(ctx *gin.Context) {
	var user model.User

	if err := ctx.BindJSON(&user); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "Invalid input body")
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

// LogOut godoc
// @Summary Log out user
// @Description Logs out the currently authenticated user.
// @Tags Auth
// @Produce json
// @Success 200 {string} string "OK"
// @Router /auth/logout [post]
func (h *Handler) LogOut(ctx *gin.Context) {
	// logout user
}

// ResetPassword godoc
// @Summary Reset user password
// @Description Sends a password reset message to the user's email.
// @Tags Auth
// @Produce json
// @Success 200 {string} string "OK"
// @Router /auth/reset-password [post]
func (h *Handler) ResetPassword(ctx *gin.Context) {
	// send message to mail
}
