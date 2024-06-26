// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// FelineService is an autogenerated mock type for the FelineService type
type FelineService struct {
	mock.Mock
}

// RoarAll provides a mock function with given fields:
func (_m *FelineService) RoarAll() {
	_m.Called()
}

// RoarAllAsync provides a mock function with given fields: asyncLen
func (_m *FelineService) RoarAllAsync(asyncLen int) {
	_m.Called(asyncLen)
}

// NewFelineService creates a new instance of FelineService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFelineService(t interface {
	mock.TestingT
	Cleanup(func())
}) *FelineService {
	mock := &FelineService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
