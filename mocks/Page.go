// Code generated by mockery v2.43.1. DO NOT EDIT.

package mocks

import (
	core "github.com/johnfercher/maroto/v2/pkg/core"
	entity "github.com/johnfercher/maroto/v2/pkg/core/entity"

	mock "github.com/stretchr/testify/mock"

	node "github.com/johnfercher/go-tree/node"
)

// Page is an autogenerated mock type for the Page type
type Page struct {
	mock.Mock
}

type Page_Expecter struct {
	mock *mock.Mock
}

func (_m *Page) EXPECT() *Page_Expecter {
	return &Page_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: rows
func (_m *Page) Add(rows ...core.Row) core.Page {
	_va := make([]interface{}, len(rows))
	for _i := range rows {
		_va[_i] = rows[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 core.Page
	if rf, ok := ret.Get(0).(func(...core.Row) core.Page); ok {
		r0 = rf(rows...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Page)
		}
	}

	return r0
}

// Page_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type Page_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - rows ...core.Row
func (_e *Page_Expecter) Add(rows ...interface{}) *Page_Add_Call {
	return &Page_Add_Call{Call: _e.mock.On("Add",
		append([]interface{}{}, rows...)...)}
}

func (_c *Page_Add_Call) Run(run func(rows ...core.Row)) *Page_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]core.Row, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(core.Row)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *Page_Add_Call) Return(_a0 core.Page) *Page_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Page_Add_Call) RunAndReturn(run func(...core.Row) core.Page) *Page_Add_Call {
	_c.Call.Return(run)
	return _c
}

// GetNumber provides a mock function with given fields:
func (_m *Page) GetNumber() int {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetNumber")
	}

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Page_GetNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNumber'
type Page_GetNumber_Call struct {
	*mock.Call
}

// GetNumber is a helper method to define mock.On call
func (_e *Page_Expecter) GetNumber() *Page_GetNumber_Call {
	return &Page_GetNumber_Call{Call: _e.mock.On("GetNumber")}
}

func (_c *Page_GetNumber_Call) Run(run func()) *Page_GetNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Page_GetNumber_Call) Return(_a0 int) *Page_GetNumber_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Page_GetNumber_Call) RunAndReturn(run func() int) *Page_GetNumber_Call {
	_c.Call.Return(run)
	return _c
}

// GetRows provides a mock function with given fields:
func (_m *Page) GetRows() []core.Row {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetRows")
	}

	var r0 []core.Row
	if rf, ok := ret.Get(0).(func() []core.Row); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]core.Row)
		}
	}

	return r0
}

// Page_GetRows_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRows'
type Page_GetRows_Call struct {
	*mock.Call
}

// GetRows is a helper method to define mock.On call
func (_e *Page_Expecter) GetRows() *Page_GetRows_Call {
	return &Page_GetRows_Call{Call: _e.mock.On("GetRows")}
}

func (_c *Page_GetRows_Call) Run(run func()) *Page_GetRows_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Page_GetRows_Call) Return(_a0 []core.Row) *Page_GetRows_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Page_GetRows_Call) RunAndReturn(run func() []core.Row) *Page_GetRows_Call {
	_c.Call.Return(run)
	return _c
}

// GetStructure provides a mock function with given fields:
func (_m *Page) GetStructure() *node.Node[core.Structure] {
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

// Page_GetStructure_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetStructure'
type Page_GetStructure_Call struct {
	*mock.Call
}

// GetStructure is a helper method to define mock.On call
func (_e *Page_Expecter) GetStructure() *Page_GetStructure_Call {
	return &Page_GetStructure_Call{Call: _e.mock.On("GetStructure")}
}

func (_c *Page_GetStructure_Call) Run(run func()) *Page_GetStructure_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Page_GetStructure_Call) Return(_a0 *node.Node[core.Structure]) *Page_GetStructure_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Page_GetStructure_Call) RunAndReturn(run func() *node.Node[core.Structure]) *Page_GetStructure_Call {
	_c.Call.Return(run)
	return _c
}

// Render provides a mock function with given fields: provider, cell
func (_m *Page) Render(provider core.Provider, cell entity.Cell) {
	_m.Called(provider, cell)
}

// Page_Render_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Render'
type Page_Render_Call struct {
	*mock.Call
}

// Render is a helper method to define mock.On call
//   - provider core.Provider
//   - cell entity.Cell
func (_e *Page_Expecter) Render(provider interface{}, cell interface{}) *Page_Render_Call {
	return &Page_Render_Call{Call: _e.mock.On("Render", provider, cell)}
}

func (_c *Page_Render_Call) Run(run func(provider core.Provider, cell entity.Cell)) *Page_Render_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(core.Provider), args[1].(entity.Cell))
	})
	return _c
}

func (_c *Page_Render_Call) Return() *Page_Render_Call {
	_c.Call.Return()
	return _c
}

func (_c *Page_Render_Call) RunAndReturn(run func(core.Provider, entity.Cell)) *Page_Render_Call {
	_c.Call.Return(run)
	return _c
}

// SetConfig provides a mock function with given fields: config
func (_m *Page) SetConfig(config *entity.Config) {
	_m.Called(config)
}

// Page_SetConfig_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetConfig'
type Page_SetConfig_Call struct {
	*mock.Call
}

// SetConfig is a helper method to define mock.On call
//   - config *entity.Config
func (_e *Page_Expecter) SetConfig(config interface{}) *Page_SetConfig_Call {
	return &Page_SetConfig_Call{Call: _e.mock.On("SetConfig", config)}
}

func (_c *Page_SetConfig_Call) Run(run func(config *entity.Config)) *Page_SetConfig_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Config))
	})
	return _c
}

func (_c *Page_SetConfig_Call) Return() *Page_SetConfig_Call {
	_c.Call.Return()
	return _c
}

func (_c *Page_SetConfig_Call) RunAndReturn(run func(*entity.Config)) *Page_SetConfig_Call {
	_c.Call.Return(run)
	return _c
}

// SetNumber provides a mock function with given fields: number, total
func (_m *Page) SetNumber(number int, total int) {
	_m.Called(number, total)
}

// Page_SetNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetNumber'
type Page_SetNumber_Call struct {
	*mock.Call
}

// SetNumber is a helper method to define mock.On call
//   - number int
//   - total int
func (_e *Page_Expecter) SetNumber(number interface{}, total interface{}) *Page_SetNumber_Call {
	return &Page_SetNumber_Call{Call: _e.mock.On("SetNumber", number, total)}
}

func (_c *Page_SetNumber_Call) Run(run func(number int, total int)) *Page_SetNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int))
	})
	return _c
}

func (_c *Page_SetNumber_Call) Return() *Page_SetNumber_Call {
	_c.Call.Return()
	return _c
}

func (_c *Page_SetNumber_Call) RunAndReturn(run func(int, int)) *Page_SetNumber_Call {
	_c.Call.Return(run)
	return _c
}

// NewPage creates a new instance of Page. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPage(t interface {
	mock.TestingT
	Cleanup(func())
},
) *Page {
	mock := &Page{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
