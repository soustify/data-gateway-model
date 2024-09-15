package main

import (
	"context"
	pb "github.com/soustify/data-gateway-model/pkg/pb/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	//conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Falha ao conectar: %v", err)
	}
	defer conn.Close()
	client := pb.NewUsersServiceClient(conn)
	result, err := client.Find(context.TODO(), &pb.UserRequest{Id: 123})
	if err != nil {
		panic(err)
	}
	log.Println(len(result.Users))
}
