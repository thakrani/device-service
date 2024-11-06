// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	models "device-service/models"

	mock "github.com/stretchr/testify/mock"
)

// IDeviceRepository is an autogenerated mock type for the IDeviceRepository type
type IDeviceRepository struct {
	mock.Mock
}

// AddDevice provides a mock function with given fields: name, brand
func (_m *IDeviceRepository) AddDevice(name string, brand string) (models.Device, error) {
	ret := _m.Called(name, brand)

	if len(ret) == 0 {
		panic("no return value specified for AddDevice")
	}

	var r0 models.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (models.Device, error)); ok {
		return rf(name, brand)
	}
	if rf, ok := ret.Get(0).(func(string, string) models.Device); ok {
		r0 = rf(name, brand)
	} else {
		r0 = ret.Get(0).(models.Device)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, brand)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteDevice provides a mock function with given fields: id
func (_m *IDeviceRepository) DeleteDevice(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteDevice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDevice provides a mock function with given fields: id
func (_m *IDeviceRepository) GetDevice(id string) (models.Device, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetDevice")
	}

	var r0 models.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.Device, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.Device); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Device)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDevices provides a mock function with given fields:
func (_m *IDeviceRepository) ListDevices() []models.Device {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ListDevices")
	}

	var r0 []models.Device
	if rf, ok := ret.Get(0).(func() []models.Device); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Device)
		}
	}

	return r0
}

// UpdateDevice provides a mock function with given fields: id, name, brand
func (_m *IDeviceRepository) UpdateDevice(id string, name *string, brand *string) (models.Device, error) {
	ret := _m.Called(id, name, brand)

	if len(ret) == 0 {
		panic("no return value specified for UpdateDevice")
	}

	var r0 models.Device
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *string, *string) (models.Device, error)); ok {
		return rf(id, name, brand)
	}
	if rf, ok := ret.Get(0).(func(string, *string, *string) models.Device); ok {
		r0 = rf(id, name, brand)
	} else {
		r0 = ret.Get(0).(models.Device)
	}

	if rf, ok := ret.Get(1).(func(string, *string, *string) error); ok {
		r1 = rf(id, name, brand)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIDeviceRepository creates a new instance of IDeviceRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIDeviceRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IDeviceRepository {
	mock := &IDeviceRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
