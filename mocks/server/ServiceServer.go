// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// ServiceServer is an autogenerated mock type for the ServiceServer type
type ServiceServer struct {
	mock.Mock
}

// Close provides a mock function with given fields: _a0
func (_m *ServiceServer) Close(_a0 context.Context) {
	_m.Called(_a0)
}

// RegisterWithHandler provides a mock function with given fields: _a0, _a1, _a2
func (_m *ServiceServer) RegisterWithHandler(_a0 context.Context, _a1 *runtime.ServeMux, _a2 *grpc.ClientConn) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterWithServer provides a mock function with given fields: _a0
func (_m *ServiceServer) RegisterWithServer(_a0 *grpc.Server) {
	_m.Called(_a0)
}