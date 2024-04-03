package handler

import (
	"errors"
	"med/pkg/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userContext         = "id"
	adminRole           = "admin"
	patinetRole         = "patient"
	doctorRole          = "doctor"
)

// type rolePermissions struct {
// 	Role        string
// 	Permissions []url.URL
// }

// var rolePermissionsMap map[string][]url.URL

// func (h *Handler) checkPermissions(role string, url url.URL) bool {
// 	permissions, ok := rolePermissionsMap[role]
// 	if !ok {
// 		return false
// 	}
// 	for _, permission := range permissions {
// 		if permission == url {
// 			return true
// 		}
// 	}
// 	return false
// }

func (h *Handler) getUserData(ctx *gin.Context) (*services.UserData, error) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty header")
		return nil, errors.New("empty header")
	}

	headerParse := strings.Split(header, " ")
	if len(headerParse) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "wrong header")
		return nil, errors.New("wrong header")
	}

	token := headerParse[1]
	userData, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return nil, err
	}

	return userData, nil
}

func (h *Handler) UserIdentity(ctx *gin.Context) {
	userRole, err := h.getUserData(ctx)
	if err != nil {
		return
	}

	ctx.Set(userContext, userRole.Id)
}

func (h *Handler) AdminIdentity(ctx *gin.Context) {
	userRole, err := h.getUserData(ctx)
	if err != nil {
		return
	}

	if userRole.Role != adminRole {
		newErrorResponse(ctx, http.StatusUnauthorized, "No permissions")
		return
	}
}

func (h *Handler) PatientIdentity(ctx *gin.Context) {
	userRole, err := h.getUserData(ctx)
	if err != nil {
		return
	}

	if userRole.Role != patinetRole {
		newErrorResponse(ctx, http.StatusUnauthorized, "No permissions")
		return
	}
}

func (h *Handler) DoctorIdentity(ctx *gin.Context) {
	userRole, err := h.getUserData(ctx)
	if err != nil {
		return
	}

	if userRole.Role != doctorRole {
		newErrorResponse(ctx, http.StatusUnauthorized, "No permissions")
		return
	}
}
