package service

import (
	"github.com/dulguundd/logError-lib/errs"
	"net/http"
)

type RepositoryGW struct {
	addr string
	port string
}

func (gw G) GetServices() ([]Service, *errs.AppError){
	var services []Service
	var err error

	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")

	return services, nil
}

func NewRepositorGW(addr string, port string) RepositoryGW {
	return RepositoryGW{addr: addr, port: port}
}