package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	PartnersCategories pb.PartnerCategoryServiceClient = &lambdaPartnersCategories{}
)

type lambdaPartnersCategories struct{}

func (c *lambdaPartnersCategories) Create(ctx context.Context, in *pb.PartnerCategory, opts ... grpc.CallOption) (*pb.PartnerCategoryResponse, error){
	var result *pb.PartnerCategoryResponse
	err := invokeAndUnmarshal(ctx, "partners_categories_service_server_create", in, &result)
	return result, err
}

func (c *lambdaPartnersCategories) FindAll(ctx context.Context, in *pb.PartnerCategoryListRequest, opts ... grpc.CallOption) (*pb.PartnerCategoryListResponse, error){
	var result *pb.PartnerCategoryListResponse
	err := invokeAndUnmarshal(ctx, "partners_categories_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaPartnersCategories) Update(ctx context.Context, in *pb.PartnerCategory, opts ... grpc.CallOption) (*pb.PartnerCategoryResponse, error){
	var result *pb.PartnerCategoryResponse
	err := invokeAndUnmarshal(ctx, "partners_categories_service_server_update", in, &result)
	return result, err
}
