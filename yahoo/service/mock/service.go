// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockYahooNews is a mock of YahooNews interface.
type MockYahooNews struct {
	ctrl     *gomock.Controller
	recorder *MockYahooNewsMockRecorder
}

// MockYahooNewsMockRecorder is the mock recorder for MockYahooNews.
type MockYahooNewsMockRecorder struct {
	mock *MockYahooNews
}

// NewMockYahooNews creates a new mock instance.
func NewMockYahooNews(ctrl *gomock.Controller) *MockYahooNews {
	mock := &MockYahooNews{ctrl: ctrl}
	mock.recorder = &MockYahooNewsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockYahooNews) EXPECT() *MockYahooNewsMockRecorder {
	return m.recorder
}

// FetchAndStore mocks base method.
func (m *MockYahooNews) FetchAndStore(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAndStore", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// FetchAndStore indicates an expected call of FetchAndStore.
func (mr *MockYahooNewsMockRecorder) FetchAndStore(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAndStore", reflect.TypeOf((*MockYahooNews)(nil).FetchAndStore), ctx)
}
