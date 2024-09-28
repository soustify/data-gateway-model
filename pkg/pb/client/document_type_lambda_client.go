package client

import (
	"context"
	"fmt"
	pb "github.com/soustify/data-gateway-model/pkg/pb/proto"
	"google.golang.org/grpc"
)

var (
	DocumentType           pb.DocumentTypeServiceClient = &documentTypeServiceClient{}
	transactionsLambdaName                              = "document_type_service_server"
)

// documentTypeServiceClient implements the DocumentTypeServiceClient interface
type documentTypeServiceClient struct{}

// Find implements DocumentTypeServiceClient.Find
func (c *documentTypeServiceClient) Find(ctx context.Context, body *pb.DocumentTypeRequest, _ ...grpc.CallOption) (*pb.DocumentTypeListResponse, error) {
	var result *pb.DocumentTypeListResponse
	err := invokeAndUnmarshal(ctx, fmt.Sprintf("%s_find", transactionsLambdaName), body, &result)
	return result, err
}

// Save implements DocumentTypeServiceClient.Save
func (c *documentTypeServiceClient) Save(ctx context.Context, body *pb.DocumentTypeRequest, _ ...grpc.CallOption) (*pb.DocumentTypeResponse, error) {
	var result *pb.DocumentTypeResponse
	err := invokeAndUnmarshal(ctx, fmt.Sprintf("%s_save", transactionsLambdaName), body, &result)
	return result, err
}
