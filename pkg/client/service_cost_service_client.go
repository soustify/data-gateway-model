package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	ServicesCost pb.ServiceCostServiceClient = &lambdaServicesCost{}
)

type lambdaServicesCost struct{}

func (c *lambdaServicesCost) Create(ctx context.Context, in *pb.ServiceCost, opts ... grpc.CallOption) (*pb.ServiceCostResponse, error){
	var result *pb.ServiceCostResponse
	err := invokeAndUnmarshal(ctx, "services_cost_service_server_create", in, &result)
	return result, err
}

func (c *lambdaServicesCost) FindAll(ctx context.Context, in *pb.ServiceCostListRequest, opts ... grpc.CallOption) (*pb.ServiceCostListResponse, error){
	var result *pb.ServiceCostListResponse
	err := invokeAndUnmarshal(ctx, "services_cost_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaServicesCost) Update(ctx context.Context, in *pb.ServiceCost, opts ... grpc.CallOption) (*pb.ServiceCostResponse, error){
	var result *pb.ServiceCostResponse
	err := invokeAndUnmarshal(ctx, "services_cost_service_server_update", in, &result)
	return result, err
}
