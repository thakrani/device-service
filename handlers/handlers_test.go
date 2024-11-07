package handlers

import (
	"device-service/models"
	"device-service/services/mocks"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddDevice(t *testing.T) {
	mockService := new(mocks.IDeviceService)
	handler := NewDeviceHandler(mockService)

	device := models.Device{ID: "1", DeviceName: "iPhone", DeviceBrand: "Apple"}
	mockService.On("AddDevice", "iPhone", "Apple").Return(device, nil)

	reqBody := `{"device_name": "iPhone", "device_brand": "Apple"}`
	req := httptest.NewRequest(http.MethodPost, "/devices", nil)
	req.Body = ioutil.NopCloser(strings.NewReader(reqBody))

	rr := httptest.NewRecorder()
	handler.AddDevice(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	var responseDevice models.Device
	json.NewDecoder(rr.Body).Decode(&responseDevice)
	assert.Equal(t, device, responseDevice)
	mockService.AssertExpectations(t)
}

func TestGetDevice(t *testing.T) {
	mockService := new(mocks.IDeviceService)
	handler := NewDeviceHandler(mockService)

	device := models.Device{ID: "1", DeviceName: "iPhone", DeviceBrand: "Apple"}
	mockService.On("GetDevice", "1").Return(device, nil)

	req := httptest.NewRequest(http.MethodGet, "/devices/1", nil)

	rr := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/devices/{id}", handler.GetDevice)

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var responseDevice models.Device
	json.NewDecoder(rr.Body).Decode(&responseDevice)

	assert.Equal(t, device, responseDevice)

	mockService.AssertExpectations(t)
}

func TestListDevices(t *testing.T) {
	mockService := new(mocks.IDeviceService)
	handler := NewDeviceHandler(mockService)

	devices := []models.Device{
		{ID: "1", DeviceName: "iPhone", DeviceBrand: "Apple"},
		{ID: "2", DeviceName: "Galaxy", DeviceBrand: "Samsung"},
	}
	mockService.On("ListDevices").Return(devices, nil)

	req := httptest.NewRequest(http.MethodGet, "/devices", nil)
	rr := httptest.NewRecorder()
	handler.ListDevices(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var responseDevices []models.Device
	json.NewDecoder(rr.Body).Decode(&responseDevices)
	assert.Equal(t, devices, responseDevices)
	mockService.AssertExpectations(t)
}

func TestDeleteDevice(t *testing.T) {
	mockService := new(mocks.IDeviceService)
	handler := NewDeviceHandler(mockService)

	mockService.On("DeleteDevice", "1").Return(nil)

	router := mux.NewRouter()
	router.HandleFunc("/devices/{id}", handler.DeleteDevice).Methods(http.MethodDelete)

	req := httptest.NewRequest(http.MethodDelete, "/devices/1", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)
	mockService.AssertExpectations(t)
}

func TestUpdateDevice(t *testing.T) {
	mockService := new(mocks.IDeviceService)
	handler := NewDeviceHandler(mockService)

	updatedDevice := models.Device{ID: "1", DeviceName: "iPhone 14", DeviceBrand: "Apple"}
	mockService.On("UpdateDevice", "1", mock.AnythingOfType("*string"), mock.AnythingOfType("*string")).Return(updatedDevice, nil)

	reqBody := `{"device_name": "iPhone 14", "device_brand": "Apple"}`
	req := httptest.NewRequest(http.MethodPut, "/devices/1", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	router := mux.NewRouter()
	router.HandleFunc("/devices/{id}", handler.UpdateDevice).Methods(http.MethodPut)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var responseDevice models.Device
	err := json.NewDecoder(rr.Body).Decode(&responseDevice)
	assert.NoError(t, err)
	assert.Equal(t, updatedDevice, responseDevice)
	mockService.AssertExpectations(t)
}

func TestSearchDeviceByBrand(t *testing.T) {
	mockService := new(mocks.IDeviceService)
	handler := NewDeviceHandler(mockService)

	devices := []models.Device{
		{ID: "1", DeviceName: "iPhone", DeviceBrand: "Apple"},
	}
	mockService.On("SearchDeviceByBrand", "Apple").Return(devices, nil)

	req := httptest.NewRequest(http.MethodGet, "/devices/search/Apple", nil)
	req.Header.Set("Content-Type", "application/json")

	router := mux.NewRouter()
	router.HandleFunc("/devices/search/{brand}", handler.SearchDeviceByBrand).Methods(http.MethodGet)

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var responseDevices []models.Device
	err := json.NewDecoder(rr.Body).Decode(&responseDevices)
	assert.NoError(t, err)
	assert.Equal(t, devices, responseDevices)
	mockService.AssertExpectations(t)
}
