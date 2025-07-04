// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	cellwriter "github.com/nh3000-org/maroto/v2/internal/providers/gofpdf/cellwriter"
	entity "github.com/nh3000-org/maroto/v2/pkg/core/entity"

	mock "github.com/stretchr/testify/mock"

	props "github.com/nh3000-org/maroto/v2/pkg/props"
)

// CellWriter is an autogenerated mock type for the CellWriter type
type CellWriter struct {
	mock.Mock
}

type CellWriter_Expecter struct {
	mock *mock.Mock
}

func (_m *CellWriter) EXPECT() *CellWriter_Expecter {
	return &CellWriter_Expecter{mock: &_m.Mock}
}

// Apply provides a mock function with given fields: width, height, config, prop
func (_m *CellWriter) Apply(width float64, height float64, config *entity.Config, prop *props.Cell) {
	_m.Called(width, height, config, prop)
}

// CellWriter_Apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Apply'
type CellWriter_Apply_Call struct {
	*mock.Call
}

// Apply is a helper method to define mock.On call
//   - width float64
//   - height float64
//   - config *entity.Config
//   - prop *props.Cell
func (_e *CellWriter_Expecter) Apply(width interface{}, height interface{}, config interface{}, prop interface{}) *CellWriter_Apply_Call {
	return &CellWriter_Apply_Call{Call: _e.mock.On("Apply", width, height, config, prop)}
}

func (_c *CellWriter_Apply_Call) Run(run func(width float64, height float64, config *entity.Config, prop *props.Cell)) *CellWriter_Apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64), args[1].(float64), args[2].(*entity.Config), args[3].(*props.Cell))
	})
	return _c
}

func (_c *CellWriter_Apply_Call) Return() *CellWriter_Apply_Call {
	_c.Call.Return()
	return _c
}

func (_c *CellWriter_Apply_Call) RunAndReturn(run func(float64, float64, *entity.Config, *props.Cell)) *CellWriter_Apply_Call {
	_c.Call.Return(run)
	return _c
}

// GetName provides a mock function with given fields:
func (_m *CellWriter) GetName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// CellWriter_GetName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetName'
type CellWriter_GetName_Call struct {
	*mock.Call
}

// GetName is a helper method to define mock.On call
func (_e *CellWriter_Expecter) GetName() *CellWriter_GetName_Call {
	return &CellWriter_GetName_Call{Call: _e.mock.On("GetName")}
}

func (_c *CellWriter_GetName_Call) Run(run func()) *CellWriter_GetName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CellWriter_GetName_Call) Return(_a0 string) *CellWriter_GetName_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CellWriter_GetName_Call) RunAndReturn(run func() string) *CellWriter_GetName_Call {
	_c.Call.Return(run)
	return _c
}

// GetNext provides a mock function with given fields:
func (_m *CellWriter) GetNext() cellwriter.CellWriter {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNext")
	}

	var r0 cellwriter.CellWriter
	if rf, ok := ret.Get(0).(func() cellwriter.CellWriter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cellwriter.CellWriter)
		}
	}

	return r0
}

// CellWriter_GetNext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNext'
type CellWriter_GetNext_Call struct {
	*mock.Call
}

// GetNext is a helper method to define mock.On call
func (_e *CellWriter_Expecter) GetNext() *CellWriter_GetNext_Call {
	return &CellWriter_GetNext_Call{Call: _e.mock.On("GetNext")}
}

func (_c *CellWriter_GetNext_Call) Run(run func()) *CellWriter_GetNext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *CellWriter_GetNext_Call) Return(_a0 cellwriter.CellWriter) *CellWriter_GetNext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *CellWriter_GetNext_Call) RunAndReturn(run func() cellwriter.CellWriter) *CellWriter_GetNext_Call {
	_c.Call.Return(run)
	return _c
}

// SetNext provides a mock function with given fields: next
func (_m *CellWriter) SetNext(next cellwriter.CellWriter) {
	_m.Called(next)
}

// CellWriter_SetNext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetNext'
type CellWriter_SetNext_Call struct {
	*mock.Call
}

// SetNext is a helper method to define mock.On call
//   - next cellwriter.CellWriter
func (_e *CellWriter_Expecter) SetNext(next interface{}) *CellWriter_SetNext_Call {
	return &CellWriter_SetNext_Call{Call: _e.mock.On("SetNext", next)}
}

func (_c *CellWriter_SetNext_Call) Run(run func(next cellwriter.CellWriter)) *CellWriter_SetNext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(cellwriter.CellWriter))
	})
	return _c
}

func (_c *CellWriter_SetNext_Call) Return() *CellWriter_SetNext_Call {
	_c.Call.Return()
	return _c
}

func (_c *CellWriter_SetNext_Call) RunAndReturn(run func(cellwriter.CellWriter)) *CellWriter_SetNext_Call {
	_c.Call.Return(run)
	return _c
}

// NewCellWriter creates a new instance of CellWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCellWriter(t interface {
	mock.TestingT
	Cleanup(func())
},
) *CellWriter {
	mock := &CellWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
