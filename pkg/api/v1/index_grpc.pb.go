// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/proto/v1/index.proto

// This document describes the `EchoService` service and all its corresponding RPCs.

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
	BroadageService_GetLiveMatches_FullMethodName     = "/core.template.v1.BroadageService/GetLiveMatches"
	BroadageService_GetPlayersFromTeam_FullMethodName = "/core.template.v1.BroadageService/GetPlayersFromTeam"
	BroadageService_GetBoxScores_FullMethodName       = "/core.template.v1.BroadageService/GetBoxScores"
	BroadageService_GetPlayByPlay_FullMethodName      = "/core.template.v1.BroadageService/GetPlayByPlay"
)

// BroadageServiceClient is the client API for BroadageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BroadageServiceClient interface {
	GetLiveMatches(ctx context.Context, in *GetLiveMatchesRequest, opts ...grpc.CallOption) (*GetLiveMatchesResponse, error)
	GetPlayersFromTeam(ctx context.Context, in *GetPlayersFromTeamRequest, opts ...grpc.CallOption) (*GetPlayersFromTeamResponse, error)
	GetBoxScores(ctx context.Context, in *GetBoxScoresRequest, opts ...grpc.CallOption) (*GetBoxScoresResponse, error)
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

func (c *broadageServiceClient) GetPlayersFromTeam(ctx context.Context, in *GetPlayersFromTeamRequest, opts ...grpc.CallOption) (*GetPlayersFromTeamResponse, error) {
	out := new(GetPlayersFromTeamResponse)
	err := c.cc.Invoke(ctx, BroadageService_GetPlayersFromTeam_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadageServiceClient) GetBoxScores(ctx context.Context, in *GetBoxScoresRequest, opts ...grpc.CallOption) (*GetBoxScoresResponse, error) {
	out := new(GetBoxScoresResponse)
	err := c.cc.Invoke(ctx, BroadageService_GetBoxScores_FullMethodName, in, out, opts...)
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
	GetPlayersFromTeam(context.Context, *GetPlayersFromTeamRequest) (*GetPlayersFromTeamResponse, error)
	GetBoxScores(context.Context, *GetBoxScoresRequest) (*GetBoxScoresResponse, error)
	GetPlayByPlay(context.Context, *GetPlayByPlayRequest) (*GetPlayByPlayResponse, error)
	mustEmbedUnimplementedBroadageServiceServer()
}

// UnimplementedBroadageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBroadageServiceServer struct {
}

func (UnimplementedBroadageServiceServer) GetLiveMatches(context.Context, *GetLiveMatchesRequest) (*GetLiveMatchesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLiveMatches not implemented")
}
func (UnimplementedBroadageServiceServer) GetPlayersFromTeam(context.Context, *GetPlayersFromTeamRequest) (*GetPlayersFromTeamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayersFromTeam not implemented")
}
func (UnimplementedBroadageServiceServer) GetBoxScores(context.Context, *GetBoxScoresRequest) (*GetBoxScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBoxScores not implemented")
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

func _BroadageService_GetPlayersFromTeam_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlayersFromTeamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadageServiceServer).GetPlayersFromTeam(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BroadageService_GetPlayersFromTeam_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadageServiceServer).GetPlayersFromTeam(ctx, req.(*GetPlayersFromTeamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BroadageService_GetBoxScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBoxScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadageServiceServer).GetBoxScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BroadageService_GetBoxScores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadageServiceServer).GetBoxScores(ctx, req.(*GetBoxScoresRequest))
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
	ServiceName: "core.template.v1.BroadageService",
	HandlerType: (*BroadageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLiveMatches",
			Handler:    _BroadageService_GetLiveMatches_Handler,
		},
		{
			MethodName: "GetPlayersFromTeam",
			Handler:    _BroadageService_GetPlayersFromTeam_Handler,
		},
		{
			MethodName: "GetBoxScores",
			Handler:    _BroadageService_GetBoxScores_Handler,
		},
		{
			MethodName: "GetPlayByPlay",
			Handler:    _BroadageService_GetPlayByPlay_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/v1/index.proto",
}