// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/ARO-RP/pkg/util/storage (interfaces: Manager)

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	reflect "reflect"

	container "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/container"
	gomock "github.com/golang/mock/gomock"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// BlobService mocks base method.
func (m *MockManager) BlobService(arg0, arg1 string) (*container.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlobService", arg0, arg1)
	ret0, _ := ret[0].(*container.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlobService indicates an expected call of BlobService.
func (mr *MockManagerMockRecorder) BlobService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlobService", reflect.TypeOf((*MockManager)(nil).BlobService), arg0, arg1)
}
