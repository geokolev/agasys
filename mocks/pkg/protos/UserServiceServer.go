// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	protos "github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	mock "github.com/stretchr/testify/mock"
)

// UserServiceServer is an autogenerated mock type for the UserServiceServer type
type UserServiceServer struct {
	mock.Mock
}

// GetUserById provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) GetUserById(_a0 context.Context, _a1 *protos.GetUser) (*protos.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *protos.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *protos.GetUser) (*protos.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *protos.GetUser) *protos.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *protos.GetUser) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUsername provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) GetUserByUsername(_a0 context.Context, _a1 *protos.GetUser) (*protos.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *protos.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *protos.GetUser) (*protos.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *protos.GetUser) *protos.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *protos.GetUser) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserStatus provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) GetUserStatus(_a0 context.Context, _a1 *protos.UserStatusRequest) (*protos.UserStatusResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *protos.UserStatusResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *protos.UserStatusRequest) (*protos.UserStatusResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *protos.UserStatusRequest) *protos.UserStatusResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.UserStatusResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *protos.UserStatusRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveUser provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) SaveUser(_a0 context.Context, _a1 *protos.User) (*protos.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *protos.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *protos.User) (*protos.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *protos.User) *protos.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *protos.User) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) UpdateUser(_a0 context.Context, _a1 *protos.User) (*protos.User, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *protos.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *protos.User) (*protos.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *protos.User) *protos.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*protos.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *protos.User) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// mustEmbedUnimplementedUserServiceServer provides a mock function with given fields:
func (_m *UserServiceServer) mustEmbedUnimplementedUserServiceServer() {
	_m.Called()
}

// NewUserServiceServer creates a new instance of UserServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserServiceServer {
	mock := &UserServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
