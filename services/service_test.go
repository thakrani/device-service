package services_test

import (
	"device-service/mocks"
	"device-service/models"
	"device-service/services"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddDevice(t *testing.T) {
	mockRepo := new(mocks.IDeviceRepository)
	deviceService := services.NewDeviceService(mockRepo)

	createdAt, _ := time.Parse("2006-01-02", "2024-11-07")
	device := models.Device{ID: "1", DeviceName: "Phone", DeviceBrand: "BrandX", CreatedAt: createdAt}
	mockRepo.On("AddDevice", "Phone", "BrandX").Return(device, nil)

	result, err := deviceService.AddDevice("Phone", "BrandX")

	assert.NoError(t, err)
	assert.Equal(t, device, result)
	mockRepo.AssertExpectations(t)
}

func TestGetDevice(t *testing.T) {
	mockRepo := new(mocks.IDeviceRepository)
	deviceService := services.NewDeviceService(mockRepo)

	createdAt, _ := time.Parse("2006-01-02", "2024-11-07")
	device := models.Device{ID: "1", DeviceName: "Phone", DeviceBrand: "BrandX", CreatedAt: createdAt}
	mockRepo.On("GetDevice", "1").Return(device, nil)

	result, err := deviceService.GetDevice("1")

	assert.NoError(t, err)
	assert.Equal(t, device, result)
	mockRepo.AssertExpectations(t)
}

func TestListDevices(t *testing.T) {
	mockRepo := new(mocks.IDeviceRepository)
	deviceService := services.NewDeviceService(mockRepo)

	createdAt, _ := time.Parse("2006-01-02", "2024-11-07")
	devices := []models.Device{
		{ID: "1", DeviceName: "Phone", DeviceBrand: "BrandX", CreatedAt: createdAt},
	}
	mockRepo.On("ListDevices").Return(devices, nil)

	result, err := deviceService.ListDevices()

	assert.NoError(t, err)
	assert.Equal(t, devices, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteDevice(t *testing.T) {
	mockRepo := new(mocks.IDeviceRepository)
	deviceService := services.NewDeviceService(mockRepo)

	mockRepo.On("DeleteDevice", "1").Return(nil)

	err := deviceService.DeleteDevice("1")

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateDevice(t *testing.T) {
	mockRepo := new(mocks.IDeviceRepository)
	deviceService := services.NewDeviceService(mockRepo)

	createdAt, _ := time.Parse("2006-01-02", "2024-11-07")
	device := models.Device{ID: "1", DeviceName: "UpdatedPhone", DeviceBrand: "BrandX", CreatedAt: createdAt}
	mockRepo.On("UpdateDevice", "1", mock.Anything, mock.Anything).Return(device, nil)

	result, err := deviceService.UpdateDevice("1", nil, nil)

	assert.NoError(t, err)
	assert.Equal(t, device, result)
	mockRepo.AssertExpectations(t)
}

func TestSearchDeviceByBrand(t *testing.T) {
	mockRepo := new(mocks.IDeviceRepository)
	deviceService := services.NewDeviceService(mockRepo)

	createdAt, _ := time.Parse("2006-01-02", "2024-11-07")
	devices := []models.Device{
		{ID: "1", DeviceName: "Phone", DeviceBrand: "BrandX", CreatedAt: createdAt},
	}
	mockRepo.On("SearchDeviceByBrand", "BrandX").Return(devices, nil)

	result, err := deviceService.SearchDeviceByBrand("BrandX")

	assert.NoError(t, err)
	assert.Equal(t, devices, result)
	mockRepo.AssertExpectations(t)
}
