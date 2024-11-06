package repository

import (
	"device-service/models"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
)

type DeviceRepository struct {
	mu      sync.RWMutex
	devices map[string]models.Device
}

type IDeviceRepository interface {
	AddDevice(name, brand string) (models.Device, error)
	GetDevice(id string) (models.Device, error)
	ListDevices() []models.Device
	UpdateDevice(id string, name, brand *string) (models.Device, error)
	DeleteDevice(id string) error
}

func NewDeviceRepository() *DeviceRepository {
	return &DeviceRepository{
		devices: make(map[string]models.Device),
	}
}

func (r *DeviceRepository) AddDevice(name, brand string) (models.Device, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// TODO: Check if device already exists

	createdAt := time.Now()
	device := models.Device{
		ID:          uuid.New().String(),
		DeviceName:  name,
		DeviceBrand: brand,
		CreatedAt:   createdAt,
	}

	r.devices[device.ID] = device
	return device, nil
}

func (r *DeviceRepository) GetDevice(id string) (models.Device, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	device, exists := r.devices[id]
	if !exists {
		return models.Device{}, errors.New("device not found")
	}
	return device, nil
}

func (r *DeviceRepository) ListDevices() []models.Device {
	r.mu.RLock()
	defer r.mu.RUnlock()

	devices := make([]models.Device, 0, len(r.devices))
	for _, device := range r.devices {
		devices = append(devices, device)
	}
	return devices
}

func (r *DeviceRepository) UpdateDevice(id string, name, brand *string) (models.Device, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	device, exists := r.devices[id]
	if !exists {
		return models.Device{}, errors.New("device not found")
	}

	if name != nil {
		device.DeviceName = *name
	}
	if brand != nil {
		device.DeviceBrand = *brand
	}

	r.devices[id] = device
	return device, nil
}

func (r *DeviceRepository) DeleteDevice(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.devices[id]; !exists {
		return errors.New("device not found")
	}
	delete(r.devices, id)
	return nil
}
