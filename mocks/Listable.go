// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	core "github.com/nh3000-org/maroto/v2/pkg/core"

	mock "github.com/stretchr/testify/mock"
)

// Listable is an autogenerated mock type for the Listable type
type Listable struct {
	mock.Mock
}

type Listable_Expecter struct {
	mock *mock.Mock
}

func (_m *Listable) EXPECT() *Listable_Expecter {
	return &Listable_Expecter{mock: &_m.Mock}
}

// GetContent provides a mock function with given fields: i
func (_m *Listable) GetContent(i int) core.Row {
	ret := _m.Called(i)

	if len(ret) == 0 {
		panic("no return value specified for GetContent")
	}

	var r0 core.Row
	if rf, ok := ret.Get(0).(func(int) core.Row); ok {
		r0 = rf(i)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Row)
		}
	}

	return r0
}

// Listable_GetContent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetContent'
type Listable_GetContent_Call struct {
	*mock.Call
}

// GetContent is a helper method to define mock.On call
//   - i int
func (_e *Listable_Expecter) GetContent(i interface{}) *Listable_GetContent_Call {
	return &Listable_GetContent_Call{Call: _e.mock.On("GetContent", i)}
}

func (_c *Listable_GetContent_Call) Run(run func(i int)) *Listable_GetContent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return _c
}

func (_c *Listable_GetContent_Call) Return(_a0 core.Row) *Listable_GetContent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Listable_GetContent_Call) RunAndReturn(run func(int) core.Row) *Listable_GetContent_Call {
	_c.Call.Return(run)
	return _c
}

// GetHeader provides a mock function with given fields:
func (_m *Listable) GetHeader() core.Row {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetHeader")
	}

	var r0 core.Row
	if rf, ok := ret.Get(0).(func() core.Row); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Row)
		}
	}

	return r0
}

// Listable_GetHeader_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHeader'
type Listable_GetHeader_Call struct {
	*mock.Call
}

// GetHeader is a helper method to define mock.On call
func (_e *Listable_Expecter) GetHeader() *Listable_GetHeader_Call {
	return &Listable_GetHeader_Call{Call: _e.mock.On("GetHeader")}
}

func (_c *Listable_GetHeader_Call) Run(run func()) *Listable_GetHeader_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Listable_GetHeader_Call) Return(_a0 core.Row) *Listable_GetHeader_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Listable_GetHeader_Call) RunAndReturn(run func() core.Row) *Listable_GetHeader_Call {
	_c.Call.Return(run)
	return _c
}

// NewListable creates a new instance of Listable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewListable(t interface {
	mock.TestingT
	Cleanup(func())
},
) *Listable {
	mock := &Listable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
