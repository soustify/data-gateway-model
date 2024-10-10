package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	PartnersService pb.PartnersServicesServiceClient = &lambdaPartnersService{}
)

type lambdaPartnersService struct{}

func (c *lambdaPartnersService) Create(ctx context.Context, in *pb.PartnersServices, opts ... grpc.CallOption) (*pb.PartnersServicesResponse, error){
	var result *pb.PartnersServicesResponse
	err := invokeAndUnmarshal(ctx, "partners_service_service_server_create", in, &result)
	return result, err
}

func (c *lambdaPartnersService) FindAll(ctx context.Context, in *pb.PartnersServicesListRequest, opts ... grpc.CallOption) (*pb.PartnersServicesListResponse, error){
	var result *pb.PartnersServicesListResponse
	err := invokeAndUnmarshal(ctx, "partners_service_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaPartnersService) Update(ctx context.Context, in *pb.PartnersServices, opts ... grpc.CallOption) (*pb.PartnersServicesResponse, error){
	var result *pb.PartnersServicesResponse
	err := invokeAndUnmarshal(ctx, "partners_service_service_server_update", in, &result)
	return result, err
}
