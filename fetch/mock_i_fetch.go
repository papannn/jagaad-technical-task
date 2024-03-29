// Code generated by mockery v2.30.1. DO NOT EDIT.

package fetch

import (
	config "jagaat-technical-task/config"
	dto "jagaat-technical-task/dto"

	mock "github.com/stretchr/testify/mock"
)

// MockIFetch is an autogenerated mock type for the IFetch type
type MockIFetch struct {
	mock.Mock
}

// FetchUserDataFromURLArr provides a mock function with given fields: cfg
func (_m *MockIFetch) FetchUserDataFromURLArr(cfg config.Config) []dto.User {
	ret := _m.Called(cfg)

	var r0 []dto.User
	if rf, ok := ret.Get(0).(func(config.Config) []dto.User); ok {
		r0 = rf(cfg)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.User)
		}
	}

	return r0
}

// NewMockIFetch creates a new instance of MockIFetch. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockIFetch(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockIFetch {
	mock := &MockIFetch{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
