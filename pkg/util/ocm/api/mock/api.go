// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/ocm/api (interfaces: API)

// Package mock_api is a generated GoMock package.
package mock_api

import (
	"context"
	"reflect"

	gomock "github.com/golang/mock/gomock"

	api "github.com/Azure/ARO-RP/pkg/util/ocm/api"
)

// MockAPI is a mock of API interface.
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI.
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance.
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// CancelClusterUpgradePolicy mocks base method.
func (m *MockAPI) CancelClusterUpgradePolicy(arg0 context.Context, arg1, arg2 string) (*api.CancelUpgradeResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelClusterUpgradePolicy", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.CancelUpgradeResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelClusterUpgradePolicy indicates an expected call of CancelClusterUpgradePolicy.
func (mr *MockAPIMockRecorder) CancelClusterUpgradePolicy(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelClusterUpgradePolicy", reflect.TypeOf((*MockAPI)(nil).CancelClusterUpgradePolicy), arg0, arg1, arg2)
}

// GetClusterList mocks base method.
func (m *MockAPI) GetClusterList(arg0 context.Context, arg1 map[string]string) (*api.ClusterList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterList", arg0, arg1)
	ret0, _ := ret[0].(*api.ClusterList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterList indicates an expected call of GetClusterList.
func (mr *MockAPIMockRecorder) GetClusterList(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterList", reflect.TypeOf((*MockAPI)(nil).GetClusterList), arg0, arg1)
}

// GetClusterUpgradePolicies mocks base method.
func (m *MockAPI) GetClusterUpgradePolicies(arg0 context.Context, arg1 string) (*api.UpgradePolicyList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterUpgradePolicies", arg0, arg1)
	ret0, _ := ret[0].(*api.UpgradePolicyList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterUpgradePolicies indicates an expected call of GetClusterUpgradePolicies.
func (mr *MockAPIMockRecorder) GetClusterUpgradePolicies(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterUpgradePolicies", reflect.TypeOf((*MockAPI)(nil).GetClusterUpgradePolicies), arg0, arg1)
}

// GetClusterUpgradePolicyState mocks base method.
func (m *MockAPI) GetClusterUpgradePolicyState(arg0 context.Context, arg1, arg2 string) (*api.UpgradePolicyState, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterUpgradePolicyState", arg0, arg1, arg2)
	ret0, _ := ret[0].(*api.UpgradePolicyState)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterUpgradePolicyState indicates an expected call of GetClusterUpgradePolicyState.
func (mr *MockAPIMockRecorder) GetClusterUpgradePolicyState(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterUpgradePolicyState", reflect.TypeOf((*MockAPI)(nil).GetClusterUpgradePolicyState), arg0, arg1, arg2)
}
