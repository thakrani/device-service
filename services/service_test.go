package services

import (
	"device-service/models"
	"device-service/repository/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddDevice(t *testing.T) {
	mockRepo := new(mocks.IDeviceRepository)
	service := NewDeviceService(mockRepo)

	device := models.Device{ID: "1", DeviceName: "iPhone", DeviceBrand: "Apple"}

	mockRepo.On("AddDevice", "iPhone", "Apple").Return(device, nil)

	result, err := service.AddDevice("iPhone", "Apple")

	assert.NoError(t, err)
	assert.Equal(t, device, result)
	mockRepo.AssertExpectations(t)
}

func TestGetDevice(t *testing.T) {
	mockRepo := new(mocks.IDeviceRepository)
	service := NewDeviceService(mockRepo)

	device := models.Device{ID: "1", DeviceName: "iPhone", DeviceBrand: "Apple"}

	mockRepo.On("GetDevice", "1").Return(device, nil)

	result, err := service.GetDevice("1")

	assert.NoError(t, err)
	assert.Equal(t, device, result)
	mockRepo.AssertExpectations(t)
}

func TestListDevices(t *testing.T) {
	mockRepo := new(mocks.IDeviceRepository)
	service := NewDeviceService(mockRepo)

	devices := []models.Device{
		{ID: "1", DeviceName: "iPhone", DeviceBrand: "Apple"},
		{ID: "2", DeviceName: "Galaxy", DeviceBrand: "Samsung"},
	}

	mockRepo.On("ListDevices").Return(devices)

	result, err := service.ListDevices()

	assert.NoError(t, err)
	assert.Equal(t, devices, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteDevice(t *testing.T) {
	mockRepo := new(mocks.IDeviceRepository)
	service := NewDeviceService(mockRepo)

	mockRepo.On("DeleteDevice", "1").Return(nil)

	err := service.DeleteDevice("1")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateDevice(t *testing.T) {
	mockRepo := new(mocks.IDeviceRepository)
	service := NewDeviceService(mockRepo)

	updatedDevice := models.Device{ID: "1", DeviceName: "iPhone 14", DeviceBrand: "Apple"}

	mockRepo.On("UpdateDevice", "1", mock.AnythingOfType("*string"), mock.AnythingOfType("*string")).Return(updatedDevice, nil)

	name := "iPhone 14"
	brand := "Apple"
	result, err := service.UpdateDevice("1", &name, &brand)

	assert.NoError(t, err)
	assert.Equal(t, updatedDevice, result)
	mockRepo.AssertExpectations(t)
}

func TestSearchDeviceByBrand(t *testing.T) {
	mockRepo := new(mocks.IDeviceRepository)
	service := NewDeviceService(mockRepo)

	devices := []models.Device{
		{ID: "1", DeviceName: "iPhone", DeviceBrand: "Apple"},
		{ID: "2", DeviceName: "Galaxy", DeviceBrand: "Samsung"},
		{ID: "3", DeviceName: "MacBook", DeviceBrand: "Apple"},
	}

	mockRepo.On("ListDevices").Return(devices)

	results, err := service.SearchDeviceByBrand("Apple")

	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, "Apple", results[0].DeviceBrand)
	assert.Equal(t, "Apple", results[1].DeviceBrand)
	mockRepo.AssertExpectations(t)
}
