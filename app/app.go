package app

import (
	"github.com/gorilla/mux"
	"kongGUI/appservice"
	"kongGUI/domain/service"
	"log"
	"net/http"
)

func Start() {
	//serviceConfig := sanityCheckService()

	router := mux.NewRouter()

	// wiring
	gwConnectionAddr, gwConnectionPort := getGWConnection()
	ServiceRepositoryGW := service.NewRepositorGW(gwConnectionAddr, gwConnectionPort)

	seh := serviceHandlers{appservice.NewServiceService(ServiceRepositoryGW)}

	//define device routes
	router.HandleFunc("/service", seh.getServices).Methods(http.MethodGet)

	//starting server
	serviceconfig.address = "0.0.0.0"
	serviceconfig.port = "8001"
	log.Fatal(http.ListenAndServe(serviceconfig.address+":"+serviceconfig.port, router))
}
