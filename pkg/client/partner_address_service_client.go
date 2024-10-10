package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	PartnerAddresses pb.PartnerAddressServiceClient = &lambdaPartnerAddresses{}
)

type lambdaPartnerAddresses struct{}

func (c *lambdaPartnerAddresses) Create(ctx context.Context, in *pb.PartnerAddress, opts ... grpc.CallOption) (*pb.PartnerAddressResponse, error){
	var result *pb.PartnerAddressResponse
	err := invokeAndUnmarshal(ctx, "partner_addresses_service_server_create", in, &result)
	return result, err
}

func (c *lambdaPartnerAddresses) FindAll(ctx context.Context, in *pb.PartnerAddressListRequest, opts ... grpc.CallOption) (*pb.PartnerAddressListResponse, error){
	var result *pb.PartnerAddressListResponse
	err := invokeAndUnmarshal(ctx, "partner_addresses_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaPartnerAddresses) Update(ctx context.Context, in *pb.PartnerAddress, opts ... grpc.CallOption) (*pb.PartnerAddressResponse, error){
	var result *pb.PartnerAddressResponse
	err := invokeAndUnmarshal(ctx, "partner_addresses_service_server_update", in, &result)
	return result, err
}
