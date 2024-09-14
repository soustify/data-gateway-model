package lambda

import (
	"context"
	"github.com/soustify/data-gateway-model/pkg/pb/proto"
	"google.golang.org/grpc"
)

var UsersClient pb.UsersServiceClient = &usersServiceLambdaClient{}

// usersServiceLambdaClient implements the UsersServiceClient interface
type usersServiceLambdaClient struct{}

// Find implements UsersServiceClient.Find
func (c *usersServiceLambdaClient) Find(param0 context.Context, param1 *pb.UserRequest, param2 ...grpc.CallOption) (*pb.UsersResponse, error) {
	panic("implement me")
}

// Save implements UsersServiceClient.Save
func (c *usersServiceLambdaClient) Save(param0 context.Context, param1 *pb.UserRequest, param2 ...grpc.CallOption) (*pb.UserResponse, error) {
	panic("implement me")
}
