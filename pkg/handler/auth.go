package handler

import "github.com/gin-gonic/gin"

func (h *Handler) LogIn(c *gin.Context) {
	// check in db for user

}

func (h *Handler) Registry(c *gin.Context) {
	// create user in db
}

func (h *Handler) LogOut(c *gin.Context) {
	// logout user
}

func (h *Handler) ResetPassword(c *gin.Context) {
	// send message to mail
}
