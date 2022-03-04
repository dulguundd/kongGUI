package app

import (
	"github.com/gorilla/mux"
	"kongGUI/domain/service"
	"log"
	"net/http"
)

func Start()  {
	//serviceConfig := sanityCheckService()

	router := mux.NewRouter()

	// wiring
	gwConnectionAddr, gwConnectionPort := getGWConnection()
	ServiceRepositoryGW := service.NewRepositorGW(gwConnectionAddr, gwConnectionPort)

	seh := serviceHandlers{service.NewServiceService(ServiceRepositoryGW)}

	//define device routes
	router.HandleFunc("/service", seh.getServices).Methods(http.MethodGet)

	//starting server
	serviceConfig.address = "0.0.0.0"
	serviceConfig.port = "8001"
	log.Fatal(http.ListenAndServe(serviceConfig.address+":"+serviceConfig.port, router))
}
