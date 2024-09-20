package client

import (
	"context"
	"fmt"
	"github.com/soustify/data-gateway-model/pkg/pb/proto"
	"google.golang.org/grpc"
)

var (
	TransactionsClient     pb.TransactionsServiceClient = &transactionsServiceClient{}
	transactionsLambdaName                              = "transactions_service_server"
)

// transactionsServiceClient implements the TransactionsServiceClient interface
type transactionsServiceClient struct{}

// Find implements TransactionsServiceClient.Find
func (c *transactionsServiceClient) Find(ctx context.Context, body *pb.TransactionRequest, _ ...grpc.CallOption) (*pb.TransactionsResponse, error) {
	var result *pb.TransactionsResponse
	err := invokeAndUnmarshal(ctx, fmt.Sprintf("%s_find", transactionsLambdaName), body, &result)
	return result, err
}

// Save implements TransactionsServiceClient.Save
func (c *transactionsServiceClient) Save(ctx context.Context, body *pb.TransactionRequest, _ ...grpc.CallOption) (*pb.TransactionResponse, error) {
	var result *pb.TransactionResponse
	err := invokeAndUnmarshal(ctx, fmt.Sprintf("%s_save", transactionsLambdaName), body, &result)
	return result, err
}
