package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	Applications pb.ApplicationServiceClient = &lambdaApplications{}
)

type lambdaApplications struct{}

func (c *lambdaApplications) Create(ctx context.Context, in *pb.Application, opts ... grpc.CallOption) (*pb.ApplicationResponse, error){
	var result *pb.ApplicationResponse
	err := invokeAndUnmarshal(ctx, "applications_service_server_create", in, &result)
	return result, err
}

func (c *lambdaApplications) FindAll(ctx context.Context, in *pb.ApplicationListRequest, opts ... grpc.CallOption) (*pb.ApplicationListResponse, error){
	var result *pb.ApplicationListResponse
	err := invokeAndUnmarshal(ctx, "applications_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaApplications) Update(ctx context.Context, in *pb.Application, opts ... grpc.CallOption) (*pb.ApplicationResponse, error){
	var result *pb.ApplicationResponse
	err := invokeAndUnmarshal(ctx, "applications_service_server_update", in, &result)
	return result, err
}
