package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Error struct {
	Message string
}

func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	log.Error().Msg(message)
	ctx.AbortWithStatusJSON(statusCode, Error{message})
}
