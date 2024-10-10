package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	Polices pb.PolicyServiceClient = &lambdaPolices{}
)

type lambdaPolices struct{}

func (c *lambdaPolices) Create(ctx context.Context, in *pb.Policy, opts ... grpc.CallOption) (*pb.PolicyResponse, error){
	var result *pb.PolicyResponse
	err := invokeAndUnmarshal(ctx, "polices_service_server_create", in, &result)
	return result, err
}

func (c *lambdaPolices) FindAll(ctx context.Context, in *pb.PolicyListRequest, opts ... grpc.CallOption) (*pb.PolicyListResponse, error){
	var result *pb.PolicyListResponse
	err := invokeAndUnmarshal(ctx, "polices_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaPolices) Update(ctx context.Context, in *pb.Policy, opts ... grpc.CallOption) (*pb.PolicyResponse, error){
	var result *pb.PolicyResponse
	err := invokeAndUnmarshal(ctx, "polices_service_server_update", in, &result)
	return result, err
}
