package handler

import (
	"triadmoko-be-golang/auth"
	"triadmoko-be-golang/service"
)

type handler struct {
	userService service.Service
	authService auth.Service
}

func NewHandlerUser(userService service.Service, authService auth.Service) *handler {
	return &handler{userService, authService}
}
