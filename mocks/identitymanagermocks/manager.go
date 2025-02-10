// Code generated by mockery v2.46.0. DO NOT EDIT.

package identitymanagermocks

import (
	context "context"

	blockchain "github.com/hyperledger/firefly/pkg/blockchain"

	core "github.com/hyperledger/firefly/pkg/core"

	fftypes "github.com/hyperledger/firefly-common/pkg/fftypes"

	mock "github.com/stretchr/testify/mock"
)

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// CachedIdentityLookupByID provides a mock function with given fields: ctx, id
func (_m *Manager) CachedIdentityLookupByID(ctx context.Context, id *fftypes.UUID) (*core.Identity, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for CachedIdentityLookupByID")
	}

	var r0 *core.Identity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) (*core.Identity, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *core.Identity); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CachedIdentityLookupMustExist provides a mock function with given fields: ctx, did
func (_m *Manager) CachedIdentityLookupMustExist(ctx context.Context, did string) (*core.Identity, bool, error) {
	ret := _m.Called(ctx, did)

	if len(ret) == 0 {
		panic("no return value specified for CachedIdentityLookupMustExist")
	}

	var r0 *core.Identity
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*core.Identity, bool, error)); ok {
		return rf(ctx, did)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *core.Identity); ok {
		r0 = rf(ctx, did)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) bool); ok {
		r1 = rf(ctx, did)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, did)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// CachedIdentityLookupNilOK provides a mock function with given fields: ctx, did
func (_m *Manager) CachedIdentityLookupNilOK(ctx context.Context, did string) (*core.Identity, bool, error) {
	ret := _m.Called(ctx, did)

	if len(ret) == 0 {
		panic("no return value specified for CachedIdentityLookupNilOK")
	}

	var r0 *core.Identity
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*core.Identity, bool, error)); ok {
		return rf(ctx, did)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *core.Identity); ok {
		r0 = rf(ctx, did)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) bool); ok {
		r1 = rf(ctx, did)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, did)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindIdentityForVerifier provides a mock function with given fields: ctx, iTypes, verifier
func (_m *Manager) FindIdentityForVerifier(ctx context.Context, iTypes []fftypes.FFEnum, verifier *core.VerifierRef) (*core.Identity, error) {
	ret := _m.Called(ctx, iTypes, verifier)

	if len(ret) == 0 {
		panic("no return value specified for FindIdentityForVerifier")
	}

	var r0 *core.Identity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []fftypes.FFEnum, *core.VerifierRef) (*core.Identity, error)); ok {
		return rf(ctx, iTypes, verifier)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []fftypes.FFEnum, *core.VerifierRef) *core.Identity); ok {
		r0 = rf(ctx, iTypes, verifier)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []fftypes.FFEnum, *core.VerifierRef) error); ok {
		r1 = rf(ctx, iTypes, verifier)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLocalNode provides a mock function with given fields: ctx
func (_m *Manager) GetLocalNode(ctx context.Context) (*core.Identity, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetLocalNode")
	}

	var r0 *core.Identity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*core.Identity, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *core.Identity); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootOrg provides a mock function with given fields: ctx
