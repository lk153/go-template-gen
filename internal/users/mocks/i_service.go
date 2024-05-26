// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	users "github.com/lk153/go-template-gen/internal/users"
	mock "github.com/stretchr/testify/mock"
)

// IService is an autogenerated mock type for the IService type
type IService struct {
	mock.Mock
}

// GetUsers provides a mock function with given fields: id
func (_m *IService) GetUsers(id uint) (*users.UserResp, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetUsers")
	}

	var r0 *users.UserResp
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*users.UserResp, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) *users.UserResp); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.UserResp)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseFile provides a mock function with given fields: data
func (_m *IService) ParseFile(data string) error {
	ret := _m.Called(data)

	if len(ret) == 0 {
		panic("no return value specified for ParseFile")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIService creates a new instance of IService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IService {
	mock := &IService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
