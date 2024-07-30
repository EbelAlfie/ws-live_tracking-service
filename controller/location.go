package controller

import (
	"live-track/main/domain"
	"live-track/main/service"
	"net/http"
)

type locationController struct {
	repo domain.LocationRepository
}

func NewLocationController() domain.LocationController {
	return locationController{
		repo: service.NewLocationRepository(),
	}
}

func (controller locationController) Publish(writer http.ResponseWriter, request http.Request) {
	controller.repo.Publish()
}
