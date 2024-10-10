package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	PartnersUsers pb.PartnerUserServiceClient = &lambdaPartnersUsers{}
)

type lambdaPartnersUsers struct{}

func (c *lambdaPartnersUsers) Create(ctx context.Context, in *pb.PartnerUser, opts ... grpc.CallOption) (*pb.PartnerUserResponse, error){
	var result *pb.PartnerUserResponse
	err := invokeAndUnmarshal(ctx, "partners_users_service_server_create", in, &result)
	return result, err
}

func (c *lambdaPartnersUsers) FindAll(ctx context.Context, in *pb.PartnerUserListRequest, opts ... grpc.CallOption) (*pb.PartnerUserListResponse, error){
	var result *pb.PartnerUserListResponse
	err := invokeAndUnmarshal(ctx, "partners_users_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaPartnersUsers) Update(ctx context.Context, in *pb.PartnerUser, opts ... grpc.CallOption) (*pb.PartnerUserResponse, error){
	var result *pb.PartnerUserResponse
	err := invokeAndUnmarshal(ctx, "partners_users_service_server_update", in, &result)
	return result, err
}
