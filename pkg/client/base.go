package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/soustify/data-gateway-model/pkg/types"
	"log"
	"os"
)

var (
	client          *lambda.Client
	accessKeyID     = os.Getenv("AWS_ACCESS_KEY")
	secretAccessKey = os.Getenv("AWS_SECRET_ACCESS")
	lambdaEndpoint  = os.Getenv("LAMBDA_ENDPOINT") // Endpoint customizado
)

func init() {
	region := os.Getenv("REGION")
	if region == "" {
		region = "us-east-1"
	}
	var cfg aws.Config
	var err error

	// Configuração base
	options := []func(*config.LoadOptions) error{
		config.WithRegion(region),
	}

	// Se as credenciais forem definidas, use-as
	if accessKeyID != "" && secretAccessKey != "" {
		options = append(options, config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, ""),
		))
	}

	// Carrega a configuração com as opções configuradas
	cfg, err = config.LoadDefaultConfig(context.TODO(), options...)
	if err != nil {
		log.Fatalf("erro ao carregar a configuração: %v", err)
	}

	if err != nil {
		log.Fatalf("Erro ao carregar configuração da AWS: %v", err)
	}

	if lambdaEndpoint != "" {
		client = lambda.NewFromConfig(cfg, func(options *lambda.Options) {
			options.BaseEndpoint = aws.String(lambdaEndpoint)
			options.EndpointResolverV2 = lambda.NewDefaultEndpointResolverV2()
		})
	} else {
		client = lambda.NewFromConfig(cfg)
	}

}

func invoke(ctx context.Context, functionName string, payload interface{}) (*lambda.InvokeOutput, error) {
	if ctx.Value(types.ContextId) == nil || ctx.Value(types.CognitoPoolId) == nil || ctx.Value(types.CognitoUserId) == nil {
		return nil, errors.New("ContextId, CognitoPollId and CognitoUserId is mandatory on context")
	}
	content := &types.LambdaParameter[any]{
		Content: &payload,
		Metadata: &types.MetadataContext{
			ContextId:     ctx.Value(types.ContextId).(string),
			CognitoPoolId: ctx.Value(types.CognitoPoolId).(string),
			CognitoUserId: ctx.Value(types.CognitoUserId).(string),
		},
	}
	jsonData, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}
	// Invocando a Lambda
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

	err = json.Unmarshal(output.Payload, response)
	if err != nil {
		return fmt.Errorf("erro ao desserializar a resposta da Lambda: %w", err)
	}
	return nil
}
