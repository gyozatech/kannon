// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	apiv1 "github.com/ludusrusso/kannon/proto/kannon/admin/apiv1"

	mock "github.com/stretchr/testify/mock"
)

// ApiServer is an autogenerated mock type for the ApiServer type
type ApiServer struct {
	mock.Mock
}

type ApiServer_Expecter struct {
	mock *mock.Mock
}

func (_m *ApiServer) EXPECT() *ApiServer_Expecter {
	return &ApiServer_Expecter{mock: &_m.Mock}
}

// CreateDomain provides a mock function with given fields: _a0, _a1
func (_m *ApiServer) CreateDomain(_a0 context.Context, _a1 *apiv1.CreateDomainRequest) (*apiv1.Domain, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *apiv1.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *apiv1.CreateDomainRequest) *apiv1.Domain); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *apiv1.CreateDomainRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApiServer_CreateDomain_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateDomain'
type ApiServer_CreateDomain_Call struct {
	*mock.Call
}

// CreateDomain is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *apiv1.CreateDomainRequest
func (_e *ApiServer_Expecter) CreateDomain(_a0 interface{}, _a1 interface{}) *ApiServer_CreateDomain_Call {
	return &ApiServer_CreateDomain_Call{Call: _e.mock.On("CreateDomain", _a0, _a1)}
}

func (_c *ApiServer_CreateDomain_Call) Run(run func(_a0 context.Context, _a1 *apiv1.CreateDomainRequest)) *ApiServer_CreateDomain_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*apiv1.CreateDomainRequest))
	})
	return _c
}

func (_c *ApiServer_CreateDomain_Call) Return(_a0 *apiv1.Domain, _a1 error) *ApiServer_CreateDomain_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// CreateTemplate provides a mock function with given fields: _a0, _a1
func (_m *ApiServer) CreateTemplate(_a0 context.Context, _a1 *apiv1.CreateTemplateReq) (*apiv1.CreateTemplateRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *apiv1.CreateTemplateRes
	if rf, ok := ret.Get(0).(func(context.Context, *apiv1.CreateTemplateReq) *apiv1.CreateTemplateRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.CreateTemplateRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *apiv1.CreateTemplateReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApiServer_CreateTemplate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateTemplate'
type ApiServer_CreateTemplate_Call struct {
	*mock.Call
}

// CreateTemplate is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *apiv1.CreateTemplateReq
func (_e *ApiServer_Expecter) CreateTemplate(_a0 interface{}, _a1 interface{}) *ApiServer_CreateTemplate_Call {
	return &ApiServer_CreateTemplate_Call{Call: _e.mock.On("CreateTemplate", _a0, _a1)}
}

func (_c *ApiServer_CreateTemplate_Call) Run(run func(_a0 context.Context, _a1 *apiv1.CreateTemplateReq)) *ApiServer_CreateTemplate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*apiv1.CreateTemplateReq))
	})
	return _c
}

func (_c *ApiServer_CreateTemplate_Call) Return(_a0 *apiv1.CreateTemplateRes, _a1 error) *ApiServer_CreateTemplate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// DeleteTemplate provides a mock function with given fields: _a0, _a1
func (_m *ApiServer) DeleteTemplate(_a0 context.Context, _a1 *apiv1.DeleteTemplateReq) (*apiv1.DeleteTemplateRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *apiv1.DeleteTemplateRes
	if rf, ok := ret.Get(0).(func(context.Context, *apiv1.DeleteTemplateReq) *apiv1.DeleteTemplateRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.DeleteTemplateRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *apiv1.DeleteTemplateReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApiServer_DeleteTemplate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteTemplate'
type ApiServer_DeleteTemplate_Call struct {
	*mock.Call
}

// DeleteTemplate is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *apiv1.DeleteTemplateReq
func (_e *ApiServer_Expecter) DeleteTemplate(_a0 interface{}, _a1 interface{}) *ApiServer_DeleteTemplate_Call {
	return &ApiServer_DeleteTemplate_Call{Call: _e.mock.On("DeleteTemplate", _a0, _a1)}
}

func (_c *ApiServer_DeleteTemplate_Call) Run(run func(_a0 context.Context, _a1 *apiv1.DeleteTemplateReq)) *ApiServer_DeleteTemplate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*apiv1.DeleteTemplateReq))
	})
	return _c
}

func (_c *ApiServer_DeleteTemplate_Call) Return(_a0 *apiv1.DeleteTemplateRes, _a1 error) *ApiServer_DeleteTemplate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetDomains provides a mock function with given fields: _a0, _a1
func (_m *ApiServer) GetDomains(_a0 context.Context, _a1 *apiv1.GetDomainsReq) (*apiv1.GetDomainsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *apiv1.GetDomainsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *apiv1.GetDomainsReq) *apiv1.GetDomainsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.GetDomainsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *apiv1.GetDomainsReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApiServer_GetDomains_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDomains'
type ApiServer_GetDomains_Call struct {
	*mock.Call
}

// GetDomains is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *apiv1.GetDomainsReq
func (_e *ApiServer_Expecter) GetDomains(_a0 interface{}, _a1 interface{}) *ApiServer_GetDomains_Call {
	return &ApiServer_GetDomains_Call{Call: _e.mock.On("GetDomains", _a0, _a1)}
}

func (_c *ApiServer_GetDomains_Call) Run(run func(_a0 context.Context, _a1 *apiv1.GetDomainsReq)) *ApiServer_GetDomains_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*apiv1.GetDomainsReq))
	})
	return _c
}

