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

func (c *documentTypeServiceClient) invoke(ctx context.Context, operation string, body interface{}, result interface{}) error {
	return invokeAndUnmarshal(ctx, fmt.Sprintf("%s_%s", transactionsLambdaName, operation), body, result)
}

func (c *documentTypeServiceClient) Create(ctx context.Context, body *pb.DocumentTypeRequest, opts ...grpc.CallOption) (*pb.DocumentTypeResponse, error) {
	var result *pb.DocumentTypeResponse
	err := c.invoke(ctx, "create", body, &result)
	return result, err
}

func (c *documentTypeServiceClient) Update(ctx context.Context, body *pb.DocumentTypeRequest, opts ...grpc.CallOption) (*pb.DocumentTypeResponse, error) {
	var result *pb.DocumentTypeResponse
	err := c.invoke(ctx, "update", body, &result)
	return result, err
}

func (c *documentTypeServiceClient) FindAll(ctx context.Context, body *pb.DocumentTypeRequest, opts ...grpc.CallOption) (*pb.DocumentTypeListResponse, error) {
	var result *pb.DocumentTypeListResponse
	err := c.invoke(ctx, "find_all", body, &result)
	return result, err
}

func (c *documentTypeServiceClient) FindOne(ctx context.Context, body *pb.DocumentTypeRequest, opts ...grpc.CallOption) (*pb.DocumentTypeResponse, error) {
	var result *pb.DocumentTypeResponse
	err := c.invoke(ctx, "find_one", body, &result)
	return result, err
}

func (c *documentTypeServiceClient) Enabled(ctx context.Context, body *pb.Entity, opts ...grpc.CallOption) (*pb.DocumentTypeResponse, error) {
	var result *pb.DocumentTypeResponse
	err := c.invoke(ctx, "enabled", body, &result)
	return result, err
}

func (c *documentTypeServiceClient) Disable(ctx context.Context, body *pb.Entity, opts ...grpc.CallOption) (*pb.DocumentTypeResponse, error) {
	var result *pb.DocumentTypeResponse
	err := c.invoke(ctx, "disable", body, &result)
	return result, err
}
