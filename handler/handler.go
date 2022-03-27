package handler

import (
	"triadmoko-be-golang/auth"
	"triadmoko-be-golang/service"
)

type handler struct {
	service     service.Service
	authService auth.Service
}

func NewHandlerUser(service service.Service, authService auth.Service) *handler {
	return &handler{service, authService}
}
