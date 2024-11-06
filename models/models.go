package models

import "time"

type Device struct {
	ID          string    `json:"id"`           // Unique identifier for the device
	DeviceName  string    `json:"device_name"`  // Name of the device
	DeviceBrand string    `json:"device_brand"` // Brand of the device
	CreatedAt   time.Time `json:"created_at"`   // Timestamp of creation
}
