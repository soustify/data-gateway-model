package adapter

import (
	"context"
	"github.com/soustify/data-gateway-model/pkg/pb"
)

type (
	documentType struct {
		function func(context.Context, *pb.DocumentTypeListRequest) (*pb.DocumentTypeListResponse, error)
	}

	DocumentType interface {
		DocumentTypeMessageAdapter(context.Context, DocumentTypeRequestPayload) (*pb.DocumentTypeResponse, error)
		DocumentTypeListRequestMessageAdapter(context.Context, DocumentTypeListRequestPayload) (*pb.DocumentTypeListResponse, error)
	}

	DocumentTypeListRequestPayload struct {
		Content *pb.DocumentTypeListRequest
	}

	DocumentTypeRequestPayload struct {
		Content *pb.DocumentType
	}
)

func (d documentType) DocumentTypeMessageAdapter(ctx context.Context, payload DocumentTypeRequestPayload) (*pb.DocumentTypeResponse, error) {
	panic("implement me")
}

func (d documentType) DocumentTypeListRequestMessageAdapter(ctx context.Context, payload DocumentTypeListRequestPayload) (*pb.DocumentTypeListResponse, error) {
	panic("implement me")
}

func NewDocumentTypeAdapter(function func(context.Context, *pb.DocumentTypeListRequest) (*pb.DocumentTypeListResponse, error)) DocumentType {
	return &documentType{
		function: function,
	}
}
