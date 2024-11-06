package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDevice(t *testing.T) {
	repo := NewDeviceRepository()
	device, err := repo.AddDevice("iPhone", "Apple")
	assert.NoError(t, err)
	assert.NotEmpty(t, device.ID)
	assert.Equal(t, "iPhone", device.DeviceName)
	assert.Equal(t, "Apple", device.DeviceBrand)
	assert.NotNil(t, device.CreatedAt)
}

func TestGetDevice(t *testing.T) {
	repo := NewDeviceRepository()
	device, _ := repo.AddDevice("iPhone", "Apple")
	fetchedDevice, err := repo.GetDevice(device.ID)
	assert.NoError(t, err)
	assert.Equal(t, device.ID, fetchedDevice.ID)
	assert.Equal(t, device.DeviceName, fetchedDevice.DeviceName)
	assert.Equal(t, device.DeviceBrand, fetchedDevice.DeviceBrand)
}

func TestGetDevice_NotFound(t *testing.T) {
	repo := NewDeviceRepository()
	_, err := repo.GetDevice("non-existent-id")
	assert.Error(t, err)
}

func TestListDevices(t *testing.T) {
	repo := NewDeviceRepository()
	repo.AddDevice("iPhone", "Apple")
	repo.AddDevice("Galaxy", "Samsung")
	devices := repo.ListDevices()
	assert.Len(t, devices, 2)
}

func TestUpdateDevice(t *testing.T) {
	repo := NewDeviceRepository()
	device, _ := repo.AddDevice("iPhone", "Apple")
	newName := "iPhone 15"
	newBrand := "Apple"
	updatedDevice, err := repo.UpdateDevice(device.ID, &newName, &newBrand)
	assert.NoError(t, err)
	assert.Equal(t, newName, updatedDevice.DeviceName)
	assert.Equal(t, newBrand, updatedDevice.DeviceBrand)
}

func TestUpdateDevice_NotFound(t *testing.T) {
	repo := NewDeviceRepository()
	newName := "iPhone 15"
	newBrand := "Apple"
	_, err := repo.UpdateDevice("non-existent-id", &newName, &newBrand)
	assert.Error(t, err)
}

func TestDeleteDevice(t *testing.T) {
	repo := NewDeviceRepository()
	device, _ := repo.AddDevice("iPhone", "Apple")
	err := repo.DeleteDevice(device.ID)
	assert.NoError(t, err)
	_, err = repo.GetDevice(device.ID)
	assert.Error(t, err)
}

func TestDeleteDevice_NotFound(t *testing.T) {
	repo := NewDeviceRepository()
	err := repo.DeleteDevice("non-existent-id")
	assert.Error(t, err)
}
