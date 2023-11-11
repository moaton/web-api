// Code generated by mockery v2.36.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Token is an autogenerated mock type for the Token type
type Token struct {
	mock.Mock
}

// CreateAccessToken provides a mock function with given fields: id, email
func (_m *Token) CreateAccessToken(id int64, email string) (string, error) {
	ret := _m.Called(id, email)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int64, string) (string, error)); ok {
		return rf(id, email)
	}
	if rf, ok := ret.Get(0).(func(int64, string) string); ok {
		r0 = rf(id, email)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int64, string) error); ok {
		r1 = rf(id, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRefreshToken provides a mock function with given fields: id
func (_m *Token) CreateRefreshToken(id int64) (string, error) {
	ret := _m.Called(id)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(int64) (string, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int64) string); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(int64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExtractIDFromToken provides a mock function with given fields: requestToken
func (_m *Token) ExtractIDFromToken(requestToken string) (int64, error) {
	ret := _m.Called(requestToken)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(requestToken)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(requestToken)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(requestToken)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsAuthorized provides a mock function with given fields: _a0
func (_m *Token) IsAuthorized(_a0 string) (bool, error) {
	ret := _m.Called(_a0)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewToken creates a new instance of Token. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewToken(t interface {
	mock.TestingT
	Cleanup(func())
}) *Token {
	mock := &Token{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
