// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go
//
// Generated by this command:
//
//	mockgen -source=storage.go -destination=mock_storage.go -package=storage Storage
//

// Package storage is a generated GoMock package.
package storage

import (
	context "context"
	reflect "reflect"

	url "github.com/bentoml/s5cmd/v2/storage/url"
	gomock "go.uber.org/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Copy mocks base method.
func (m *MockStorage) Copy(ctx context.Context, src, dst *url.URL, metadata Metadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Copy", ctx, src, dst, metadata)
	ret0, _ := ret[0].(error)
	return ret0
}

// Copy indicates an expected call of Copy.
func (mr *MockStorageMockRecorder) Copy(ctx, src, dst, metadata any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockStorage)(nil).Copy), ctx, src, dst, metadata)
}

// Delete mocks base method.
func (m *MockStorage) Delete(ctx context.Context, src *url.URL) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, src)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockStorageMockRecorder) Delete(ctx, src any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockStorage)(nil).Delete), ctx, src)
}

// List mocks base method.
func (m *MockStorage) List(ctx context.Context, src *url.URL, followSymlinks bool) <-chan *Object {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, src, followSymlinks)
	ret0, _ := ret[0].(<-chan *Object)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockStorageMockRecorder) List(ctx, src, followSymlinks any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStorage)(nil).List), ctx, src, followSymlinks)
}

// MultiDelete mocks base method.
func (m *MockStorage) MultiDelete(ctx context.Context, urls <-chan *url.URL) <-chan *Object {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MultiDelete", ctx, urls)
	ret0, _ := ret[0].(<-chan *Object)
	return ret0
}

// MultiDelete indicates an expected call of MultiDelete.
func (mr *MockStorageMockRecorder) MultiDelete(ctx, urls any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MultiDelete", reflect.TypeOf((*MockStorage)(nil).MultiDelete), ctx, urls)
}

// Stat mocks base method.
func (m *MockStorage) Stat(ctx context.Context, src *url.URL) (*Object, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat", ctx, src)
	ret0, _ := ret[0].(*Object)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat.
func (mr *MockStorageMockRecorder) Stat(ctx, src any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockStorage)(nil).Stat), ctx, src)
}
