package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	Partners pb.PartnerServiceClient = &lambdaPartners{}
)

type lambdaPartners struct{}

func (c *lambdaPartners) Create(ctx context.Context, in *pb.Partner, opts ... grpc.CallOption) (*pb.PartnerResponse, error){
	var result *pb.PartnerResponse
	err := invokeAndUnmarshal(ctx, "partners_service_server_create", in, &result)
	return result, err
}

func (c *lambdaPartners) FindAll(ctx context.Context, in *pb.PartnerListRequest, opts ... grpc.CallOption) (*pb.PartnerListResponse, error){
	var result *pb.PartnerListResponse
	err := invokeAndUnmarshal(ctx, "partners_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaPartners) Update(ctx context.Context, in *pb.Partner, opts ... grpc.CallOption) (*pb.PartnerResponse, error){
	var result *pb.PartnerResponse
	err := invokeAndUnmarshal(ctx, "partners_service_server_update", in, &result)
	return result, err
}
