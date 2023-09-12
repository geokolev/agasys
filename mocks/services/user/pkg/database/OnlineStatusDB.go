// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	database "github.com/jonsch318/royalafg/services/user/pkg/database"
	mock "github.com/stretchr/testify/mock"
)

// OnlineStatusDB is an autogenerated mock type for the OnlineStatusDB type
type OnlineStatusDB struct {
	mock.Mock
}

// GetOnlineStatus provides a mock function with given fields: id
func (_m *OnlineStatusDB) GetOnlineStatus(id string) (*database.OnlineStatus, error) {
	ret := _m.Called(id)

	var r0 *database.OnlineStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*database.OnlineStatus, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *database.OnlineStatus); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*database.OnlineStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetOnlineStatus provides a mock function with given fields: id, status
func (_m *OnlineStatusDB) SetOnlineStatus(id string, status *database.OnlineStatus) error {
	ret := _m.Called(id, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *database.OnlineStatus) error); ok {
		r0 = rf(id, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOnlineStatusDB creates a new instance of OnlineStatusDB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOnlineStatusDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *OnlineStatusDB {
	mock := &OnlineStatusDB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
