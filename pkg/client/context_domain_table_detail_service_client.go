package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	ContextDomainTablesDetails pb.ContextDomainTableDetailServiceClient = &lambdaContextDomainTablesDetails{}
)

type lambdaContextDomainTablesDetails struct{}

func (c *lambdaContextDomainTablesDetails) Create(ctx context.Context, in *pb.ContextDomainTableDetail, opts ... grpc.CallOption) (*pb.ContextDomainTableDetailResponse, error){
	var result *pb.ContextDomainTableDetailResponse
	err := invokeAndUnmarshal(ctx, "context_domain_tables_details_service_server_create", in, &result)
	return result, err
}

func (c *lambdaContextDomainTablesDetails) FindAll(ctx context.Context, in *pb.ContextDomainTableDetailListRequest, opts ... grpc.CallOption) (*pb.ContextDomainTableDetailListResponse, error){
	var result *pb.ContextDomainTableDetailListResponse
	err := invokeAndUnmarshal(ctx, "context_domain_tables_details_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaContextDomainTablesDetails) Update(ctx context.Context, in *pb.ContextDomainTableDetail, opts ... grpc.CallOption) (*pb.ContextDomainTableDetailResponse, error){
	var result *pb.ContextDomainTableDetailResponse
	err := invokeAndUnmarshal(ctx, "context_domain_tables_details_service_server_update", in, &result)
	return result, err
}
