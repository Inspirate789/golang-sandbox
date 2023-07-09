package bidirectional_streaming

import (
	"context"
	"github.com/Inspirate789/golang-sandbox/internal/adapters/grpc/api/bidirectional_streaming"
	"github.com/Inspirate789/golang-sandbox/internal/adapters/grpc/api/calculation"
	"github.com/Inspirate789/golang-sandbox/internal/adapters/grpc/interfaces"
	"github.com/Inspirate789/golang-sandbox/internal/models"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	grpcConn   *grpc.ClientConn
	grpcClient bidirectional_streaming.BidirectionalCalculatorClient
}

func NewClient() interfaces.Client {
	return &client{}
}

func (c *client) Open(target string) error {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	var err error
	c.grpcConn, err = grpc.Dial(target, opts...)
	if err != nil {
		return errors.Wrapf(err, "fail to dial with %s", target)
	}

	c.grpcClient = bidirectional_streaming.NewBidirectionalCalculatorClient(c.grpcConn)

	return nil
}

func (c *client) Calculate(input models.Calculation) (models.Calculation, error) {
	stream, err := c.grpcClient.Calculate(context.Background())
	if err != nil {
		return input, err
	}

	request := &calculation.Calculation{
		Base:   input.Base,
		Result: input.Result,
	}

	err = stream.Send(request)
	if err != nil {
		return models.Calculation{}, err
	}

	response, err := stream.Recv()
	if err != nil {
		return models.Calculation{}, err
	}

	return models.Calculation{
		Base:   response.Base,
		Result: response.Result,
	}, nil
}

func (c *client) Close() error {
	return c.grpcConn.Close()
}
