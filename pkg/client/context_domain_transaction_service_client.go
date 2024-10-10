package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	ContextDomainTransaction pb.ContextDomainTransactionServiceClient = &lambdaContextDomainTransaction{}
)

type lambdaContextDomainTransaction struct{}

func (c *lambdaContextDomainTransaction) Create(ctx context.Context, in *pb.ContextDomainTransaction, opts ... grpc.CallOption) (*pb.ContextDomainTransactionResponse, error){
	var result *pb.ContextDomainTransactionResponse
	err := invokeAndUnmarshal(ctx, "context_domain_transaction_service_server_create", in, &result)
	return result, err
}

func (c *lambdaContextDomainTransaction) FindAll(ctx context.Context, in *pb.ContextDomainTransactionListRequest, opts ... grpc.CallOption) (*pb.ContextDomainTransactionListResponse, error){
	var result *pb.ContextDomainTransactionListResponse
	err := invokeAndUnmarshal(ctx, "context_domain_transaction_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaContextDomainTransaction) Update(ctx context.Context, in *pb.ContextDomainTransaction, opts ... grpc.CallOption) (*pb.ContextDomainTransactionResponse, error){
	var result *pb.ContextDomainTransactionResponse
	err := invokeAndUnmarshal(ctx, "context_domain_transaction_service_server_update", in, &result)
	return result, err
}
