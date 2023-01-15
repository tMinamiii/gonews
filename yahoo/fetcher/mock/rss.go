// Code generated by MockGen. DO NOT EDIT.
// Source: rss.go

// Package fetcher is a generated GoMock package.
package fetcher

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/tMinamii/gonews/yahoo/model"
)

// MockRSSFetcher is a mock of RSSFetcher interface.
type MockRSSFetcher struct {
	ctrl     *gomock.Controller
	recorder *MockRSSFetcherMockRecorder
}

// MockRSSFetcherMockRecorder is the mock recorder for MockRSSFetcher.
type MockRSSFetcherMockRecorder struct {
	mock *MockRSSFetcher
}

// NewMockRSSFetcher creates a new mock instance.
func NewMockRSSFetcher(ctrl *gomock.Controller) *MockRSSFetcher {
	mock := &MockRSSFetcher{ctrl: ctrl}
	mock.recorder = &MockRSSFetcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRSSFetcher) EXPECT() *MockRSSFetcherMockRecorder {
	return m.recorder
}

// Fetch mocks base method.
func (m *MockRSSFetcher) Fetch(ctx context.Context, URL string) (*model.YahooNewsRSS, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fetch", ctx, URL)
	ret0, _ := ret[0].(*model.YahooNewsRSS)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fetch indicates an expected call of Fetch.
func (mr *MockRSSFetcherMockRecorder) Fetch(ctx, URL interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fetch", reflect.TypeOf((*MockRSSFetcher)(nil).Fetch), ctx, URL)
}