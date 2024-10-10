package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	Address pb.AddressServiceClient = &lambdaAddress{}
)

type lambdaAddress struct{}

func (c *lambdaAddress) Create(ctx context.Context, in *pb.Address, opts ... grpc.CallOption) (*pb.AddressResponse, error){
	var result *pb.AddressResponse
	err := invokeAndUnmarshal(ctx, "address_service_server_create", in, &result)
	return result, err
}

func (c *lambdaAddress) FindAll(ctx context.Context, in *pb.AddressListRequest, opts ... grpc.CallOption) (*pb.AddressListResponse, error){
	var result *pb.AddressListResponse
	err := invokeAndUnmarshal(ctx, "address_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaAddress) Update(ctx context.Context, in *pb.Address, opts ... grpc.CallOption) (*pb.AddressResponse, error){
	var result *pb.AddressResponse
	err := invokeAndUnmarshal(ctx, "address_service_server_update", in, &result)
	return result, err
}
