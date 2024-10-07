package client

import (
	"context"
	"fmt"
	"github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	DocumentType           pb.DocumentTypeServiceClient = &documentTypeServiceClient{}
	transactionsLambdaName                              = "document_type_service_server"
)

// documentTypeServiceClient implements the DocumentTypeServiceClient interface
type documentTypeServiceClient struct{}

func (c *documentTypeServiceClient) Create(ctx context.Context, body *pb.DocumentTypeRequest, opts ...grpc.CallOption) (*pb.DocumentTypeResponse, error) {
	var result *pb.DocumentTypeResponse
	err := invokeAndUnmarshal(ctx, fmt.Sprintf("%s_create", transactionsLambdaName), body, &result)
	return result, err
}

func (c *documentTypeServiceClient) Update(ctx context.Context, body *pb.DocumentTypeRequest, opts ...grpc.CallOption) (*pb.DocumentTypeResponse, error) {
	var result *pb.DocumentTypeResponse
	err := invokeAndUnmarshal(ctx, fmt.Sprintf("%s_update", transactionsLambdaName), body, &result)
	return result, err
}

func (c *documentTypeServiceClient) FindAll(ctx context.Context, body *pb.DocumentTypeRequest, opts ...grpc.CallOption) (*pb.DocumentTypeListResponse, error) {
	var result *pb.DocumentTypeListResponse
	err := invokeAndUnmarshal(ctx, fmt.Sprintf("%s_find_all", transactionsLambdaName), body, &result)
	return result, err
}

func (c *documentTypeServiceClient) FindOne(ctx context.Context, body *pb.DocumentTypeRequest, opts ...grpc.CallOption) (*pb.DocumentTypeResponse, error) {
	var result *pb.DocumentTypeResponse
	err := invokeAndUnmarshal(ctx, fmt.Sprintf("%s_find_one", transactionsLambdaName), body, &result)
	return result, err
}

func (c *documentTypeServiceClient) Enabled(ctx context.Context, body *pb.Entity, opts ...grpc.CallOption) (*pb.DocumentTypeResponse, error) {
	var result *pb.DocumentTypeResponse
	err := invokeAndUnmarshal(ctx, fmt.Sprintf("%s_enabled", transactionsLambdaName), body, &result)
	return result, err
}

func (c *documentTypeServiceClient) Disable(ctx context.Context, body *pb.Entity, opts ...grpc.CallOption) (*pb.DocumentTypeResponse, error) {
	var result *pb.DocumentTypeResponse
	err := invokeAndUnmarshal(ctx, fmt.Sprintf("%s_disable", transactionsLambdaName), body, &result)
	return result, err
}
