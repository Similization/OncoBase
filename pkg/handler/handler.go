package handler

import (
	"med/pkg/services"
)

type Handler struct {
	services *services.Service
}

func NewHandler(s *services.Service) *Handler {
	return &Handler{services: s}
}
