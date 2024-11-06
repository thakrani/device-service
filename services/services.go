package services

import (
	"device-service/models"
	"device-service/repository"
)

type DeviceService struct {
	repo repository.IDeviceRepository
}

type IDeviceService interface {
	AddDevice(name, brand string) (models.Device, error)
	GetDevice(id string) (models.Device, error)
	ListDevices() ([]models.Device, error)
	DeleteDevice(id string) error
	UpdateDevice(id string, name *string, brand *string) (models.Device, error)
	SearchDeviceByBrand(brand string) ([]models.Device, error)
}

func NewDeviceService(repo repository.IDeviceRepository) IDeviceService {
	return &DeviceService{repo: repo}
}

func (s *DeviceService) AddDevice(name, brand string) (models.Device, error) {
	return s.repo.AddDevice(name, brand)
}

func (s *DeviceService) GetDevice(id string) (models.Device, error) {
	return s.repo.GetDevice(id)
}

func (s *DeviceService) ListDevices() ([]models.Device, error) {
	return s.repo.ListDevices(), nil
}

func (s *DeviceService) DeleteDevice(id string) error {
	return s.repo.DeleteDevice(id)
}

func (s *DeviceService) UpdateDevice(id string, name *string, brand *string) (models.Device, error) {
	return s.repo.UpdateDevice(id, name, brand)
}

func (s *DeviceService) SearchDeviceByBrand(brand string) ([]models.Device, error) {
	var results []models.Device
	for _, device := range s.repo.ListDevices() {
		if device.DeviceBrand == brand {
			results = append(results, device)
		}
	}
	return results, nil
}
