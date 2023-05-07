// Code generated by mockery v2.26.1. DO NOT EDIT.

package home

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MockHandler is an autogenerated mock type for the Handler type
type MockHandler struct {
	mock.Mock
}

type MockHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHandler) EXPECT() *MockHandler_Expecter {
	return &MockHandler_Expecter{mock: &_m.Mock}
}

// Home provides a mock function with given fields: w, r
func (_m *MockHandler) Home(w http.ResponseWriter, r *http.Request) {
	_m.Called(w, r)
}

// MockHandler_Home_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Home'
type MockHandler_Home_Call struct {
	*mock.Call
}

// Home is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *MockHandler_Expecter) Home(w interface{}, r interface{}) *MockHandler_Home_Call {
	return &MockHandler_Home_Call{Call: _e.mock.On("Home", w, r)}
}

func (_c *MockHandler_Home_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *MockHandler_Home_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockHandler_Home_Call) Return() *MockHandler_Home_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockHandler_Home_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request)) *MockHandler_Home_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockHandler creates a new instance of MockHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockHandler(t mockConstructorTestingTNewMockHandler) *MockHandler {
	mock := &MockHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}