func (_c *ApiServer_GetDomains_Call) Return(_a0 *apiv1.GetDomainsResponse, _a1 error) *ApiServer_GetDomains_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTemplate provides a mock function with given fields: _a0, _a1
func (_m *ApiServer) GetTemplate(_a0 context.Context, _a1 *apiv1.GetTemplateReq) (*apiv1.GetTemplateRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *apiv1.GetTemplateRes
	if rf, ok := ret.Get(0).(func(context.Context, *apiv1.GetTemplateReq) *apiv1.GetTemplateRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.GetTemplateRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *apiv1.GetTemplateReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApiServer_GetTemplate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTemplate'
type ApiServer_GetTemplate_Call struct {
	*mock.Call
}

// GetTemplate is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *apiv1.GetTemplateReq
func (_e *ApiServer_Expecter) GetTemplate(_a0 interface{}, _a1 interface{}) *ApiServer_GetTemplate_Call {
	return &ApiServer_GetTemplate_Call{Call: _e.mock.On("GetTemplate", _a0, _a1)}
}

func (_c *ApiServer_GetTemplate_Call) Run(run func(_a0 context.Context, _a1 *apiv1.GetTemplateReq)) *ApiServer_GetTemplate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*apiv1.GetTemplateReq))
	})
	return _c
}

func (_c *ApiServer_GetTemplate_Call) Return(_a0 *apiv1.GetTemplateRes, _a1 error) *ApiServer_GetTemplate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetTemplates provides a mock function with given fields: _a0, _a1
func (_m *ApiServer) GetTemplates(_a0 context.Context, _a1 *apiv1.GetTemplatesReq) (*apiv1.GetTemplatesRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *apiv1.GetTemplatesRes
	if rf, ok := ret.Get(0).(func(context.Context, *apiv1.GetTemplatesReq) *apiv1.GetTemplatesRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.GetTemplatesRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *apiv1.GetTemplatesReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApiServer_GetTemplates_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTemplates'
type ApiServer_GetTemplates_Call struct {
	*mock.Call
}

// GetTemplates is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *apiv1.GetTemplatesReq
func (_e *ApiServer_Expecter) GetTemplates(_a0 interface{}, _a1 interface{}) *ApiServer_GetTemplates_Call {
	return &ApiServer_GetTemplates_Call{Call: _e.mock.On("GetTemplates", _a0, _a1)}
}

func (_c *ApiServer_GetTemplates_Call) Run(run func(_a0 context.Context, _a1 *apiv1.GetTemplatesReq)) *ApiServer_GetTemplates_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*apiv1.GetTemplatesReq))
	})
	return _c
}

func (_c *ApiServer_GetTemplates_Call) Return(_a0 *apiv1.GetTemplatesRes, _a1 error) *ApiServer_GetTemplates_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// RegenerateDomainKey provides a mock function with given fields: _a0, _a1
func (_m *ApiServer) RegenerateDomainKey(_a0 context.Context, _a1 *apiv1.RegenerateDomainKeyRequest) (*apiv1.Domain, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *apiv1.Domain
	if rf, ok := ret.Get(0).(func(context.Context, *apiv1.RegenerateDomainKeyRequest) *apiv1.Domain); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *apiv1.RegenerateDomainKeyRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApiServer_RegenerateDomainKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegenerateDomainKey'
type ApiServer_RegenerateDomainKey_Call struct {
	*mock.Call
}

// RegenerateDomainKey is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *apiv1.RegenerateDomainKeyRequest
func (_e *ApiServer_Expecter) RegenerateDomainKey(_a0 interface{}, _a1 interface{}) *ApiServer_RegenerateDomainKey_Call {
	return &ApiServer_RegenerateDomainKey_Call{Call: _e.mock.On("RegenerateDomainKey", _a0, _a1)}
}

func (_c *ApiServer_RegenerateDomainKey_Call) Run(run func(_a0 context.Context, _a1 *apiv1.RegenerateDomainKeyRequest)) *ApiServer_RegenerateDomainKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*apiv1.RegenerateDomainKeyRequest))
	})
	return _c
}

func (_c *ApiServer_RegenerateDomainKey_Call) Return(_a0 *apiv1.Domain, _a1 error) *ApiServer_RegenerateDomainKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// UpdateTemplate provides a mock function with given fields: _a0, _a1
func (_m *ApiServer) UpdateTemplate(_a0 context.Context, _a1 *apiv1.UpdateTemplateReq) (*apiv1.UpdateTemplateRes, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *apiv1.UpdateTemplateRes
	if rf, ok := ret.Get(0).(func(context.Context, *apiv1.UpdateTemplateReq) *apiv1.UpdateTemplateRes); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*apiv1.UpdateTemplateRes)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *apiv1.UpdateTemplateReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ApiServer_UpdateTemplate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateTemplate'
type ApiServer_UpdateTemplate_Call struct {
	*mock.Call
}

// UpdateTemplate is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *apiv1.UpdateTemplateReq
func (_e *ApiServer_Expecter) UpdateTemplate(_a0 interface{}, _a1 interface{}) *ApiServer_UpdateTemplate_Call {
	return &ApiServer_UpdateTemplate_Call{Call: _e.mock.On("UpdateTemplate", _a0, _a1)}
}

func (_c *ApiServer_UpdateTemplate_Call) Run(run func(_a0 context.Context, _a1 *apiv1.UpdateTemplateReq)) *ApiServer_UpdateTemplate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*apiv1.UpdateTemplateReq))
	})
	return _c
}

func (_c *ApiServer_UpdateTemplate_Call) Return(_a0 *apiv1.UpdateTemplateRes, _a1 error) *ApiServer_UpdateTemplate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewApiServer interface {
	mock.TestingT
	Cleanup(func())
}

// NewApiServer creates a new instance of ApiServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewApiServer(t mockConstructorTestingTNewApiServer) *ApiServer {
	mock := &ApiServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
