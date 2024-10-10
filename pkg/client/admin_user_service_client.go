package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	AdminUser pb.AdminUserServiceClient = &lambdaAdminUser{}
)

type lambdaAdminUser struct{}

func (c *lambdaAdminUser) Create(ctx context.Context, in *pb.AdminUser, opts ... grpc.CallOption) (*pb.AdminUserResponse, error){
	var result *pb.AdminUserResponse
	err := invokeAndUnmarshal(ctx, "admin_user_service_server_create", in, &result)
	return result, err
}

func (c *lambdaAdminUser) FindAll(ctx context.Context, in *pb.AdminUserListRequest, opts ... grpc.CallOption) (*pb.AdminUserListResponse, error){
	var result *pb.AdminUserListResponse
	err := invokeAndUnmarshal(ctx, "admin_user_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaAdminUser) Update(ctx context.Context, in *pb.AdminUser, opts ... grpc.CallOption) (*pb.AdminUserResponse, error){
	var result *pb.AdminUserResponse
	err := invokeAndUnmarshal(ctx, "admin_user_service_server_update", in, &result)
	return result, err
}
