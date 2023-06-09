// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Prompter is an autogenerated mock type for the Prompter type
type Prompter struct {
	mock.Mock
}

// IsTerminal provides a mock function with given fields:
func (_m *Prompter) IsTerminal() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PasswordPrompt provides a mock function with given fields: _a0
func (_m *Prompter) PasswordPrompt(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Prompt provides a mock function with given fields: _a0
func (_m *Prompter) Prompt(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewPrompter interface {
	mock.TestingT
	Cleanup(func())
}

// NewPrompter creates a new instance of Prompter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPrompter(t mockConstructorTestingTNewPrompter) *Prompter {
	mock := &Prompter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