func (_m *Manager) GetRootOrg(ctx context.Context) (*core.Identity, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetRootOrg")
	}

	var r0 *core.Identity
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*core.Identity, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *core.Identity); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRootOrgDID provides a mock function with given fields: ctx
func (_m *Manager) GetRootOrgDID(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetRootOrgDID")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResolveIdentitySigner provides a mock function with given fields: ctx, _a1
func (_m *Manager) ResolveIdentitySigner(ctx context.Context, _a1 *core.Identity) (*core.SignerRef, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ResolveIdentitySigner")
	}

	var r0 *core.SignerRef
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.Identity) (*core.SignerRef, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.Identity) *core.SignerRef); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.SignerRef)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.Identity) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResolveInputSigningIdentity provides a mock function with given fields: ctx, signerRef
func (_m *Manager) ResolveInputSigningIdentity(ctx context.Context, signerRef *core.SignerRef) error {
	ret := _m.Called(ctx, signerRef)

	if len(ret) == 0 {
		panic("no return value specified for ResolveInputSigningIdentity")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.SignerRef) error); ok {
		r0 = rf(ctx, signerRef)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResolveInputSigningKey provides a mock function with given fields: ctx, inputKey, keyNormalizationMode
func (_m *Manager) ResolveInputSigningKey(ctx context.Context, inputKey string, keyNormalizationMode int) (string, error) {
	ret := _m.Called(ctx, inputKey, keyNormalizationMode)

	if len(ret) == 0 {
		panic("no return value specified for ResolveInputSigningKey")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) (string, error)); ok {
		return rf(ctx, inputKey, keyNormalizationMode)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int) string); ok {
		r0 = rf(ctx, inputKey, keyNormalizationMode)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, inputKey, keyNormalizationMode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResolveInputVerifierRef provides a mock function with given fields: ctx, inputKey, intent
func (_m *Manager) ResolveInputVerifierRef(ctx context.Context, inputKey *core.VerifierRef, intent blockchain.ResolveKeyIntent) (*core.VerifierRef, error) {
	ret := _m.Called(ctx, inputKey, intent)

	if len(ret) == 0 {
		panic("no return value specified for ResolveInputVerifierRef")
	}

	var r0 *core.VerifierRef
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.VerifierRef, blockchain.ResolveKeyIntent) (*core.VerifierRef, error)); ok {
		return rf(ctx, inputKey, intent)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.VerifierRef, blockchain.ResolveKeyIntent) *core.VerifierRef); ok {
		r0 = rf(ctx, inputKey, intent)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.VerifierRef)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.VerifierRef, blockchain.ResolveKeyIntent) error); ok {
		r1 = rf(ctx, inputKey, intent)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResolveMultipartyRootVerifier provides a mock function with given fields: ctx
func (_m *Manager) ResolveMultipartyRootVerifier(ctx context.Context) (*core.VerifierRef, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ResolveMultipartyRootVerifier")
	}

	var r0 *core.VerifierRef
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*core.VerifierRef, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *core.VerifierRef); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.VerifierRef)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResolveQuerySigningKey provides a mock function with given fields: ctx, inputKey, keyNormalizationMode
func (_m *Manager) ResolveQuerySigningKey(ctx context.Context, inputKey string, keyNormalizationMode int) (string, error) {
	ret := _m.Called(ctx, inputKey, keyNormalizationMode)

	if len(ret) == 0 {
		panic("no return value specified for ResolveQuerySigningKey")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) (string, error)); ok {
		return rf(ctx, inputKey, keyNormalizationMode)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, int) string); ok {
		r0 = rf(ctx, inputKey, keyNormalizationMode)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, int) error); ok {
		r1 = rf(ctx, inputKey, keyNormalizationMode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ValidateNodeOwner provides a mock function with given fields: ctx, node, _a2
func (_m *Manager) ValidateNodeOwner(ctx context.Context, node *core.Identity, _a2 *core.Identity) (bool, error) {
	ret := _m.Called(ctx, node, _a2)

	if len(ret) == 0 {
		panic("no return value specified for ValidateNodeOwner")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.Identity, *core.Identity) (bool, error)); ok {
		return rf(ctx, node, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.Identity, *core.Identity) bool); ok {
		r0 = rf(ctx, node, _a2)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.Identity, *core.Identity) error); ok {
		r1 = rf(ctx, node, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// VerifyIdentityChain provides a mock function with given fields: ctx, _a1
func (_m *Manager) VerifyIdentityChain(ctx context.Context, _a1 *core.Identity) (*core.Identity, bool, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for VerifyIdentityChain")
	}

	var r0 *core.Identity
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.Identity) (*core.Identity, bool, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *core.Identity) *core.Identity); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Identity)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *core.Identity) bool); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *core.Identity) error); ok {
		r2 = rf(ctx, _a1)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
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
