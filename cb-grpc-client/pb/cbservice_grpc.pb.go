// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.22.5
// source: pb/cbservice.proto

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

// CbServiceClient is the client API for CbService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CbServiceClient interface {
	GetHealth(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	BookTicket(ctx context.Context, in *BookingRequest, opts ...grpc.CallOption) (*BookingResponse, error)
	GetSeatArrangenments(ctx context.Context, in *SeatArrangmentRequest, opts ...grpc.CallOption) (*SeatArrangmentResponse, error)
	RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error)
	ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*ModifySeatResponse, error)
	GetReceipt(ctx context.Context, in *ReceiptRequest, opts ...grpc.CallOption) (*BookingResponse, error)
}

type cbServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCbServiceClient(cc grpc.ClientConnInterface) CbServiceClient {
	return &cbServiceClient{cc}
}

func (c *cbServiceClient) GetHealth(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/CbService/GetHealth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cbServiceClient) BookTicket(ctx context.Context, in *BookingRequest, opts ...grpc.CallOption) (*BookingResponse, error) {
	out := new(BookingResponse)
	err := c.cc.Invoke(ctx, "/CbService/BookTicket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cbServiceClient) GetSeatArrangenments(ctx context.Context, in *SeatArrangmentRequest, opts ...grpc.CallOption) (*SeatArrangmentResponse, error) {
	out := new(SeatArrangmentResponse)
	err := c.cc.Invoke(ctx, "/CbService/GetSeatArrangenments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cbServiceClient) RemoveUser(ctx context.Context, in *RemoveUserRequest, opts ...grpc.CallOption) (*RemoveUserResponse, error) {
	out := new(RemoveUserResponse)
	err := c.cc.Invoke(ctx, "/CbService/RemoveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cbServiceClient) ModifySeat(ctx context.Context, in *ModifySeatRequest, opts ...grpc.CallOption) (*ModifySeatResponse, error) {
	out := new(ModifySeatResponse)
	err := c.cc.Invoke(ctx, "/CbService/ModifySeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cbServiceClient) GetReceipt(ctx context.Context, in *ReceiptRequest, opts ...grpc.CallOption) (*BookingResponse, error) {
	out := new(BookingResponse)
	err := c.cc.Invoke(ctx, "/CbService/GetReceipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CbServiceServer is the server API for CbService service.
// All implementations must embed UnimplementedCbServiceServer
// for forward compatibility
type CbServiceServer interface {
	GetHealth(context.Context, *HealthRequest) (*HealthResponse, error)
	BookTicket(context.Context, *BookingRequest) (*BookingResponse, error)
	GetSeatArrangenments(context.Context, *SeatArrangmentRequest) (*SeatArrangmentResponse, error)
	RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error)
	ModifySeat(context.Context, *ModifySeatRequest) (*ModifySeatResponse, error)
	GetReceipt(context.Context, *ReceiptRequest) (*BookingResponse, error)
	mustEmbedUnimplementedCbServiceServer()
}

// UnimplementedCbServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCbServiceServer struct {
}

func (UnimplementedCbServiceServer) GetHealth(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHealth not implemented")
}
func (UnimplementedCbServiceServer) BookTicket(context.Context, *BookingRequest) (*BookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BookTicket not implemented")
}
func (UnimplementedCbServiceServer) GetSeatArrangenments(context.Context, *SeatArrangmentRequest) (*SeatArrangmentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSeatArrangenments not implemented")
}
func (UnimplementedCbServiceServer) RemoveUser(context.Context, *RemoveUserRequest) (*RemoveUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUser not implemented")
}
func (UnimplementedCbServiceServer) ModifySeat(context.Context, *ModifySeatRequest) (*ModifySeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifySeat not implemented")
}
func (UnimplementedCbServiceServer) GetReceipt(context.Context, *ReceiptRequest) (*BookingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}
func (UnimplementedCbServiceServer) mustEmbedUnimplementedCbServiceServer() {}

// UnsafeCbServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CbServiceServer will
// result in compilation errors.
type UnsafeCbServiceServer interface {
	mustEmbedUnimplementedCbServiceServer()
}

func RegisterCbServiceServer(s grpc.ServiceRegistrar, srv CbServiceServer) {
	s.RegisterService(&CbService_ServiceDesc, srv)
}

func _CbService_GetHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CbServiceServer).GetHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CbService/GetHealth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CbServiceServer).GetHealth(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CbService_BookTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CbServiceServer).BookTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CbService/BookTicket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CbServiceServer).BookTicket(ctx, req.(*BookingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CbService_GetSeatArrangenments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeatArrangmentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CbServiceServer).GetSeatArrangenments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CbService/GetSeatArrangenments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CbServiceServer).GetSeatArrangenments(ctx, req.(*SeatArrangmentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CbService_RemoveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CbServiceServer).RemoveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CbService/RemoveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CbServiceServer).RemoveUser(ctx, req.(*RemoveUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CbService_ModifySeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CbServiceServer).ModifySeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CbService/ModifySeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CbServiceServer).ModifySeat(ctx, req.(*ModifySeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CbService_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CbServiceServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CbService/GetReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CbServiceServer).GetReceipt(ctx, req.(*ReceiptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CbService_ServiceDesc is the grpc.ServiceDesc for CbService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CbService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CbService",
	HandlerType: (*CbServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetHealth",
			Handler:    _CbService_GetHealth_Handler,
		},
		{
			MethodName: "BookTicket",
			Handler:    _CbService_BookTicket_Handler,
		},
		{
			MethodName: "GetSeatArrangenments",
			Handler:    _CbService_GetSeatArrangenments_Handler,
		},
		{
			MethodName: "RemoveUser",
			Handler:    _CbService_RemoveUser_Handler,
		},
		{
			MethodName: "ModifySeat",
			Handler:    _CbService_ModifySeat_Handler,
		},
		{
			MethodName: "GetReceipt",
			Handler:    _CbService_GetReceipt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/cbservice.proto",
}
