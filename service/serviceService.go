package service

import (
	"github.com/dulguundd/logError-lib/errs"
	"kongGUI/domain/service"
)

type DefaultServiceService struct {
	repo service.ServiceRepository
}

type DeviceService interface {
	GetServices() (service.Service, *errs.AppError)
	//GetService(string) (*dto.DeviceResponse, *errs.AppError)
}

func NewServiceService(repository service.ServiceRepository) DefaultServiceService {
	return DefaultServiceService{(repository)}
}
