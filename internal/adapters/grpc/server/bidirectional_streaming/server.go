package bidirectional_streaming

import (
	"github.com/Inspirate789/golang-sandbox/internal/adapters/grpc/api/bidirectional_streaming"
	"github.com/Inspirate789/golang-sandbox/internal/adapters/grpc/api/calculation"
	"github.com/Inspirate789/golang-sandbox/internal/adapters/grpc/interfaces"
	"github.com/Inspirate789/golang-sandbox/internal/models"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"io"
	"net"
)

type server struct {
	listener   net.Listener
	grpcServer *grpc.Server
	callback   models.CalculationCallback
	bidirectional_streaming.UnimplementedBidirectionalCalculatorServer
}

func NewServer(address string, callback models.CalculationCallback) (interfaces.Server, error) {
	s := server{
		listener:   nil,
		grpcServer: nil,
		callback:   callback,
		UnimplementedBidirectionalCalculatorServer: bidirectional_streaming.UnimplementedBidirectionalCalculatorServer{},
	}

	var err error
	s.listener, err = net.Listen("tcp", address)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to listen %s", address)
	}

	s.grpcServer = grpc.NewServer()
	bidirectional_streaming.RegisterBidirectionalCalculatorServer(s.grpcServer, &s)

	return &s, nil
}

func (s *server) Serve() error {
	return s.grpcServer.Serve(s.listener)
}

func (s *server) Calculate(stream bidirectional_streaming.BidirectionalCalculator_CalculateServer) error {
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		result, err := s.callback(models.Calculation{
			Base:   request.Base,
			Result: request.Result,
		})
		if err != nil {
			return err
		}

		err = stream.Send(&calculation.Calculation{
			Base:   result.Base,
			Result: result.Result,
		})
		if err != nil {
			return err
		}
	}
}

func (s *server) Close() error {
	s.grpcServer.Stop()

	return nil
}
