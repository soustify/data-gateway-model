package main

import (
	"context"
	"fmt"
	"github.com/soustify/data-gateway-model/pkg/client"
	"github.com/soustify/data-gateway-model/pkg/pb"
)

func main() {
	lambda()
}

func lambda() {
	call := client.DocumentType
	data, err := call.FindAll(context.Background(), &pb.DocumentTypeListRequest{}, nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", data.Response)
}

func grpc() {
	//conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatalf("Falha ao conectar: %v", err)
	//}
	//defer conn.Close()
	//client := pb.NewTransactionsServiceClient(conn)
	//result, err := client.Find(context.TODO(), &pb.TransactionRequest{Id: 123})
	//if err != nil {
	//	panic(err)
	//}
	//log.Println(len(result.Transactions))
}
