package unary

import (
	"context"
	"github.com/Inspirate789/golang-sandbox/internal/adapters/grpc/api/calculation"
	"github.com/Inspirate789/golang-sandbox/internal/adapters/grpc/api/unary"
	"github.com/Inspirate789/golang-sandbox/internal/adapters/grpc/interfaces"
	"github.com/Inspirate789/golang-sandbox/internal/models"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	listener   net.Listener
	grpcServer *grpc.Server
	callback   models.CalculationCallback
	unary.UnimplementedUnaryCalculatorServer
}

func NewServer(address string, callback models.CalculationCallback) (interfaces.Server, error) {
	s := server{
		listener:                           nil,
		grpcServer:                         nil,
		callback:                           callback,
		UnimplementedUnaryCalculatorServer: unary.UnimplementedUnaryCalculatorServer{},
	}

	var err error
	s.listener, err = net.Listen("tcp", address)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to listen %s", address)
	}

	s.grpcServer = grpc.NewServer()
	unary.RegisterUnaryCalculatorServer(s.grpcServer, &s)

	return &s, nil
}

func (s *server) Serve() error {
	return s.grpcServer.Serve(s.listener)
}

func (s *server) Calculate(_ context.Context, request *calculation.Calculation) (*calculation.Calculation, error) {
	result, err := s.callback(models.Calculation{
		Base:   request.Base,
		Result: request.Result,
	})
	if err != nil {
		return nil, err
	}

	return &calculation.Calculation{
		Base:   result.Base,
		Result: result.Result,
	}, nil
}

func (s *server) Close() error {
	s.grpcServer.Stop()

	return nil
}
