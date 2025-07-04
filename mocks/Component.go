// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	core "github.com/nh3000-org/maroto/v2/pkg/core"
	entity "github.com/nh3000-org/maroto/v2/pkg/core/entity"

	mock "github.com/stretchr/testify/mock"

	node "github.com/nh3000-org/go-tree/node"
)

// Component is an autogenerated mock type for the Component type
type Component struct {
	mock.Mock
}

type Component_Expecter struct {
	mock *mock.Mock
}

func (_m *Component) EXPECT() *Component_Expecter {
	return &Component_Expecter{mock: &_m.Mock}
}

// GetHeight provides a mock function with given fields: provider, cell
func (_m *Component) GetHeight(provider core.Provider, cell *entity.Cell) float64 {
	ret := _m.Called(provider, cell)

	if len(ret) == 0 {
		panic("no return value specified for GetHeight")
	}

	var r0 float64
	if rf, ok := ret.Get(0).(func(core.Provider, *entity.Cell) float64); ok {
		r0 = rf(provider, cell)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// Component_GetHeight_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetHeight'
type Component_GetHeight_Call struct {
	*mock.Call
}

// GetHeight is a helper method to define mock.On call
//   - provider core.Provider
//   - cell *entity.Cell
func (_e *Component_Expecter) GetHeight(provider interface{}, cell interface{}) *Component_GetHeight_Call {
	return &Component_GetHeight_Call{Call: _e.mock.On("GetHeight", provider, cell)}
}

func (_c *Component_GetHeight_Call) Run(run func(provider core.Provider, cell *entity.Cell)) *Component_GetHeight_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(core.Provider), args[1].(*entity.Cell))
	})
	return _c
}

func (_c *Component_GetHeight_Call) Return(_a0 float64) *Component_GetHeight_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Component_GetHeight_Call) RunAndReturn(run func(core.Provider, *entity.Cell) float64) *Component_GetHeight_Call {
	_c.Call.Return(run)
	return _c
}

// GetStructure provides a mock function with given fields:
func (_m *Component) GetStructure() *node.Node[core.Structure] {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetStructure")
	}

	var r0 *node.Node[core.Structure]
	if rf, ok := ret.Get(0).(func() *node.Node[core.Structure]); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*node.Node[core.Structure])
		}
	}

	return r0
}

// Component_GetStructure_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStructure'
type Component_GetStructure_Call struct {
	*mock.Call
}

// GetStructure is a helper method to define mock.On call
func (_e *Component_Expecter) GetStructure() *Component_GetStructure_Call {
	return &Component_GetStructure_Call{Call: _e.mock.On("GetStructure")}
}

func (_c *Component_GetStructure_Call) Run(run func()) *Component_GetStructure_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Component_GetStructure_Call) Return(_a0 *node.Node[core.Structure]) *Component_GetStructure_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Component_GetStructure_Call) RunAndReturn(run func() *node.Node[core.Structure]) *Component_GetStructure_Call {
	_c.Call.Return(run)
	return _c
}

// Render provides a mock function with given fields: provider, cell
func (_m *Component) Render(provider core.Provider, cell *entity.Cell) {
	_m.Called(provider, cell)
}

// Component_Render_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Render'
type Component_Render_Call struct {
	*mock.Call
}

// Render is a helper method to define mock.On call
//   - provider core.Provider
//   - cell *entity.Cell
func (_e *Component_Expecter) Render(provider interface{}, cell interface{}) *Component_Render_Call {
	return &Component_Render_Call{Call: _e.mock.On("Render", provider, cell)}
}

func (_c *Component_Render_Call) Run(run func(provider core.Provider, cell *entity.Cell)) *Component_Render_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(core.Provider), args[1].(*entity.Cell))
	})
	return _c
}

func (_c *Component_Render_Call) Return() *Component_Render_Call {
	_c.Call.Return()
	return _c
}

func (_c *Component_Render_Call) RunAndReturn(run func(core.Provider, *entity.Cell)) *Component_Render_Call {
	_c.Call.Return(run)
	return _c
}

// SetConfig provides a mock function with given fields: config
func (_m *Component) SetConfig(config *entity.Config) {
	_m.Called(config)
}

// Component_SetConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetConfig'
type Component_SetConfig_Call struct {
	*mock.Call
}

// SetConfig is a helper method to define mock.On call
//   - config *entity.Config
func (_e *Component_Expecter) SetConfig(config interface{}) *Component_SetConfig_Call {
	return &Component_SetConfig_Call{Call: _e.mock.On("SetConfig", config)}
}

func (_c *Component_SetConfig_Call) Run(run func(config *entity.Config)) *Component_SetConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Config))
	})
	return _c
}

func (_c *Component_SetConfig_Call) Return() *Component_SetConfig_Call {
	_c.Call.Return()
	return _c
}

func (_c *Component_SetConfig_Call) RunAndReturn(run func(*entity.Config)) *Component_SetConfig_Call {
	_c.Call.Return(run)
	return _c
}

// NewComponent creates a new instance of Component. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewComponent(t interface {
	mock.TestingT
	Cleanup(func())
},
) *Component {
	mock := &Component{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
