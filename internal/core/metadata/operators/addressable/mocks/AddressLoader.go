// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import models "github.com/edgexfoundry/go-mod-core-contracts/v2/models"

// AddressLoader is an autogenerated mock type for the AddressLoader type
type AddressLoader struct {
	mock.Mock
}

// GetAddressableById provides a mock function with given fields: id
func (_m *AddressLoader) GetAddressableById(id string) (models.Addressable, error) {
	ret := _m.Called(id)

	var r0 models.Addressable
	if rf, ok := ret.Get(0).(func(string) models.Addressable); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Addressable)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressableByName provides a mock function with given fields: n
func (_m *AddressLoader) GetAddressableByName(n string) (models.Addressable, error) {
	ret := _m.Called(n)

	var r0 models.Addressable
	if rf, ok := ret.Get(0).(func(string) models.Addressable); ok {
		r0 = rf(n)
	} else {
		r0 = ret.Get(0).(models.Addressable)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressables provides a mock function with given fields:
func (_m *AddressLoader) GetAddressables() ([]models.Addressable, error) {
	ret := _m.Called()

	var r0 []models.Addressable
	if rf, ok := ret.Get(0).(func() []models.Addressable); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Addressable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressablesByAddress provides a mock function with given fields: address
func (_m *AddressLoader) GetAddressablesByAddress(address string) ([]models.Addressable, error) {
	ret := _m.Called(address)

	var r0 []models.Addressable
	if rf, ok := ret.Get(0).(func(string) []models.Addressable); ok {
		r0 = rf(address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Addressable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressablesByPort provides a mock function with given fields: p
func (_m *AddressLoader) GetAddressablesByPort(p int) ([]models.Addressable, error) {
	ret := _m.Called(p)

	var r0 []models.Addressable
	if rf, ok := ret.Get(0).(func(int) []models.Addressable); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Addressable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressablesByPublisher provides a mock function with given fields: p
func (_m *AddressLoader) GetAddressablesByPublisher(p string) ([]models.Addressable, error) {
	ret := _m.Called(p)

	var r0 []models.Addressable
	if rf, ok := ret.Get(0).(func(string) []models.Addressable); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Addressable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAddressablesByTopic provides a mock function with given fields: t
func (_m *AddressLoader) GetAddressablesByTopic(t string) ([]models.Addressable, error) {
	ret := _m.Called(t)

	var r0 []models.Addressable
	if rf, ok := ret.Get(0).(func(string) []models.Addressable); ok {
		r0 = rf(t)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Addressable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(t)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
