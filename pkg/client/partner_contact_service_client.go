package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	PartnersContacts pb.PartnerContactServiceClient = &lambdaPartnersContacts{}
)

type lambdaPartnersContacts struct{}

func (c *lambdaPartnersContacts) Create(ctx context.Context, in *pb.PartnerContact, opts ... grpc.CallOption) (*pb.PartnerContactResponse, error){
	var result *pb.PartnerContactResponse
	err := invokeAndUnmarshal(ctx, "partners_contacts_service_server_create", in, &result)
	return result, err
}

func (c *lambdaPartnersContacts) FindAll(ctx context.Context, in *pb.PartnerContactListRequest, opts ... grpc.CallOption) (*pb.PartnerContactListResponse, error){
	var result *pb.PartnerContactListResponse
	err := invokeAndUnmarshal(ctx, "partners_contacts_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaPartnersContacts) Update(ctx context.Context, in *pb.PartnerContact, opts ... grpc.CallOption) (*pb.PartnerContactResponse, error){
	var result *pb.PartnerContactResponse
	err := invokeAndUnmarshal(ctx, "partners_contacts_service_server_update", in, &result)
	return result, err
}
