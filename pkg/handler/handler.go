package handler

import (
	"github.com/sculptorvoid/langi_backend/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
