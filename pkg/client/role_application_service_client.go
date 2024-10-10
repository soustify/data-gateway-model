package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	RolesApplications pb.RoleApplicationServiceClient = &lambdaRolesApplications{}
)

type lambdaRolesApplications struct{}

func (c *lambdaRolesApplications) Create(ctx context.Context, in *pb.RoleApplication, opts ... grpc.CallOption) (*pb.RoleApplicationResponse, error){
	var result *pb.RoleApplicationResponse
	err := invokeAndUnmarshal(ctx, "roles_applications_service_server_create", in, &result)
	return result, err
}

func (c *lambdaRolesApplications) FindAll(ctx context.Context, in *pb.RoleApplicationListRequest, opts ... grpc.CallOption) (*pb.RoleApplicationListResponse, error){
	var result *pb.RoleApplicationListResponse
	err := invokeAndUnmarshal(ctx, "roles_applications_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaRolesApplications) Update(ctx context.Context, in *pb.RoleApplication, opts ... grpc.CallOption) (*pb.RoleApplicationResponse, error){
	var result *pb.RoleApplicationResponse
	err := invokeAndUnmarshal(ctx, "roles_applications_service_server_update", in, &result)
	return result, err
}
