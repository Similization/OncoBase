package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// ErrorResponse represents an error response sent to the client.
type ErrorResponse struct {
	Message string `json:"message"`
}

// newErrorResponse logs the error and sends an error response to the client with the provided message and status code.
func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	log.Error().Msg(message)
	ctx.AbortWithStatusJSON(statusCode, ErrorResponse{Message: message})
}
