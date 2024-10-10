package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	Services pb.ServiceServiceClient = &lambdaServices{}
)

type lambdaServices struct{}

func (c *lambdaServices) Create(ctx context.Context, in *pb.Service, opts ... grpc.CallOption) (*pb.ServiceResponse, error){
	var result *pb.ServiceResponse
	err := invokeAndUnmarshal(ctx, "services_service_server_create", in, &result)
	return result, err
}

func (c *lambdaServices) FindAll(ctx context.Context, in *pb.ServiceListRequest, opts ... grpc.CallOption) (*pb.ServiceListResponse, error){
	var result *pb.ServiceListResponse
	err := invokeAndUnmarshal(ctx, "services_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaServices) Update(ctx context.Context, in *pb.Service, opts ... grpc.CallOption) (*pb.ServiceResponse, error){
	var result *pb.ServiceResponse
	err := invokeAndUnmarshal(ctx, "services_service_server_update", in, &result)
	return result, err
}
