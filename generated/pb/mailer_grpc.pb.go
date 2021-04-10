// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// MailerClient is the client API for Mailer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MailerClient interface {
	SendHTML(ctx context.Context, in *SendHTMLRequest, opts ...grpc.CallOption) (*SendResponse, error)
	SendTemplate(ctx context.Context, in *SendTemplateRequest, opts ...grpc.CallOption) (*SendResponse, error)
}

type mailerClient struct {
	cc grpc.ClientConnInterface
}

func NewMailerClient(cc grpc.ClientConnInterface) MailerClient {
	return &mailerClient{cc}
}

func (c *mailerClient) SendHTML(ctx context.Context, in *SendHTMLRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/kannon.Mailer/SendHTML", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mailerClient) SendTemplate(ctx context.Context, in *SendTemplateRequest, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/kannon.Mailer/SendTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MailerServer is the server API for Mailer service.
// All implementations should embed UnimplementedMailerServer
// for forward compatibility
type MailerServer interface {
	SendHTML(context.Context, *SendHTMLRequest) (*SendResponse, error)
	SendTemplate(context.Context, *SendTemplateRequest) (*SendResponse, error)
}

// UnimplementedMailerServer should be embedded to have forward compatible implementations.
type UnimplementedMailerServer struct {
}

func (UnimplementedMailerServer) SendHTML(context.Context, *SendHTMLRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendHTML not implemented")
}
func (UnimplementedMailerServer) SendTemplate(context.Context, *SendTemplateRequest) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTemplate not implemented")
}

// UnsafeMailerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MailerServer will
// result in compilation errors.
type UnsafeMailerServer interface {
	mustEmbedUnimplementedMailerServer()
}

func RegisterMailerServer(s grpc.ServiceRegistrar, srv MailerServer) {
	s.RegisterService(&Mailer_ServiceDesc, srv)
}

func _Mailer_SendHTML_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendHTMLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServer).SendHTML(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kannon.Mailer/SendHTML",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServer).SendHTML(ctx, req.(*SendHTMLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mailer_SendTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MailerServer).SendTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kannon.Mailer/SendTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MailerServer).SendTemplate(ctx, req.(*SendTemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mailer_ServiceDesc is the grpc.ServiceDesc for Mailer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mailer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kannon.Mailer",
	HandlerType: (*MailerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendHTML",
			Handler:    _Mailer_SendHTML_Handler,
		},
		{
			MethodName: "SendTemplate",
			Handler:    _Mailer_SendTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mailer.proto",
}
