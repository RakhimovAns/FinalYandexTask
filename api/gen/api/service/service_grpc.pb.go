// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: api/service/service.proto

package desc

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

// CalculusClient is the client API for Calculus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculusClient interface {
	Calculate(ctx context.Context, in *Expression, opts ...grpc.CallOption) (*ID, error)
	Status(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Result, error)
	Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*Result, error)
	Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*Token, error)
	Logout(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type calculusClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculusClient(cc grpc.ClientConnInterface) CalculusClient {
	return &calculusClient{cc}
}

func (c *calculusClient) Calculate(ctx context.Context, in *Expression, opts ...grpc.CallOption) (*ID, error) {
	out := new(ID)
	err := c.cc.Invoke(ctx, "/desc.calculus/Calculate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculusClient) Status(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/desc.calculus/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculusClient) Register(ctx context.Context, in *User, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/desc.calculus/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculusClient) Login(ctx context.Context, in *User, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/desc.calculus/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculusClient) Logout(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/desc.calculus/Logout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculusServer is the server API for Calculus service.
// All implementations must embed UnimplementedCalculusServer
// for forward compatibility
type CalculusServer interface {
	Calculate(context.Context, *Expression) (*ID, error)
	Status(context.Context, *ID) (*Result, error)
	Register(context.Context, *User) (*Result, error)
	Login(context.Context, *User) (*Token, error)
	Logout(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedCalculusServer()
}

// UnimplementedCalculusServer must be embedded to have forward compatible implementations.
type UnimplementedCalculusServer struct {
}

func (UnimplementedCalculusServer) Calculate(context.Context, *Expression) (*ID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedCalculusServer) Status(context.Context, *ID) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedCalculusServer) Register(context.Context, *User) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedCalculusServer) Login(context.Context, *User) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedCalculusServer) Logout(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedCalculusServer) mustEmbedUnimplementedCalculusServer() {}

// UnsafeCalculusServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculusServer will
// result in compilation errors.
type UnsafeCalculusServer interface {
	mustEmbedUnimplementedCalculusServer()
}

func RegisterCalculusServer(s grpc.ServiceRegistrar, srv CalculusServer) {
	s.RegisterService(&Calculus_ServiceDesc, srv)
}

func _Calculus_Calculate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Expression)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculusServer).Calculate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desc.calculus/Calculate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculusServer).Calculate(ctx, req.(*Expression))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculus_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculusServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desc.calculus/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculusServer).Status(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculus_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculusServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desc.calculus/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculusServer).Register(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculus_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculusServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desc.calculus/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculusServer).Login(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _Calculus_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculusServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/desc.calculus/Logout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculusServer).Logout(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Calculus_ServiceDesc is the grpc.ServiceDesc for Calculus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Calculus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "desc.calculus",
	HandlerType: (*CalculusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Calculate",
			Handler:    _Calculus_Calculate_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _Calculus_Status_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _Calculus_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Calculus_Login_Handler,
		},
		{
			MethodName: "Logout",
			Handler:    _Calculus_Logout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/service/service.proto",
}
