// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	extension "github.com/nh3000-org/maroto/v2/pkg/consts/extension"
	entity "github.com/nh3000-org/maroto/v2/pkg/core/entity"

	mock "github.com/stretchr/testify/mock"
)

// Cache is an autogenerated mock type for the Cache type
type Cache struct {
	mock.Mock
}

type Cache_Expecter struct {
	mock *mock.Mock
}

func (_m *Cache) EXPECT() *Cache_Expecter {
	return &Cache_Expecter{mock: &_m.Mock}
}

// AddImage provides a mock function with given fields: value, image
func (_m *Cache) AddImage(value string, image *entity.Image) {
	_m.Called(value, image)
}

// Cache_AddImage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddImage'
type Cache_AddImage_Call struct {
	*mock.Call
}

// AddImage is a helper method to define mock.On call
//   - value string
//   - image *entity.Image
func (_e *Cache_Expecter) AddImage(value interface{}, image interface{}) *Cache_AddImage_Call {
	return &Cache_AddImage_Call{Call: _e.mock.On("AddImage", value, image)}
}

func (_c *Cache_AddImage_Call) Run(run func(value string, image *entity.Image)) *Cache_AddImage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(*entity.Image))
	})
	return _c
}

func (_c *Cache_AddImage_Call) Return() *Cache_AddImage_Call {
	_c.Call.Return()
	return _c
}

func (_c *Cache_AddImage_Call) RunAndReturn(run func(string, *entity.Image)) *Cache_AddImage_Call {
	_c.Call.Return(run)
	return _c
}

// GetImage provides a mock function with given fields: value, _a1
func (_m *Cache) GetImage(value string, _a1 extension.Type) (*entity.Image, error) {
	ret := _m.Called(value, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetImage")
	}

	var r0 *entity.Image
	var r1 error
	if rf, ok := ret.Get(0).(func(string, extension.Type) (*entity.Image, error)); ok {
		return rf(value, _a1)
	}
	if rf, ok := ret.Get(0).(func(string, extension.Type) *entity.Image); ok {
		r0 = rf(value, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Image)
		}
	}

	if rf, ok := ret.Get(1).(func(string, extension.Type) error); ok {
		r1 = rf(value, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Cache_GetImage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetImage'
type Cache_GetImage_Call struct {
	*mock.Call
}

// GetImage is a helper method to define mock.On call
//   - value string
//   - _a1 extension.Type
func (_e *Cache_Expecter) GetImage(value interface{}, _a1 interface{}) *Cache_GetImage_Call {
	return &Cache_GetImage_Call{Call: _e.mock.On("GetImage", value, _a1)}
}

func (_c *Cache_GetImage_Call) Run(run func(value string, _a1 extension.Type)) *Cache_GetImage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(extension.Type))
	})
	return _c
}

func (_c *Cache_GetImage_Call) Return(_a0 *entity.Image, _a1 error) *Cache_GetImage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Cache_GetImage_Call) RunAndReturn(run func(string, extension.Type) (*entity.Image, error)) *Cache_GetImage_Call {
	_c.Call.Return(run)
	return _c
}

// LoadImage provides a mock function with given fields: value, _a1
func (_m *Cache) LoadImage(value string, _a1 extension.Type) error {
	ret := _m.Called(value, _a1)

	if len(ret) == 0 {
		panic("no return value specified for LoadImage")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, extension.Type) error); ok {
		r0 = rf(value, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Cache_LoadImage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadImage'
type Cache_LoadImage_Call struct {
	*mock.Call
}

// LoadImage is a helper method to define mock.On call
//   - value string
//   - _a1 extension.Type
func (_e *Cache_Expecter) LoadImage(value interface{}, _a1 interface{}) *Cache_LoadImage_Call {
	return &Cache_LoadImage_Call{Call: _e.mock.On("LoadImage", value, _a1)}
}

func (_c *Cache_LoadImage_Call) Run(run func(value string, _a1 extension.Type)) *Cache_LoadImage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(extension.Type))
	})
	return _c
}

func (_c *Cache_LoadImage_Call) Return(_a0 error) *Cache_LoadImage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Cache_LoadImage_Call) RunAndReturn(run func(string, extension.Type) error) *Cache_LoadImage_Call {
	_c.Call.Return(run)
	return _c
}

// NewCache creates a new instance of Cache. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCache(t interface {
	mock.TestingT
	Cleanup(func())
},
) *Cache {
	mock := &Cache{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
