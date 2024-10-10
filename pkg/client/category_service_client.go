package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	Categories pb.CategoryServiceClient = &lambdaCategories{}
)

type lambdaCategories struct{}

func (c *lambdaCategories) Create(ctx context.Context, in *pb.Category, opts ... grpc.CallOption) (*pb.CategoryResponse, error){
	var result *pb.CategoryResponse
	err := invokeAndUnmarshal(ctx, "categories_service_server_create", in, &result)
	return result, err
}

func (c *lambdaCategories) FindAll(ctx context.Context, in *pb.CategoryListRequest, opts ... grpc.CallOption) (*pb.CategoryListResponse, error){
	var result *pb.CategoryListResponse
	err := invokeAndUnmarshal(ctx, "categories_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaCategories) Update(ctx context.Context, in *pb.Category, opts ... grpc.CallOption) (*pb.CategoryResponse, error){
	var result *pb.CategoryResponse
	err := invokeAndUnmarshal(ctx, "categories_service_server_update", in, &result)
	return result, err
}
