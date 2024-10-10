package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	ContextDomainChanges pb.ContextDomainChangeServiceClient = &lambdaContextDomainChanges{}
)

type lambdaContextDomainChanges struct{}

func (c *lambdaContextDomainChanges) Create(ctx context.Context, in *pb.ContextDomainChange, opts ... grpc.CallOption) (*pb.ContextDomainChangeResponse, error){
	var result *pb.ContextDomainChangeResponse
	err := invokeAndUnmarshal(ctx, "context_domain_changes_service_server_create", in, &result)
	return result, err
}

func (c *lambdaContextDomainChanges) FindAll(ctx context.Context, in *pb.ContextDomainChangeListRequest, opts ... grpc.CallOption) (*pb.ContextDomainChangeListResponse, error){
	var result *pb.ContextDomainChangeListResponse
	err := invokeAndUnmarshal(ctx, "context_domain_changes_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaContextDomainChanges) Update(ctx context.Context, in *pb.ContextDomainChange, opts ... grpc.CallOption) (*pb.ContextDomainChangeResponse, error){
	var result *pb.ContextDomainChangeResponse
	err := invokeAndUnmarshal(ctx, "context_domain_changes_service_server_update", in, &result)
	return result, err
}
