// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/proto/v1/proto_broadage.proto

// This document describes the `BroadageService` service and all its corresponding RPCs.

package _go

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

const (
	BroadageService_GetLiveMatches_FullMethodName  = "/core.broadage.v1.BroadageService/GetLiveMatches"
	BroadageService_GetTeamPlayers_FullMethodName  = "/core.broadage.v1.BroadageService/GetTeamPlayers"
	BroadageService_GetMatchPlayers_FullMethodName = "/core.broadage.v1.BroadageService/GetMatchPlayers"
	BroadageService_GetPlayByPlay_FullMethodName   = "/core.broadage.v1.BroadageService/GetPlayByPlay"
)

// BroadageServiceClient is the client API for BroadageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BroadageServiceClient interface {
	GetLiveMatches(ctx context.Context, in *GetLiveMatchesRequest, opts ...grpc.CallOption) (*GetLiveMatchesResponse, error)
	GetTeamPlayers(ctx context.Context, in *GetTeamPlayersRequest, opts ...grpc.CallOption) (*GetTeamPlayersResponse, error)
	GetMatchPlayers(ctx context.Context, in *GetMatchPlayersRequest, opts ...grpc.CallOption) (*GetMatchPlayersResponse, error)
	GetPlayByPlay(ctx context.Context, in *GetPlayByPlayRequest, opts ...grpc.CallOption) (*GetPlayByPlayResponse, error)
}

type broadageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBroadageServiceClient(cc grpc.ClientConnInterface) BroadageServiceClient {
	return &broadageServiceClient{cc}
}

func (c *broadageServiceClient) GetLiveMatches(ctx context.Context, in *GetLiveMatchesRequest, opts ...grpc.CallOption) (*GetLiveMatchesResponse, error) {
	out := new(GetLiveMatchesResponse)
	err := c.cc.Invoke(ctx, BroadageService_GetLiveMatches_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadageServiceClient) GetTeamPlayers(ctx context.Context, in *GetTeamPlayersRequest, opts ...grpc.CallOption) (*GetTeamPlayersResponse, error) {
	out := new(GetTeamPlayersResponse)
	err := c.cc.Invoke(ctx, BroadageService_GetTeamPlayers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadageServiceClient) GetMatchPlayers(ctx context.Context, in *GetMatchPlayersRequest, opts ...grpc.CallOption) (*GetMatchPlayersResponse, error) {
	out := new(GetMatchPlayersResponse)
	err := c.cc.Invoke(ctx, BroadageService_GetMatchPlayers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadageServiceClient) GetPlayByPlay(ctx context.Context, in *GetPlayByPlayRequest, opts ...grpc.CallOption) (*GetPlayByPlayResponse, error) {
	out := new(GetPlayByPlayResponse)
	err := c.cc.Invoke(ctx, BroadageService_GetPlayByPlay_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BroadageServiceServer is the server API for BroadageService service.
// All implementations must embed UnimplementedBroadageServiceServer
// for forward compatibility
type BroadageServiceServer interface {
	GetLiveMatches(context.Context, *GetLiveMatchesRequest) (*GetLiveMatchesResponse, error)
	GetTeamPlayers(context.Context, *GetTeamPlayersRequest) (*GetTeamPlayersResponse, error)
	GetMatchPlayers(context.Context, *GetMatchPlayersRequest) (*GetMatchPlayersResponse, error)
	GetPlayByPlay(context.Context, *GetPlayByPlayRequest) (*GetPlayByPlayResponse, error)
	mustEmbedUnimplementedBroadageServiceServer()
}

// UnimplementedBroadageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBroadageServiceServer struct {
}

func (UnimplementedBroadageServiceServer) GetLiveMatches(context.Context, *GetLiveMatchesRequest) (*GetLiveMatchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLiveMatches not implemented")
}
func (UnimplementedBroadageServiceServer) GetTeamPlayers(context.Context, *GetTeamPlayersRequest) (*GetTeamPlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeamPlayers not implemented")
}
func (UnimplementedBroadageServiceServer) GetMatchPlayers(context.Context, *GetMatchPlayersRequest) (*GetMatchPlayersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMatchPlayers not implemented")
}
func (UnimplementedBroadageServiceServer) GetPlayByPlay(context.Context, *GetPlayByPlayRequest) (*GetPlayByPlayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayByPlay not implemented")
}
func (UnimplementedBroadageServiceServer) mustEmbedUnimplementedBroadageServiceServer() {}

// UnsafeBroadageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BroadageServiceServer will
// result in compilation errors.
type UnsafeBroadageServiceServer interface {
	mustEmbedUnimplementedBroadageServiceServer()
}

func RegisterBroadageServiceServer(s grpc.ServiceRegistrar, srv BroadageServiceServer) {
	s.RegisterService(&BroadageService_ServiceDesc, srv)
}

func _BroadageService_GetLiveMatches_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLiveMatchesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadageServiceServer).GetLiveMatches(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BroadageService_GetLiveMatches_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadageServiceServer).GetLiveMatches(ctx, req.(*GetLiveMatchesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BroadageService_GetTeamPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamPlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadageServiceServer).GetTeamPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BroadageService_GetTeamPlayers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadageServiceServer).GetTeamPlayers(ctx, req.(*GetTeamPlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BroadageService_GetMatchPlayers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMatchPlayersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadageServiceServer).GetMatchPlayers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BroadageService_GetMatchPlayers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadageServiceServer).GetMatchPlayers(ctx, req.(*GetMatchPlayersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BroadageService_GetPlayByPlay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayByPlayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadageServiceServer).GetPlayByPlay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BroadageService_GetPlayByPlay_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadageServiceServer).GetPlayByPlay(ctx, req.(*GetPlayByPlayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BroadageService_ServiceDesc is the grpc.ServiceDesc for BroadageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BroadageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "core.broadage.v1.BroadageService",
	HandlerType: (*BroadageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLiveMatches",
			Handler:    _BroadageService_GetLiveMatches_Handler,
		},
		{
			MethodName: "GetTeamPlayers",
			Handler:    _BroadageService_GetTeamPlayers_Handler,
		},
		{
			MethodName: "GetMatchPlayers",
			Handler:    _BroadageService_GetMatchPlayers_Handler,
		},
		{
			MethodName: "GetPlayByPlay",
			Handler:    _BroadageService_GetPlayByPlay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/v1/proto_broadage.proto",
}