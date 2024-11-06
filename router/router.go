package router

import (
	"device-service/handlers"
	"device-service/repository"
	"device-service/services"

	"github.com/gorilla/mux"
)

func InitializeRoutes(r *mux.Router) {
	repo := repository.NewDeviceRepository()
	service := services.NewDeviceService(repo)
	handler := handlers.NewDeviceHandler(service)
	r.HandleFunc("/devices", handler.AddDevice).Methods("POST")
	r.HandleFunc("/devices", handler.ListDevices).Methods("GET")
	r.HandleFunc("/devices/{id}", handler.GetDevice).Methods("GET")
	r.HandleFunc("/devices/{id}", handler.UpdateDevice).Methods("PUT")
	r.HandleFunc("/devices/{id}", handler.DeleteDevice).Methods("DELETE")
	r.HandleFunc("/devices/search/{brand}", handler.SearchDeviceByBrand).Methods("GET")
}
