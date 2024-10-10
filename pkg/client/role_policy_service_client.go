package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	RolesPolices pb.RolePolicyServiceClient = &lambdaRolesPolices{}
)

type lambdaRolesPolices struct{}

func (c *lambdaRolesPolices) Create(ctx context.Context, in *pb.RolePolicy, opts ... grpc.CallOption) (*pb.RolePolicyResponse, error){
	var result *pb.RolePolicyResponse
	err := invokeAndUnmarshal(ctx, "roles_polices_service_server_create", in, &result)
	return result, err
}

func (c *lambdaRolesPolices) FindAll(ctx context.Context, in *pb.RolePolicyListRequest, opts ... grpc.CallOption) (*pb.RolePolicyListResponse, error){
	var result *pb.RolePolicyListResponse
	err := invokeAndUnmarshal(ctx, "roles_polices_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaRolesPolices) Update(ctx context.Context, in *pb.RolePolicy, opts ... grpc.CallOption) (*pb.RolePolicyResponse, error){
	var result *pb.RolePolicyResponse
	err := invokeAndUnmarshal(ctx, "roles_polices_service_server_update", in, &result)
	return result, err
}
