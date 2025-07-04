// Code generated by mockery v2.42.0. DO NOT EDIT.

package mocks

import (
	extension "github.com/nh3000-org/maroto/v2/pkg/consts/extension"
	entity "github.com/nh3000-org/maroto/v2/pkg/core/entity"

	gofpdf "github.com/jung-kurt/gofpdf"

	mock "github.com/stretchr/testify/mock"

	props "github.com/nh3000-org/maroto/v2/pkg/props"

	uuid "github.com/google/uuid"
)

// Image is an autogenerated mock type for the Image type
type Image struct {
	mock.Mock
}

type Image_Expecter struct {
	mock *mock.Mock
}

func (_m *Image) EXPECT() *Image_Expecter {
	return &Image_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: img, cell, margins, prop, _a4, flow
func (_m *Image) Add(img *entity.Image, cell *entity.Cell, margins *entity.Margins, prop *props.Rect, _a4 extension.Type, flow bool) error {
	ret := _m.Called(img, cell, margins, prop, _a4, flow)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Image, *entity.Cell, *entity.Margins, *props.Rect, extension.Type, bool) error); ok {
		r0 = rf(img, cell, margins, prop, _a4, flow)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Image_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type Image_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - img *entity.Image
//   - cell *entity.Cell
//   - margins *entity.Margins
//   - prop *props.Rect
//   - _a4 extension.Type
//   - flow bool
func (_e *Image_Expecter) Add(img interface{}, cell interface{}, margins interface{}, prop interface{}, _a4 interface{}, flow interface{}) *Image_Add_Call {
	return &Image_Add_Call{Call: _e.mock.On("Add", img, cell, margins, prop, _a4, flow)}
}

func (_c *Image_Add_Call) Run(run func(img *entity.Image, cell *entity.Cell, margins *entity.Margins, prop *props.Rect, _a4 extension.Type, flow bool)) *Image_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Image), args[1].(*entity.Cell), args[2].(*entity.Margins), args[3].(*props.Rect), args[4].(extension.Type), args[5].(bool))
	})
	return _c
}

func (_c *Image_Add_Call) Return(_a0 error) *Image_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Image_Add_Call) RunAndReturn(run func(*entity.Image, *entity.Cell, *entity.Margins, *props.Rect, extension.Type, bool) error) *Image_Add_Call {
	_c.Call.Return(run)
	return _c
}

// GetImageInfo provides a mock function with given fields: img, _a1
func (_m *Image) GetImageInfo(img *entity.Image, _a1 extension.Type) (*gofpdf.ImageInfoType, uuid.UUID) {
	ret := _m.Called(img, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetImageInfo")
	}

	var r0 *gofpdf.ImageInfoType
	var r1 uuid.UUID
	if rf, ok := ret.Get(0).(func(*entity.Image, extension.Type) (*gofpdf.ImageInfoType, uuid.UUID)); ok {
		return rf(img, _a1)
	}
	if rf, ok := ret.Get(0).(func(*entity.Image, extension.Type) *gofpdf.ImageInfoType); ok {
		r0 = rf(img, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gofpdf.ImageInfoType)
		}
	}

	if rf, ok := ret.Get(1).(func(*entity.Image, extension.Type) uuid.UUID); ok {
		r1 = rf(img, _a1)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(uuid.UUID)
		}
	}

	return r0, r1
}

// Image_GetImageInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetImageInfo'
type Image_GetImageInfo_Call struct {
	*mock.Call
}

// GetImageInfo is a helper method to define mock.On call
//   - img *entity.Image
//   - _a1 extension.Type
func (_e *Image_Expecter) GetImageInfo(img interface{}, _a1 interface{}) *Image_GetImageInfo_Call {
	return &Image_GetImageInfo_Call{Call: _e.mock.On("GetImageInfo", img, _a1)}
}

func (_c *Image_GetImageInfo_Call) Run(run func(img *entity.Image, _a1 extension.Type)) *Image_GetImageInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Image), args[1].(extension.Type))
	})
	return _c
}

func (_c *Image_GetImageInfo_Call) Return(_a0 *gofpdf.ImageInfoType, _a1 uuid.UUID) *Image_GetImageInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Image_GetImageInfo_Call) RunAndReturn(run func(*entity.Image, extension.Type) (*gofpdf.ImageInfoType, uuid.UUID)) *Image_GetImageInfo_Call {
	_c.Call.Return(run)
	return _c
}

// NewImage creates a new instance of Image. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewImage(t interface {
	mock.TestingT
	Cleanup(func())
},
) *Image {
	mock := &Image{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
