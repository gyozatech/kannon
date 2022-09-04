// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: kannon/admin/apiv1/adminapiv1.proto

package apiv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ApiClient is the client API for Api service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiClient interface {
	GetDomains(ctx context.Context, in *GetDomainsReq, opts ...grpc.CallOption) (*GetDomainsResponse, error)
	CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*Domain, error)
	RegenerateDomainKey(ctx context.Context, in *RegenerateDomainKeyRequest, opts ...grpc.CallOption) (*Domain, error)
	CreateTemplate(ctx context.Context, in *CreateTemplateReq, opts ...grpc.CallOption) (*CreateTemplateRes, error)
	UpdateTemplate(ctx context.Context, in *UpdateTemplateReq, opts ...grpc.CallOption) (*UpdateTemplateRes, error)
	DeleteTemplate(ctx context.Context, in *DeleteTemplateReq, opts ...grpc.CallOption) (*DeleteTemplateRes, error)
	GetTemplate(ctx context.Context, in *GetTemplateReq, opts ...grpc.CallOption) (*GetTemplateRes, error)
	GetTemplates(ctx context.Context, in *GetTemplatesReq, opts ...grpc.CallOption) (*GetTemplatesRes, error)
}

type apiClient struct {
	cc grpc.ClientConnInterface
}

