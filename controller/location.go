package controller

import (
	"live-track/main/domain"
	"live-track/main/service"
)

type locationController struct {
	repo domain.LocationRepository
}

func NewLocationController() domain.LocationController {
	return locationController{
		repo: service.NewLocationRepository(),
	}
}

func (controller locationController) Publish() {
	controller.repo.Publish()
}
