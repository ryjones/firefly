// Code generated by mockery v1.0.0. DO NOT EDIT.

package broadcastmocks

import (
	context "context"

	fftypes "github.com/hyperledger/firefly-common/pkg/fftypes"
	core "github.com/hyperledger/firefly/pkg/core"

	mock "github.com/stretchr/testify/mock"

	sysmessaging "github.com/hyperledger/firefly/internal/sysmessaging"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// BroadcastDatatype provides a mock function with given fields: ctx, datatype, waitConfirm
func (_m *Manager) BroadcastDatatype(ctx context.Context, datatype *core.Datatype, waitConfirm bool) (*core.Message, error) {
	ret := _m.Called(ctx, datatype, waitConfirm)

	var r0 *core.Message
	if rf, ok := ret.Get(0).(func(context.Context, *core.Datatype, bool) *core.Message); ok {
		r0 = rf(ctx, datatype, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.Datatype, bool) error); ok {
		r1 = rf(ctx, datatype, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BroadcastDefinition provides a mock function with given fields: ctx, def, signingIdentity, tag, waitConfirm
func (_m *Manager) BroadcastDefinition(ctx context.Context, def core.Definition, signingIdentity *core.SignerRef, tag string, waitConfirm bool) (*core.Message, error) {
	ret := _m.Called(ctx, def, signingIdentity, tag, waitConfirm)

	var r0 *core.Message
	if rf, ok := ret.Get(0).(func(context.Context, core.Definition, *core.SignerRef, string, bool) *core.Message); ok {
		r0 = rf(ctx, def, signingIdentity, tag, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, core.Definition, *core.SignerRef, string, bool) error); ok {
		r1 = rf(ctx, def, signingIdentity, tag, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BroadcastDefinitionAsNode provides a mock function with given fields: ctx, def, tag, waitConfirm
func (_m *Manager) BroadcastDefinitionAsNode(ctx context.Context, def core.Definition, tag string, waitConfirm bool) (*core.Message, error) {
	ret := _m.Called(ctx, def, tag, waitConfirm)

	var r0 *core.Message
	if rf, ok := ret.Get(0).(func(context.Context, core.Definition, string, bool) *core.Message); ok {
		r0 = rf(ctx, def, tag, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, core.Definition, string, bool) error); ok {
		r1 = rf(ctx, def, tag, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BroadcastIdentityClaim provides a mock function with given fields: ctx, def, signingIdentity, tag, waitConfirm
func (_m *Manager) BroadcastIdentityClaim(ctx context.Context, def *core.IdentityClaim, signingIdentity *core.SignerRef, tag string, waitConfirm bool) (*core.Message, error) {
	ret := _m.Called(ctx, def, signingIdentity, tag, waitConfirm)

	var r0 *core.Message
	if rf, ok := ret.Get(0).(func(context.Context, *core.IdentityClaim, *core.SignerRef, string, bool) *core.Message); ok {
		r0 = rf(ctx, def, signingIdentity, tag, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.IdentityClaim, *core.SignerRef, string, bool) error); ok {
		r1 = rf(ctx, def, signingIdentity, tag, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BroadcastMessage provides a mock function with given fields: ctx, in, waitConfirm
func (_m *Manager) BroadcastMessage(ctx context.Context, in *core.MessageInOut, waitConfirm bool) (*core.Message, error) {
	ret := _m.Called(ctx, in, waitConfirm)

	var r0 *core.Message
	if rf, ok := ret.Get(0).(func(context.Context, *core.MessageInOut, bool) *core.Message); ok {
		r0 = rf(ctx, in, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.MessageInOut, bool) error); ok {
		r1 = rf(ctx, in, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BroadcastTokenPool provides a mock function with given fields: ctx, pool, waitConfirm
func (_m *Manager) BroadcastTokenPool(ctx context.Context, pool *core.TokenPoolAnnouncement, waitConfirm bool) (*core.Message, error) {
	ret := _m.Called(ctx, pool, waitConfirm)

	var r0 *core.Message
	if rf, ok := ret.Get(0).(func(context.Context, *core.TokenPoolAnnouncement, bool) *core.Message); ok {
		r0 = rf(ctx, pool, waitConfirm)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Message)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.TokenPoolAnnouncement, bool) error); ok {
		r1 = rf(ctx, pool, waitConfirm)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *Manager) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewBroadcast provides a mock function with given fields: in
func (_m *Manager) NewBroadcast(in *core.MessageInOut) sysmessaging.MessageSender {
	ret := _m.Called(in)

	var r0 sysmessaging.MessageSender
	if rf, ok := ret.Get(0).(func(*core.MessageInOut) sysmessaging.MessageSender); ok {
		r0 = rf(in)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(sysmessaging.MessageSender)
		}
	}

	return r0
}

// PrepareOperation provides a mock function with given fields: ctx, op
func (_m *Manager) PrepareOperation(ctx context.Context, op *core.Operation) (*core.PreparedOperation, error) {
	ret := _m.Called(ctx, op)

	var r0 *core.PreparedOperation
	if rf, ok := ret.Get(0).(func(context.Context, *core.Operation) *core.PreparedOperation); ok {
		r0 = rf(ctx, op)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.PreparedOperation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.Operation) error); ok {
		r1 = rf(ctx, op)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunOperation provides a mock function with given fields: ctx, op
func (_m *Manager) RunOperation(ctx context.Context, op *core.PreparedOperation) (fftypes.JSONObject, bool, error) {
	ret := _m.Called(ctx, op)

	var r0 fftypes.JSONObject
	if rf, ok := ret.Get(0).(func(context.Context, *core.PreparedOperation) fftypes.JSONObject); ok {
		r0 = rf(ctx, op)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(fftypes.JSONObject)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(context.Context, *core.PreparedOperation) bool); ok {
		r1 = rf(ctx, op)
	} else {
		r1 = ret.Get(1).(bool)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *core.PreparedOperation) error); ok {
		r2 = rf(ctx, op)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Start provides a mock function with given fields:
func (_m *Manager) Start() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitStop provides a mock function with given fields:
func (_m *Manager) WaitStop() {
	_m.Called()
}
