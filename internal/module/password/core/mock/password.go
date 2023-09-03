// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/module/password/core/password.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	core "monolith/internal/module/password/core"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPasswordUsecase is a mock of PasswordUsecase interface.
type MockPasswordUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordUsecaseMockRecorder
}

// MockPasswordUsecaseMockRecorder is the mock recorder for MockPasswordUsecase.
type MockPasswordUsecaseMockRecorder struct {
	mock *MockPasswordUsecase
}

// NewMockPasswordUsecase creates a new mock instance.
func NewMockPasswordUsecase(ctrl *gomock.Controller) *MockPasswordUsecase {
	mock := &MockPasswordUsecase{ctrl: ctrl}
	mock.recorder = &MockPasswordUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasswordUsecase) EXPECT() *MockPasswordUsecaseMockRecorder {
	return m.recorder
}

// ChangePassword mocks base method.
func (m *MockPasswordUsecase) ChangePassword(arg0 context.Context, arg1 core.Password) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangePassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangePassword indicates an expected call of ChangePassword.
func (mr *MockPasswordUsecaseMockRecorder) ChangePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangePassword", reflect.TypeOf((*MockPasswordUsecase)(nil).ChangePassword), arg0, arg1)
}

// MockPasswordRepository is a mock of PasswordRepository interface.
type MockPasswordRepository struct {
	ctrl     *gomock.Controller
	recorder *MockPasswordRepositoryMockRecorder
}

// MockPasswordRepositoryMockRecorder is the mock recorder for MockPasswordRepository.
type MockPasswordRepositoryMockRecorder struct {
	mock *MockPasswordRepository
}

// NewMockPasswordRepository creates a new mock instance.
func NewMockPasswordRepository(ctrl *gomock.Controller) *MockPasswordRepository {
	mock := &MockPasswordRepository{ctrl: ctrl}
	mock.recorder = &MockPasswordRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPasswordRepository) EXPECT() *MockPasswordRepositoryMockRecorder {
	return m.recorder
}

// GeneratePassword mocks base method.
func (m *MockPasswordRepository) GeneratePassword(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GeneratePassword", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GeneratePassword indicates an expected call of GeneratePassword.
func (mr *MockPasswordRepositoryMockRecorder) GeneratePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GeneratePassword", reflect.TypeOf((*MockPasswordRepository)(nil).GeneratePassword), arg0, arg1)
}