// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DatabaseI is an autogenerated mock type for the DatabaseI type
type DatabaseI struct {
	mock.Mock
}

// Create provides a mock function with given fields: value
func (_m *DatabaseI) Create(value interface{}) error {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: value
func (_m *DatabaseI) Delete(value interface{}) error {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchAll provides a mock function with given fields: value
func (_m *DatabaseI) FetchAll(value interface{}) error {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for FetchAll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchByID provides a mock function with given fields: value
func (_m *DatabaseI) FetchByID(value interface{}) error {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for FetchByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: value
func (_m *DatabaseI) Update(value interface{}) error {
	ret := _m.Called(value)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDatabaseI creates a new instance of DatabaseI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDatabaseI(t interface {
	mock.TestingT
	Cleanup(func())
}) *DatabaseI {
	mock := &DatabaseI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
