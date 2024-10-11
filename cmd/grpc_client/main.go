package main

import (
	"context"
	"fmt"
	"github.com/soustify/data-gateway-model/pkg/client"
	"github.com/soustify/data-gateway-model/pkg/pb"
	"github.com/soustify/data-gateway-model/pkg/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

func main() {
	lambda()
	//grpcClient()
}

func lambda() {
	call := client.DocumentType
	ctx := types.CreateContext(types.MetadataContext{
		ContextId:     "1e5a7459-91c6-4c8a-bc30-6dfa597e1g2a",
		CognitoPoolId: "1e2a7459-91c6-4c8a-bc30-6dfa597e1g2a",
		CognitoUserId: "1e3a7459-91c6-4c8a-bc30-6dfa597e1g2a",
	})

	data, err := call.FindAll(ctx, &pb.DocumentTypeListRequest{
		Payload: &pb.DocumentType{
			Entity: &pb.Entity{
				Id:     "123123123",
				Active: false,
			},
			Name: "TESTE_!!ASDS",
		},
	}, nil)

	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", data.Response)
}

func grpcClient() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	//conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Não foi possível conectar: %v", err)
	}
	defer conn.Close()

	client := pb.NewDocumentTypeServiceClient(conn)

	md := metadata.Pairs(
		"context-id", "1e5a7459-91c6-4c8a-bc30-6dfa597e1c1a",
		"cognito-pool-id", "1e5a7459-91c6-4c8a-bc30-6dfa597e1c1a",
		"cognito-user-id", "1e5a7459-91c6-4c8a-bc30-6dfa597e1c1a")

	ctx, cancel := context.WithTimeout(metadata.NewOutgoingContext(context.Background(), md), 30*time.Minute)
	defer cancel()

	result, err := client.FindAll(ctx, &pb.DocumentTypeListRequest{
		Payload: &pb.DocumentType{
			Entity: &pb.Entity{Active: false},
		},
		Metadata: &pb.Metadata{
			Page: 1,
			Size: 15,
		},
	})

	if err != nil {
		log.Fatalf("Erro ao chamar Save: %v", err)
	}
	log.Printf("Resposta de Save: %v", result)
}
