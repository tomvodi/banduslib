// Code generated by MockGen. DO NOT EDIT.
// Source: embellishment_expander.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	music_model "banduslib/internal/common/music_model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSymbolExpander is a mock of SymbolExpander interface.
type MockSymbolExpander struct {
	ctrl     *gomock.Controller
	recorder *MockSymbolExpanderMockRecorder
}

// MockSymbolExpanderMockRecorder is the mock recorder for MockSymbolExpander.
type MockSymbolExpanderMockRecorder struct {
	mock *MockSymbolExpander
}

// NewMockSymbolExpander creates a new mock instance.
func NewMockSymbolExpander(ctrl *gomock.Controller) *MockSymbolExpander {
	mock := &MockSymbolExpander{ctrl: ctrl}
	mock.recorder = &MockSymbolExpanderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSymbolExpander) EXPECT() *MockSymbolExpanderMockRecorder {
	return m.recorder
}

// ExpandSymbol mocks base method.
func (m *MockSymbolExpander) ExpandSymbol(symbol *music_model.Symbol) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExpandSymbol", symbol)
}

// ExpandSymbol indicates an expected call of ExpandSymbol.
func (mr *MockSymbolExpanderMockRecorder) ExpandSymbol(symbol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpandSymbol", reflect.TypeOf((*MockSymbolExpander)(nil).ExpandSymbol), symbol)
}

// MockEmbellishmentExpander is a mock of EmbellishmentExpander interface.
type MockEmbellishmentExpander struct {
	ctrl     *gomock.Controller
	recorder *MockEmbellishmentExpanderMockRecorder
}

// MockEmbellishmentExpanderMockRecorder is the mock recorder for MockEmbellishmentExpander.
type MockEmbellishmentExpanderMockRecorder struct {
	mock *MockEmbellishmentExpander
}

// NewMockEmbellishmentExpander creates a new mock instance.
func NewMockEmbellishmentExpander(ctrl *gomock.Controller) *MockEmbellishmentExpander {
	mock := &MockEmbellishmentExpander{ctrl: ctrl}
	mock.recorder = &MockEmbellishmentExpanderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEmbellishmentExpander) EXPECT() *MockEmbellishmentExpanderMockRecorder {
	return m.recorder
}

// ExpandModel mocks base method.
func (m *MockEmbellishmentExpander) ExpandModel(model music_model.MusicModel) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExpandModel", model)
}

// ExpandModel indicates an expected call of ExpandModel.
func (mr *MockEmbellishmentExpanderMockRecorder) ExpandModel(model interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpandModel", reflect.TypeOf((*MockEmbellishmentExpander)(nil).ExpandModel), model)
}

// ExpandSymbol mocks base method.
func (m *MockEmbellishmentExpander) ExpandSymbol(symbol *music_model.Symbol) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExpandSymbol", symbol)
}

// ExpandSymbol indicates an expected call of ExpandSymbol.
func (mr *MockEmbellishmentExpanderMockRecorder) ExpandSymbol(symbol interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpandSymbol", reflect.TypeOf((*MockEmbellishmentExpander)(nil).ExpandSymbol), symbol)
}

// ExpandTune mocks base method.
func (m *MockEmbellishmentExpander) ExpandTune(tune *music_model.Tune) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExpandTune", tune)
}

// ExpandTune indicates an expected call of ExpandTune.
func (mr *MockEmbellishmentExpanderMockRecorder) ExpandTune(tune interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExpandTune", reflect.TypeOf((*MockEmbellishmentExpander)(nil).ExpandTune), tune)
}
