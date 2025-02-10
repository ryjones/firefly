// Code generated by mockery v2.46.0. DO NOT EDIT.

package batchmocks

import (
	context "context"

	batch "github.com/hyperledger/firefly/internal/batch"

	fftypes "github.com/hyperledger/firefly-common/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// CancelBatch provides a mock function with given fields: ctx, batchID
func (_m *Manager) CancelBatch(ctx context.Context, batchID string) error {
	ret := _m.Called(ctx, batchID)

	if len(ret) == 0 {
		panic("no return value specified for CancelBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, batchID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *Manager) Close() {
	_m.Called()
}

// LoadContexts provides a mock function with given fields: ctx, payload
func (_m *Manager) LoadContexts(ctx context.Context, payload *batch.DispatchPayload) error {
	ret := _m.Called(ctx, payload)

	if len(ret) == 0 {
		panic("no return value specified for LoadContexts")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *batch.DispatchPayload) error); ok {
		r0 = rf(ctx, payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMessages provides a mock function with given fields:
func (_m *Manager) NewMessages() chan<- int64 {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for NewMessages")
	}

	var r0 chan<- int64
	if rf, ok := ret.Get(0).(func() chan<- int64); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan<- int64)
		}
	}

	return r0
}

// RegisterDispatcher provides a mock function with given fields: name, pinned, msgTypes, handler, batchOptions
func (_m *Manager) RegisterDispatcher(name string, pinned bool, msgTypes []fftypes.FFEnum, handler batch.DispatchHandler, batchOptions batch.DispatcherOptions) {
	_m.Called(name, pinned, msgTypes, handler, batchOptions)
}

// Start provides a mock function with given fields:
func (_m *Manager) Start() error {
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

// Status provides a mock function with given fields:
func (_m *Manager) Status() *batch.ManagerStatus {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Status")
	}

	var r0 *batch.ManagerStatus
	if rf, ok := ret.Get(0).(func() *batch.ManagerStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*batch.ManagerStatus)
		}
	}

	return r0
}

// WaitStop provides a mock function with given fields:
func (_m *Manager) WaitStop() {
	_m.Called()
}

// NewManager creates a new instance of Manager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *Manager {
	mock := &Manager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
