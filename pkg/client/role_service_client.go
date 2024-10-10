package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	Roles pb.RoleServiceClient = &lambdaRoles{}
)

type lambdaRoles struct{}

func (c *lambdaRoles) Create(ctx context.Context, in *pb.Role, opts ... grpc.CallOption) (*pb.RoleResponse, error){
	var result *pb.RoleResponse
	err := invokeAndUnmarshal(ctx, "roles_service_server_create", in, &result)
	return result, err
}

func (c *lambdaRoles) FindAll(ctx context.Context, in *pb.RoleListRequest, opts ... grpc.CallOption) (*pb.RoleListResponse, error){
	var result *pb.RoleListResponse
	err := invokeAndUnmarshal(ctx, "roles_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaRoles) Update(ctx context.Context, in *pb.Role, opts ... grpc.CallOption) (*pb.RoleResponse, error){
	var result *pb.RoleResponse
	err := invokeAndUnmarshal(ctx, "roles_service_server_update", in, &result)
	return result, err
}
