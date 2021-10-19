// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	models "BookStore/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// AddBook provides a mock function with given fields: _a0
func (_m *Repository) AddBook(_a0 models.BStore) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.BStore) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteBookByID provides a mock function with given fields: _a0
func (_m *Repository) DeleteBookByID(_a0 int) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBooks provides a mock function with given fields:
func (_m *Repository) GetBooks() ([]models.BStore, error) {
	ret := _m.Called()

	var r0 []models.BStore
	if rf, ok := ret.Get(0).(func() []models.BStore); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.BStore)
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

// UpdateBookByID provides a mock function with given fields: _a0
func (_m *Repository) UpdateBookByID(_a0 models.BStore) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.BStore) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
