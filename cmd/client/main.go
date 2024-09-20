package main

import (
	"context"
	"github.com/soustify/data-gateway-model/pkg/pb/client"
	pb "github.com/soustify/data-gateway-model/pkg/pb/proto"
)

func main() {

	result, err := client.TransactionsClient.Find(context.Background(), &pb.TransactionRequest{Id: 123})
	if err != nil {
		panic(err)
	}
	println(len(result.Transactions))
}
