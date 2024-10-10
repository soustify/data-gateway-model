package client

import (
	"context"
    "github.com/soustify/data-gateway-model/pkg/pb"
	"google.golang.org/grpc"
)

var (
	Contacts pb.ContactServiceClient = &lambdaContacts{}
)

type lambdaContacts struct{}

func (c *lambdaContacts) Create(ctx context.Context, in *pb.Contact, opts ... grpc.CallOption) (*pb.ContactResponse, error){
	var result *pb.ContactResponse
	err := invokeAndUnmarshal(ctx, "contacts_service_server_create", in, &result)
	return result, err
}

func (c *lambdaContacts) FindAll(ctx context.Context, in *pb.ContactListRequest, opts ... grpc.CallOption) (*pb.ContactListResponse, error){
	var result *pb.ContactListResponse
	err := invokeAndUnmarshal(ctx, "contacts_service_server_find_all", in, &result)
	return result, err
}

func (c *lambdaContacts) Update(ctx context.Context, in *pb.Contact, opts ... grpc.CallOption) (*pb.ContactResponse, error){
	var result *pb.ContactResponse
	err := invokeAndUnmarshal(ctx, "contacts_service_server_update", in, &result)
	return result, err
}
