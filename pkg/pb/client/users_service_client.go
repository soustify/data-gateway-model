package client

import (
	"context"
	"fmt"
	"github.com/soustify/data-gateway-model/pkg/pb/proto"
	"google.golang.org/grpc"
)

var (
	UsersClient     pb.UsersServiceClient = &usersServiceClient{}
	usersLambdaName                       = "users_service_server"
)

// usersServiceClient implements the UsersServiceClient interface
type usersServiceClient struct{}

// Find implements UsersServiceClient.Find
func (c *usersServiceClient) Find(ctx context.Context, body *pb.UserRequest, _ ...grpc.CallOption) (*pb.UsersResponse, error) {
	var result *pb.UsersResponse
	err := invokeAndUnmarshal(ctx, fmt.Sprintf("%s_find", usersLambdaName), body, &result)
	return result, err
}

// Save implements UsersServiceClient.Save
func (c *usersServiceClient) Save(ctx context.Context, body *pb.UserRequest, _ ...grpc.CallOption) (*pb.UserResponse, error) {
	var result *pb.UserResponse
	err := invokeAndUnmarshal(ctx, fmt.Sprintf("%s_save", usersLambdaName), body, &result)
	return result, err
}
