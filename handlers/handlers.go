package handlers

import (
	"device-service/services"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type DeviceHandler struct {
	services services.IDeviceService
}

func NewDeviceHandler(services services.IDeviceService) *DeviceHandler {
	return &DeviceHandler{services: services}
}

func (h *DeviceHandler) AddDevice(w http.ResponseWriter, r *http.Request) {
	var deviceData struct {
		Name  string `json:"device_name"`
		Brand string `json:"device_brand"`
	}

	if err := json.NewDecoder(r.Body).Decode(&deviceData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	device, err := h.services.AddDevice(deviceData.Name, deviceData.Brand)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(device)
}

func (h *DeviceHandler) GetDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	device, err := h.services.GetDevice(id)
	if err != nil {
		http.Error(w, "Device not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(device)
}

func (h *DeviceHandler) ListDevices(w http.ResponseWriter, r *http.Request) {
	devices, err := h.services.ListDevices()
	if err != nil {
		http.Error(w, "Unable to list devices", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(devices)
}

func (h *DeviceHandler) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := h.services.DeleteDevice(id); err != nil {
		http.Error(w, "Device not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *DeviceHandler) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updateData struct {
		Name  *string `json:"device_name"`
		Brand *string `json:"device_brand"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	device, err := h.services.UpdateDevice(id, updateData.Name, updateData.Brand)
	if err != nil {
		http.Error(w, "Device not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(device)
}

func (h *DeviceHandler) SearchDeviceByBrand(w http.ResponseWriter, r *http.Request) {
	brand := r.URL.Query().Get("brand")

	devices, err := h.services.SearchDeviceByBrand(brand)
	if err != nil {
		http.Error(w, "Error searching devices", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(devices)
}
