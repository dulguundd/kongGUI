package appservice

import (
	"github.com/dulguundd/logError-lib/errs"
	"kongGUI/domain/service"
)

type DefaultServiceService struct {
	repo service.ServiceRepository
}

func (d DefaultServiceService) GetServices() (service.Service, *errs.AppError) {
	return service.GetServices()
}

type ServiceService interface {
	GetServices() (service.Service, *errs.AppError)
	//GetService(string) (*dto.DeviceResponse, *errs.AppError)
}

func NewServiceService(repository service.ServiceRepository) DefaultServiceService {
	return DefaultServiceService{(repository)}
}
