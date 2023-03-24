// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/api/validate/dynamic (interfaces: Dynamic)

// Package mock_dynamic is a generated GoMock package.
package mock_dynamic

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	api "github.com/Azure/ARO-RP/pkg/api"
	dynamic "github.com/Azure/ARO-RP/pkg/api/validate/dynamic"
)

// MockDynamic is a mock of Dynamic interface.
type MockDynamic struct {
	ctrl     *gomock.Controller
	recorder *MockDynamicMockRecorder
}

// MockDynamicMockRecorder is the mock recorder for MockDynamic.
type MockDynamicMockRecorder struct {
	mock *MockDynamic
}

// NewMockDynamic creates a new mock instance.
func NewMockDynamic(ctrl *gomock.Controller) *MockDynamic {
	mock := &MockDynamic{ctrl: ctrl}
	mock.recorder = &MockDynamicMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDynamic) EXPECT() *MockDynamicMockRecorder {
	return m.recorder
}

// ValidateDiskEncryptionSets mocks base method.
func (m *MockDynamic) ValidateDiskEncryptionSets(arg0 context.Context, arg1 *api.OpenShiftCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateDiskEncryptionSets", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateDiskEncryptionSets indicates an expected call of ValidateDiskEncryptionSets.
func (mr *MockDynamicMockRecorder) ValidateDiskEncryptionSets(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateDiskEncryptionSets", reflect.TypeOf((*MockDynamic)(nil).ValidateDiskEncryptionSets), arg0, arg1)
}

// ValidateEncryptionAtHost mocks base method.
func (m *MockDynamic) ValidateEncryptionAtHost(arg0 context.Context, arg1 *api.OpenShiftCluster) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateEncryptionAtHost", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateEncryptionAtHost indicates an expected call of ValidateEncryptionAtHost.
func (mr *MockDynamicMockRecorder) ValidateEncryptionAtHost(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateEncryptionAtHost", reflect.TypeOf((*MockDynamic)(nil).ValidateEncryptionAtHost), arg0, arg1)
}

// ValidateProviders mocks base method.
func (m *MockDynamic) ValidateProviders(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateProviders", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateProviders indicates an expected call of ValidateProviders.
func (mr *MockDynamicMockRecorder) ValidateProviders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateProviders", reflect.TypeOf((*MockDynamic)(nil).ValidateProviders), arg0)
}

// ValidateServicePrincipal mocks base method.
func (m *MockDynamic) ValidateServicePrincipal(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateServicePrincipal", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateServicePrincipal indicates an expected call of ValidateServicePrincipal.
func (mr *MockDynamicMockRecorder) ValidateServicePrincipal(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateServicePrincipal", reflect.TypeOf((*MockDynamic)(nil).ValidateServicePrincipal), arg0, arg1, arg2, arg3)
}

// ValidateSubnets mocks base method.
func (m *MockDynamic) ValidateSubnets(arg0 context.Context, arg1 *api.OpenShiftCluster, arg2 []dynamic.Subnet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateSubnets", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateSubnets indicates an expected call of ValidateSubnets.
func (mr *MockDynamicMockRecorder) ValidateSubnets(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateSubnets", reflect.TypeOf((*MockDynamic)(nil).ValidateSubnets), arg0, arg1, arg2)
}

// ValidateVnet mocks base method.
func (m *MockDynamic) ValidateVnet(arg0 context.Context, arg1 string, arg2 []dynamic.Subnet, arg3 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ValidateVnet", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// ValidateVnet indicates an expected call of ValidateVnet.
func (mr *MockDynamicMockRecorder) ValidateVnet(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateVnet", reflect.TypeOf((*MockDynamic)(nil).ValidateVnet), varargs...)
}
