package handler

import (
	"errors"
	"net/http"
	"strings"

	services "med/pkg/service"

	"github.com/gin-gonic/gin"
)

// Constants for context keys and roles
const (
	authorizationHeader = "Authorization"

	userContext       = "id"
	bloodCountContext = "blood_count_id"
	diseaseContext    = "disease_id"
	doctorContext     = "doctor_id"
	patientContext    = "patient_id"
	procedureContext  = "procedure_id"

	adminRole   = "admin"
	patientRole = "patient"
	doctorRole  = "doctor"
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

// getUserData retrieves user data from the authorization header
func (h *Handler) getUserData(ctx *gin.Context) (*services.UserData, error) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return nil, errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return nil, errors.New("malformed header")
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(ctx, http.StatusUnauthorized, "token is empty")
		return nil, errors.New("token is empty")
	}

	token := headerParts[1]
	userData, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return nil, err
	}

	return userData, nil
}

// UserIdentity middleware sets user ID in context
func (h *Handler) UserIdentity(ctx *gin.Context) {
	userData, err := h.getUserData(ctx)
	if err != nil {
		return
	}

	ctx.Set(userContext, userData.Id)
}

// AdminIdentity middleware checks if the user is an admin
func (h *Handler) AdminIdentity(ctx *gin.Context) {
	userData, err := h.getUserData(ctx)
	if err != nil {
		return
	}

	if userData.Role != adminRole {
		newErrorResponse(ctx, http.StatusUnauthorized, "insufficient permissions")
		return
	}
}

// PatientIdentity middleware checks if the user is a patient
func (h *Handler) PatientIdentity(ctx *gin.Context) {
	userData, err := h.getUserData(ctx)
	if err != nil {
		return
	}

	if userData.Role != patientRole {
		newErrorResponse(ctx, http.StatusUnauthorized, "insufficient permissions")
		return
	}
}

// DoctorIdentity middleware checks if the user is a doctor
func (h *Handler) DoctorIdentity(ctx *gin.Context) {
	userData, err := h.getUserData(ctx)
	if err != nil {
		return
	}

	if userData.Role != doctorRole {
		newErrorResponse(ctx, http.StatusUnauthorized, "insufficient permissions")
		return
	}
}
