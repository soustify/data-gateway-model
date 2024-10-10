package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	CategoriesServices pb.CategoriesServicesServiceClient = &lambdaCategoriesServices{}
)

type lambdaCategoriesServices struct{}

func (c *lambdaCategoriesServices) Create(ctx context.Context, in *pb.CategoriesServices, opts ... grpc.CallOption) (*pb.CategoriesServicesResponse, error){
	var result *pb.CategoriesServicesResponse
	err := invokeAndUnmarshal(ctx, "categories_services_service_server_create", in, &result)
	return result, err
}

func (c *lambdaCategoriesServices) FindAll(ctx context.Context, in *pb.CategoriesServicesListRequest, opts ... grpc.CallOption) (*pb.CategoriesServicesListResponse, error){
	var result *pb.CategoriesServicesListResponse
	err := invokeAndUnmarshal(ctx, "categories_services_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaCategoriesServices) Update(ctx context.Context, in *pb.CategoriesServices, opts ... grpc.CallOption) (*pb.CategoriesServicesResponse, error){
	var result *pb.CategoriesServicesResponse
	err := invokeAndUnmarshal(ctx, "categories_services_service_server_update", in, &result)
	return result, err
}
