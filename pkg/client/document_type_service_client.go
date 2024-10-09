package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	DocumentType pb.DocumentTypeServiceClient = &lambdaDocumentType{}
)

type lambdaDocumentType struct{}

func (c *lambdaDocumentType) Create(ctx context.Context, in *pb.DocumentType, opts ... grpc.CallOption) (*pb.DocumentTypeResponse, error){
	var result *pb.DocumentTypeResponse
	err := invokeAndUnmarshal(ctx, "document_type_service_server_create", in, &result)
	return result, err
}

func (c *lambdaDocumentType) FindAll(ctx context.Context, in *pb.DocumentTypeListRequest, opts ... grpc.CallOption) (*pb.DocumentTypeListResponse, error){
	var result *pb.DocumentTypeListResponse
	err := invokeAndUnmarshal(ctx, "document_type_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaDocumentType) Update(ctx context.Context, in *pb.DocumentType, opts ... grpc.CallOption) (*pb.DocumentTypeResponse, error){
	var result *pb.DocumentTypeResponse
	err := invokeAndUnmarshal(ctx, "document_type_service_server_update", in, &result)
	return result, err
}
