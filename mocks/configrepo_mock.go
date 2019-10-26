// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/configrepo/init.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	go_version "github.com/hashicorp/go-version"
	reflect "reflect"
	time "time"
)

// MockRepo is a mock of Repo interface
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// Init mocks base method
func (m *MockRepo) Init() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init")
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockRepoMockRecorder) Init() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockRepo)(nil).Init))
}

// GetAppsVersions mocks base method
func (m *MockRepo) GetAppsVersions() map[string][]*go_version.Version {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppsVersions")
	ret0, _ := ret[0].(map[string][]*go_version.Version)
	return ret0
}

// GetAppsVersions indicates an expected call of GetAppsVersions
func (mr *MockRepoMockRecorder) GetAppsVersions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppsVersions", reflect.TypeOf((*MockRepo)(nil).GetAppsVersions))
}

// GetFile mocks base method
func (m *MockRepo) GetFile(targetApp, targetVersion, path string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFile", targetApp, targetVersion, path)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFile indicates an expected call of GetFile
func (mr *MockRepoMockRecorder) GetFile(targetApp, targetVersion, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFile", reflect.TypeOf((*MockRepo)(nil).GetFile), targetApp, targetVersion, path)
}

// Pull mocks base method
func (m *MockRepo) Pull() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pull")
	ret0, _ := ret[0].(error)
	return ret0
}

// Pull indicates an expected call of Pull
func (mr *MockRepoMockRecorder) Pull() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pull", reflect.TypeOf((*MockRepo)(nil).Pull))
}

// StartPullingEvery mocks base method
func (m *MockRepo) StartPullingEvery(period time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartPullingEvery", period)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartPullingEvery indicates an expected call of StartPullingEvery
func (mr *MockRepoMockRecorder) StartPullingEvery(period interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartPullingEvery", reflect.TypeOf((*MockRepo)(nil).StartPullingEvery), period)
}

// StopPulling mocks base method
func (m *MockRepo) StopPulling() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StopPulling")
}

// StopPulling indicates an expected call of StopPulling
func (mr *MockRepoMockRecorder) StopPulling() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopPulling", reflect.TypeOf((*MockRepo)(nil).StopPulling))
}
