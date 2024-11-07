package repository

import (
	"database/sql"
	"device-service/models"
	"errors"
	"time"

	"github.com/google/uuid"
)

type DeviceRepository struct {
	DB *sql.DB
}

type IDeviceRepository interface {
	AddDevice(name, brand string) (models.Device, error)
	GetDevice(id string) (models.Device, error)
	ListDevices() ([]models.Device, error)
	UpdateDevice(id string, name, brand *string) (models.Device, error)
	DeleteDevice(id string) error
	SearchDeviceByBrand(brand string) ([]models.Device, error)
}

func NewDeviceRepository(db *sql.DB) IDeviceRepository {
	return &DeviceRepository{DB: db}
}

func (r *DeviceRepository) AddDevice(name, brand string) (models.Device, error) {
	createdAt := time.Now()
	id := uuid.New().String()
	device := models.Device{
		ID:          id,
		DeviceName:  name,
		DeviceBrand: brand,
		CreatedAt:   createdAt,
	}

	_, err := r.DB.Exec(
		"INSERT INTO devices (id, device_name, device_brand, created_at) VALUES ($1, $2, $3, $4)",
		id, name, brand, createdAt,
	)
	if err != nil {
		return models.Device{}, err
	}

	return device, nil
}

func (r *DeviceRepository) GetDevice(id string) (models.Device, error) {
	var device models.Device
	err := r.DB.QueryRow("SELECT id, device_name, device_brand, created_at FROM devices WHERE id = $1", id).
		Scan(&device.ID, &device.DeviceName, &device.DeviceBrand, &device.CreatedAt)

	if err == sql.ErrNoRows {
		return models.Device{}, errors.New("device not found")
	} else if err != nil {
		return models.Device{}, err
	}

	return device, nil
}

func (r *DeviceRepository) ListDevices() ([]models.Device, error) {
	rows, err := r.DB.Query("SELECT id, device_name, device_brand, created_at FROM devices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	devices := []models.Device{}
	for rows.Next() {
		var device models.Device
		if err := rows.Scan(&device.ID, &device.DeviceName, &device.DeviceBrand, &device.CreatedAt); err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}

	return devices, nil
}

func (r *DeviceRepository) UpdateDevice(id string, name, brand *string) (models.Device, error) {
	if name != nil {
		_, err := r.DB.Exec("UPDATE devices SET device_name = $1 WHERE id = $2", *name, id)
		if err != nil {
			return models.Device{}, err
		}
	}
	if brand != nil {
		_, err := r.DB.Exec("UPDATE devices SET device_brand = $1 WHERE id = $2", *brand, id)
		if err != nil {
			return models.Device{}, err
		}
	}
	return r.GetDevice(id)
}

func (r *DeviceRepository) DeleteDevice(id string) error {
	_, err := r.DB.Exec("DELETE FROM devices WHERE id = $1", id)
	return err
}

func (r *DeviceRepository) SearchDeviceByBrand(brand string) ([]models.Device, error) {
	rows, err := r.DB.Query("SELECT id, device_name, device_brand, created_at FROM  devices WHERE device_brand = $1", brand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	devices := []models.Device{}
	for rows.Next() {
		var device models.Device
		if err := rows.Scan(&device.ID, &device.DeviceName, &device.DeviceBrand, &device.CreatedAt); err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}

	return devices, nil
}
