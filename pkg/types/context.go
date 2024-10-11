package types

import "context"

const (
	ContextId     = "context-id"
	CognitoPoolId = "cognito-pool-id"
	CognitoUserId = "cognito-user-id"
)

type (
	LambdaParameter[T any] struct {
		Content  *T
		Metadata *MetadataContext
	}

	MetadataContext struct {
		ContextId     string
		CognitoPoolId string
		CognitoUserId string
	}
)

func (c *LambdaParameter[any]) GenerateContext() context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, ContextId, c.Metadata.ContextId)
	ctx = context.WithValue(ctx, CognitoPoolId, c.Metadata.ContextId)
	ctx = context.WithValue(ctx, CognitoUserId, c.Metadata.ContextId)
	return ctx
}

func CreateContext(value MetadataContext) context.Context {
	ctx := context.Background()
	ctx = context.WithValue(ctx, ContextId, value.ContextId)
	ctx = context.WithValue(ctx, CognitoPoolId, value.CognitoPoolId)
	ctx = context.WithValue(ctx, CognitoUserId, value.CognitoUserId)
	return ctx
}
