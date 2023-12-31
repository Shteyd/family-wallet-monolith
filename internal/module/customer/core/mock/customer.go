// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/module/customer/core/customer.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	core "monolith/internal/module/customer/core"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCustomerUsecase is a mock of CustomerUsecase interface.
type MockCustomerUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerUsecaseMockRecorder
}

// MockCustomerUsecaseMockRecorder is the mock recorder for MockCustomerUsecase.
type MockCustomerUsecaseMockRecorder struct {
	mock *MockCustomerUsecase
}

// NewMockCustomerUsecase creates a new mock instance.
func NewMockCustomerUsecase(ctrl *gomock.Controller) *MockCustomerUsecase {
	mock := &MockCustomerUsecase{ctrl: ctrl}
	mock.recorder = &MockCustomerUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerUsecase) EXPECT() *MockCustomerUsecaseMockRecorder {
	return m.recorder
}

// GetByCreds mocks base method.
func (m *MockCustomerUsecase) GetByCreds(arg0 context.Context, arg1 core.Customer) (core.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCreds", arg0, arg1)
	ret0, _ := ret[0].(core.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCreds indicates an expected call of GetByCreds.
func (mr *MockCustomerUsecaseMockRecorder) GetByCreds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCreds", reflect.TypeOf((*MockCustomerUsecase)(nil).GetByCreds), arg0, arg1)
}

// Save mocks base method.
func (m *MockCustomerUsecase) Save(arg0 context.Context, arg1 core.Customer) (core.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(core.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockCustomerUsecaseMockRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCustomerUsecase)(nil).Save), arg0, arg1)
}

// MockCustomerRepository is a mock of CustomerRepository interface.
type MockCustomerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerRepositoryMockRecorder
}

// MockCustomerRepositoryMockRecorder is the mock recorder for MockCustomerRepository.
type MockCustomerRepositoryMockRecorder struct {
	mock *MockCustomerRepository
}

// NewMockCustomerRepository creates a new mock instance.
func NewMockCustomerRepository(ctrl *gomock.Controller) *MockCustomerRepository {
	mock := &MockCustomerRepository{ctrl: ctrl}
	mock.recorder = &MockCustomerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerRepository) EXPECT() *MockCustomerRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCustomerRepository) Create(arg0 context.Context, arg1 core.Customer) (core.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(core.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCustomerRepositoryMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCustomerRepository)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockCustomerRepository) Delete(arg0 context.Context, arg1 core.Customer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCustomerRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCustomerRepository)(nil).Delete), arg0, arg1)
}

// GetByCreds mocks base method.
func (m *MockCustomerRepository) GetByCreds(arg0 context.Context, arg1 core.Customer) (core.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByCreds", arg0, arg1)
	ret0, _ := ret[0].(core.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByCreds indicates an expected call of GetByCreds.
func (mr *MockCustomerRepositoryMockRecorder) GetByCreds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByCreds", reflect.TypeOf((*MockCustomerRepository)(nil).GetByCreds), arg0, arg1)
}

// GetById mocks base method.
func (m *MockCustomerRepository) GetById(arg0 context.Context, arg1 core.Customer) (core.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0, arg1)
	ret0, _ := ret[0].(core.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockCustomerRepositoryMockRecorder) GetById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockCustomerRepository)(nil).GetById), arg0, arg1)
}

// Update mocks base method.
func (m *MockCustomerRepository) Update(arg0 context.Context, arg1 core.Customer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCustomerRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCustomerRepository)(nil).Update), arg0, arg1)
}

// UpdateEmailConfirmation mocks base method.
func (m *MockCustomerRepository) UpdateEmailConfirmation(arg0 context.Context, arg1 core.Customer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEmailConfirmation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEmailConfirmation indicates an expected call of UpdateEmailConfirmation.
func (mr *MockCustomerRepositoryMockRecorder) UpdateEmailConfirmation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEmailConfirmation", reflect.TypeOf((*MockCustomerRepository)(nil).UpdateEmailConfirmation), arg0, arg1)
}

// UpdatePassword mocks base method.
func (m *MockCustomerRepository) UpdatePassword(arg0 context.Context, arg1 core.Customer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockCustomerRepositoryMockRecorder) UpdatePassword(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockCustomerRepository)(nil).UpdatePassword), arg0, arg1)
}
