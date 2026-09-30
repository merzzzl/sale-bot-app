package controller

import (
	"github.com/ltbots/backend/internal/service"
	"github.com/ltbots/protocols/go/api"
)

type ControllerService struct {
	api.UnimplementedControllerWebServer

	botToken string
	service  *service.Service
}

func NewController(service *service.Service, botToken string) *ControllerService {
	return &ControllerService{
		service:  service,
		botToken: botToken,
	}
}