func NewApiClient(cc grpc.ClientConnInterface) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) GetDomains(ctx context.Context, in *GetDomainsReq, opts ...grpc.CallOption) (*GetDomainsResponse, error) {
	out := new(GetDomainsResponse)
	err := c.cc.Invoke(ctx, "/pkg.kannon.admin.apiv1.Api/GetDomains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) CreateDomain(ctx context.Context, in *CreateDomainRequest, opts ...grpc.CallOption) (*Domain, error) {
	out := new(Domain)
	err := c.cc.Invoke(ctx, "/pkg.kannon.admin.apiv1.Api/CreateDomain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) RegenerateDomainKey(ctx context.Context, in *RegenerateDomainKeyRequest, opts ...grpc.CallOption) (*Domain, error) {
	out := new(Domain)
	err := c.cc.Invoke(ctx, "/pkg.kannon.admin.apiv1.Api/RegenerateDomainKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) CreateTemplate(ctx context.Context, in *CreateTemplateReq, opts ...grpc.CallOption) (*CreateTemplateRes, error) {
	out := new(CreateTemplateRes)
	err := c.cc.Invoke(ctx, "/pkg.kannon.admin.apiv1.Api/CreateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) UpdateTemplate(ctx context.Context, in *UpdateTemplateReq, opts ...grpc.CallOption) (*UpdateTemplateRes, error) {
	out := new(UpdateTemplateRes)
	err := c.cc.Invoke(ctx, "/pkg.kannon.admin.apiv1.Api/UpdateTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeleteTemplate(ctx context.Context, in *DeleteTemplateReq, opts ...grpc.CallOption) (*DeleteTemplateRes, error) {
	out := new(DeleteTemplateRes)
	err := c.cc.Invoke(ctx, "/pkg.kannon.admin.apiv1.Api/DeleteTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetTemplate(ctx context.Context, in *GetTemplateReq, opts ...grpc.CallOption) (*GetTemplateRes, error) {
	out := new(GetTemplateRes)
	err := c.cc.Invoke(ctx, "/pkg.kannon.admin.apiv1.Api/GetTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetTemplates(ctx context.Context, in *GetTemplatesReq, opts ...grpc.CallOption) (*GetTemplatesRes, error) {
	out := new(GetTemplatesRes)
	err := c.cc.Invoke(ctx, "/pkg.kannon.admin.apiv1.Api/GetTemplates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiServer is the server API for Api service.
// All implementations should embed UnimplementedApiServer
// for forward compatibility
type ApiServer interface {
	GetDomains(context.Context, *GetDomainsReq) (*GetDomainsResponse, error)
	CreateDomain(context.Context, *CreateDomainRequest) (*Domain, error)
	RegenerateDomainKey(context.Context, *RegenerateDomainKeyRequest) (*Domain, error)
	CreateTemplate(context.Context, *CreateTemplateReq) (*CreateTemplateRes, error)
	UpdateTemplate(context.Context, *UpdateTemplateReq) (*UpdateTemplateRes, error)
	DeleteTemplate(context.Context, *DeleteTemplateReq) (*DeleteTemplateRes, error)
	GetTemplate(context.Context, *GetTemplateReq) (*GetTemplateRes, error)
	GetTemplates(context.Context, *GetTemplatesReq) (*GetTemplatesRes, error)
}

// UnimplementedApiServer should be embedded to have forward compatible implementations.
type UnimplementedApiServer struct {
}

func (UnimplementedApiServer) GetDomains(context.Context, *GetDomainsReq) (*GetDomainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDomains not implemented")
}
func (UnimplementedApiServer) CreateDomain(context.Context, *CreateDomainRequest) (*Domain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDomain not implemented")
}
func (UnimplementedApiServer) RegenerateDomainKey(context.Context, *RegenerateDomainKeyRequest) (*Domain, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegenerateDomainKey not implemented")
}
func (UnimplementedApiServer) CreateTemplate(context.Context, *CreateTemplateReq) (*CreateTemplateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTemplate not implemented")
}
func (UnimplementedApiServer) UpdateTemplate(context.Context, *UpdateTemplateReq) (*UpdateTemplateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTemplate not implemented")
}
func (UnimplementedApiServer) DeleteTemplate(context.Context, *DeleteTemplateReq) (*DeleteTemplateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTemplate not implemented")
}
func (UnimplementedApiServer) GetTemplate(context.Context, *GetTemplateReq) (*GetTemplateRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplate not implemented")
}
func (UnimplementedApiServer) GetTemplates(context.Context, *GetTemplatesReq) (*GetTemplatesRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTemplates not implemented")
}

// UnsafeApiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServer will
// result in compilation errors.
type UnsafeApiServer interface {
	mustEmbedUnimplementedApiServer()
}

func RegisterApiServer(s grpc.ServiceRegistrar, srv ApiServer) {
	s.RegisterService(&Api_ServiceDesc, srv)
}

func _Api_GetDomains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDomainsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetDomains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.kannon.admin.apiv1.Api/GetDomains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetDomains(ctx, req.(*GetDomainsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_CreateDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).CreateDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.kannon.admin.apiv1.Api/CreateDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).CreateDomain(ctx, req.(*CreateDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_RegenerateDomainKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegenerateDomainKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).RegenerateDomainKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.kannon.admin.apiv1.Api/RegenerateDomainKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).RegenerateDomainKey(ctx, req.(*RegenerateDomainKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_CreateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTemplateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).CreateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.kannon.admin.apiv1.Api/CreateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).CreateTemplate(ctx, req.(*CreateTemplateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_UpdateTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTemplateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).UpdateTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.kannon.admin.apiv1.Api/UpdateTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).UpdateTemplate(ctx, req.(*UpdateTemplateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_DeleteTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTemplateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).DeleteTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.kannon.admin.apiv1.Api/DeleteTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).DeleteTemplate(ctx, req.(*DeleteTemplateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.kannon.admin.apiv1.Api/GetTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetTemplate(ctx, req.(*GetTemplateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Api_GetTemplates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTemplatesReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServer).GetTemplates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pkg.kannon.admin.apiv1.Api/GetTemplates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServer).GetTemplates(ctx, req.(*GetTemplatesReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Api_ServiceDesc is the grpc.ServiceDesc for Api service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Api_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pkg.kannon.admin.apiv1.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDomains",
			Handler:    _Api_GetDomains_Handler,
		},
		{
			MethodName: "CreateDomain",
			Handler:    _Api_CreateDomain_Handler,
		},
		{
			MethodName: "RegenerateDomainKey",
			Handler:    _Api_RegenerateDomainKey_Handler,
		},
		{
			MethodName: "CreateTemplate",
			Handler:    _Api_CreateTemplate_Handler,
		},
		{
			MethodName: "UpdateTemplate",
			Handler:    _Api_UpdateTemplate_Handler,
		},
		{
			MethodName: "DeleteTemplate",
			Handler:    _Api_DeleteTemplate_Handler,
		},
		{
			MethodName: "GetTemplate",
			Handler:    _Api_GetTemplate_Handler,
		},
		{
			MethodName: "GetTemplates",
			Handler:    _Api_GetTemplates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kannon/admin/apiv1/adminapiv1.proto",
}
