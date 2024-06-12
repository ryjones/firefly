<<<<<<< HEAD
// Code generated by mockery v2.40.2. DO NOT EDIT.
=======
// Code generated by mockery v2.42.1. DO NOT EDIT.
>>>>>>> origin/main

package identitymocks

import (
	context "context"

	config "github.com/hyperledger/firefly-common/pkg/config"

	identity "github.com/hyperledger/firefly/pkg/identity"

	mock "github.com/stretchr/testify/mock"
)

// Plugin is an autogenerated mock type for the Plugin type
type Plugin struct {
	mock.Mock
}

// Capabilities provides a mock function with given fields:
func (_m *Plugin) Capabilities() *identity.Capabilities {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Capabilities")
	}

	var r0 *identity.Capabilities
	if rf, ok := ret.Get(0).(func() *identity.Capabilities); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*identity.Capabilities)
		}
	}

	return r0
}

// Init provides a mock function with given fields: ctx, _a1
func (_m *Plugin) Init(ctx context.Context, _a1 config.Section) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Init")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, config.Section) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InitConfig provides a mock function with given fields: _a0
func (_m *Plugin) InitConfig(_a0 config.Section) {
	_m.Called(_a0)
}

// Name provides a mock function with given fields:
func (_m *Plugin) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SetHandler provides a mock function with given fields: namespace, handler
func (_m *Plugin) SetHandler(namespace string, handler identity.Callbacks) {
	_m.Called(namespace, handler)
}

// Start provides a mock function with given fields:
func (_m *Plugin) Start() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPlugin creates a new instance of Plugin. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPlugin(t interface {
	mock.TestingT
	Cleanup(func())
}) *Plugin {
	mock := &Plugin{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
