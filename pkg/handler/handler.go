package handler

import (
	"med/pkg/services"
)

type Handler struct {
	AccountHandler AccountHandlers
	services       *services.Service
}

func NewHandler(s *services.Service) *Handler {
	return &Handler{
		AccountHandler: AccountHandlers{},
		services:       s,
	}
}
