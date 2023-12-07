package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userContext         = "email"
)

func (h *Handler) UserIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty header")
		return
	}

	headerParse := strings.Split(header, " ")
	if len(headerParse) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "wrong header")
		return
	}

	token := headerParse[1]
	email, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
	}

	ctx.Set(userContext, email)

}
