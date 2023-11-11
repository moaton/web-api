// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

// Auth provides a mock function with given fields: w, r
func (_m *Handler) Auth(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// CreateUser provides a mock function with given fields: w, r
func (_m *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// DeleteUser provides a mock function with given fields: w, r
func (_m *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// Refresh provides a mock function with given fields: w, r
func (_m *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// UpdateUser provides a mock function with given fields: w, r
func (_m *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
