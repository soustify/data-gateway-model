package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	Customers pb.CustomerServiceClient = &lambdaCustomers{}
)

type lambdaCustomers struct{}

func (c *lambdaCustomers) Create(ctx context.Context, in *pb.Customer, opts ... grpc.CallOption) (*pb.CustomerResponse, error){
	var result *pb.CustomerResponse
	err := invokeAndUnmarshal(ctx, "customers_service_server_create", in, &result)
	return result, err
}

func (c *lambdaCustomers) FindAll(ctx context.Context, in *pb.CustomerListRequest, opts ... grpc.CallOption) (*pb.CustomerListResponse, error){
	var result *pb.CustomerListResponse
	err := invokeAndUnmarshal(ctx, "customers_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaCustomers) Update(ctx context.Context, in *pb.Customer, opts ... grpc.CallOption) (*pb.CustomerResponse, error){
	var result *pb.CustomerResponse
	err := invokeAndUnmarshal(ctx, "customers_service_server_update", in, &result)
	return result, err
}
