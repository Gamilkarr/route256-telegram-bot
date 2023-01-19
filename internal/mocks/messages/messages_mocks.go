// Code generated by MockGen. DO NOT EDIT.
// Source: internal/model/messages/incoming_msg.go

// Package mock_messages is a generated GoMock package.
package mock_messages

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMessageSender is a mock of MessageSender interface.
type MockMessageSender struct {
	ctrl     *gomock.Controller
	recorder *MockMessageSenderMockRecorder
}

// MockMessageSenderMockRecorder is the mock recorder for MockMessageSender.
type MockMessageSenderMockRecorder struct {
	mock *MockMessageSender
}

// NewMockMessageSender creates a new mock instance.
func NewMockMessageSender(ctrl *gomock.Controller) *MockMessageSender {
	mock := &MockMessageSender{ctrl: ctrl}
	mock.recorder = &MockMessageSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageSender) EXPECT() *MockMessageSenderMockRecorder {
	return m.recorder
}

// SendMessenge mocks base method.
func (m *MockMessageSender) SendMessenge(text string, userID int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessenge", text, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessenge indicates an expected call of SendMessenge.
func (mr *MockMessageSenderMockRecorder) SendMessenge(text, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessenge", reflect.TypeOf((*MockMessageSender)(nil).SendMessenge), text, userID)
}
