package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"log"
	"os"
)

var (
	client *lambda.Client
)

func init() {
	region := os.Getenv("REGION")
	if region == "" {
		region = "us-east-1"
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("Erro ao carregar configuração da AWS: %v", err)
	}
	client = lambda.NewFromConfig(cfg)
}

func invoke(ctx context.Context, functionName string, payload interface{}) (*lambda.InvokeOutput, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	output, err := client.Invoke(ctx, &lambda.InvokeInput{
		FunctionName: aws.String(functionName),
		Payload:      jsonData,
	})

	if output.StatusCode != 200 {
		return nil, fmt.Errorf("invocação falhou com status code: %d", output.StatusCode)
	}

	return output, err

}

// Função genérica para invocar a Lambda e desserializar a resposta
func invokeAndUnmarshal(ctx context.Context, functionName string, request interface{}, response interface{}) error {
	output, err := invoke(ctx, functionName, request)
	if err != nil {
		return err
	}
	// Desserializar a resposta
	err = json.Unmarshal(output.Payload, response)
	if err != nil {
		return fmt.Errorf("erro ao desserializar a resposta da Lambda: %w", err)
	}
	return nil
}
