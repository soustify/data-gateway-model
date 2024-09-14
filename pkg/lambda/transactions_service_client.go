package lambda

import (
	"context"
	"github.com/soustify/data-gateway-model/pkg/pb/proto"
	"google.golang.org/grpc"
)

var TransactionsClient pb.TransactionsServiceClient = &transactionsServiceClient{}

// transactionsServiceClient implements the TransactionsServiceClient interface
type transactionsServiceClient struct{}

// Find implements TransactionsServiceClient.Find
func (c *transactionsServiceClient) Find(param0 context.Context, param1 *pb.TransactionRequest, param2 ...grpc.CallOption) (*pb.TransactionsResponse, error) {
	panic("implement me")
}

// Save implements TransactionsServiceClient.Save
func (c *transactionsServiceClient) Save(param0 context.Context, param1 *pb.TransactionRequest, param2 ...grpc.CallOption) (*pb.TransactionResponse, error) {
	panic("implement me")
}
