// Code generated by mockery v2.0.0-alpha.2. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Godview is an autogenerated mock type for the Godview type
type Godview struct {
	mock.Mock
}

// Localize provides a mock function with given fields:
func (_m *Godview) Localize() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
