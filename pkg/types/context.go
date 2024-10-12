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

func (c *LambdaParameter[any]) GenerateContext(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, ContextId, c.Metadata.ContextId)
	ctx = context.WithValue(ctx, CognitoPoolId, c.Metadata.CognitoPoolId)
	ctx = context.WithValue(ctx, CognitoUserId, c.Metadata.CognitoUserId)
	return ctx
}
