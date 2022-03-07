package app

import (
	"kongGUI/appservice"
	"net/http"
	"time"
)

type serviceHandlers struct {
	service appservice.ServiceService
}

func (seh *serviceHandlers) getServices(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	devices, err := seh.service.GetAllDevice()

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, devices)
		serviceLatencyLogger(start)
	}
}
