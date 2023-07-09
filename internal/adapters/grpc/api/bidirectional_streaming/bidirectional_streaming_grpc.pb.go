// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: bidirectional_streaming.proto

package bidirectional_streaming

import (
	context "context"
	calculation "github.com/Inspirate789/golang-sandbox/internal/adapters/grpc/api/calculation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BidirectionalCalculatorClient is the client API for BidirectionalCalculator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BidirectionalCalculatorClient interface {
	Calculate(ctx context.Context, opts ...grpc.CallOption) (BidirectionalCalculator_CalculateClient, error)
}

type bidirectionalCalculatorClient struct {
	cc grpc.ClientConnInterface
}

func NewBidirectionalCalculatorClient(cc grpc.ClientConnInterface) BidirectionalCalculatorClient {
	return &bidirectionalCalculatorClient{cc}
}

func (c *bidirectionalCalculatorClient) Calculate(ctx context.Context, opts ...grpc.CallOption) (BidirectionalCalculator_CalculateClient, error) {
	stream, err := c.cc.NewStream(ctx, &BidirectionalCalculator_ServiceDesc.Streams[0], "/calculation.BidirectionalCalculator/Calculate", opts...)
	if err != nil {
		return nil, err
	}
	x := &bidirectionalCalculatorCalculateClient{stream}
	return x, nil
}

type BidirectionalCalculator_CalculateClient interface {
	Send(*calculation.Calculation) error
	Recv() (*calculation.Calculation, error)
	grpc.ClientStream
}

type bidirectionalCalculatorCalculateClient struct {
	grpc.ClientStream
}

func (x *bidirectionalCalculatorCalculateClient) Send(m *calculation.Calculation) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bidirectionalCalculatorCalculateClient) Recv() (*calculation.Calculation, error) {
	m := new(calculation.Calculation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BidirectionalCalculatorServer is the server API for BidirectionalCalculator service.
// All implementations must embed UnimplementedBidirectionalCalculatorServer
// for forward compatibility
type BidirectionalCalculatorServer interface {
	Calculate(BidirectionalCalculator_CalculateServer) error
	mustEmbedUnimplementedBidirectionalCalculatorServer()
}

// UnimplementedBidirectionalCalculatorServer must be embedded to have forward compatible implementations.
type UnimplementedBidirectionalCalculatorServer struct {
}

func (UnimplementedBidirectionalCalculatorServer) Calculate(BidirectionalCalculator_CalculateServer) error {
	return status.Errorf(codes.Unimplemented, "method Calculate not implemented")
}
func (UnimplementedBidirectionalCalculatorServer) mustEmbedUnimplementedBidirectionalCalculatorServer() {
}

// UnsafeBidirectionalCalculatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BidirectionalCalculatorServer will
// result in compilation errors.
type UnsafeBidirectionalCalculatorServer interface {
	mustEmbedUnimplementedBidirectionalCalculatorServer()
}

func RegisterBidirectionalCalculatorServer(s grpc.ServiceRegistrar, srv BidirectionalCalculatorServer) {
	s.RegisterService(&BidirectionalCalculator_ServiceDesc, srv)
}

func _BidirectionalCalculator_Calculate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BidirectionalCalculatorServer).Calculate(&bidirectionalCalculatorCalculateServer{stream})
}

type BidirectionalCalculator_CalculateServer interface {
	Send(*calculation.Calculation) error
	Recv() (*calculation.Calculation, error)
	grpc.ServerStream
}

type bidirectionalCalculatorCalculateServer struct {
	grpc.ServerStream
}

func (x *bidirectionalCalculatorCalculateServer) Send(m *calculation.Calculation) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bidirectionalCalculatorCalculateServer) Recv() (*calculation.Calculation, error) {
	m := new(calculation.Calculation)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BidirectionalCalculator_ServiceDesc is the grpc.ServiceDesc for BidirectionalCalculator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BidirectionalCalculator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculation.BidirectionalCalculator",
	HandlerType: (*BidirectionalCalculatorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Calculate",
			Handler:       _BidirectionalCalculator_Calculate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "bidirectional_streaming.proto",
}
