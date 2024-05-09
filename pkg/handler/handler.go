package handler

import (
	services "med/pkg/service"
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
