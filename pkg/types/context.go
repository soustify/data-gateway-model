package types

import (
	"context"
	"errors"
	"fmt"
)

const (
	ContextId     = "context-id"
	CognitoPoolId = "cognito-pool-id"
	CognitoUserId = "cognito-user-id"
)

type (
	LambdaParameter[T any] struct {
		Content  *T
		Metadata *MetadataContext
		fields   []string
	}

	MetadataContext struct {
		ContextId     string
		CognitoPoolId string
		CognitoUserId string
	}
)

func FullValidation(meta *MetadataContext) error {
	if meta == nil {
		return errors.New("metada is null")
	}
	if meta.ContextId == "" || meta.CognitoPoolId == "" || meta.CognitoUserId == "" {
		return errors.New(fmt.Sprintf("metadata fields invalid, Context: %s, PollId: %s UserId: %s", meta.ContextId, meta.CognitoPoolId, meta.CognitoUserId))
	}
	return nil
}

func OnlyContextValidation(meta *MetadataContext) error {
	if meta == nil {
		return errors.New("metada is null")
	}
	if meta.ContextId == "" {
		return errors.New(fmt.Sprintf("metadata fields invalid, Context: %s, PollId: %s UserId: %s", meta.ContextId, meta.CognitoPoolId, meta.CognitoUserId))
	}
	return nil
}

func ContextAndUserIdValidation(meta *MetadataContext) error {
	if meta == nil {
		return errors.New("metada is null")
	}
	if meta.ContextId == "" || meta.CognitoUserId == "" {
		return errors.New(fmt.Sprintf("metadata fields invalid, Context: %s, UserId: %s", meta.ContextId, meta.CognitoUserId))
	}
	return nil
}

func ContextAndPoolIdValidation(meta *MetadataContext) error {
	if meta == nil {
		return errors.New("metada is null")
	}
	if meta.ContextId == "" || meta.CognitoPoolId == "" {
		return errors.New(fmt.Sprintf("metadata fields invalid, Context: %s, PollId: %s", meta.ContextId, meta.CognitoPoolId))
	}
	return nil
}

func UserIdAndPoolIdValidation(meta *MetadataContext) error {
	if meta == nil {
		return errors.New("metada is null")
	}
	if meta.CognitoPoolId == "" || meta.CognitoUserId == "" {
		return errors.New(fmt.Sprintf("metadata fields invalid, PollId: %s UserId: %s", meta.CognitoPoolId, meta.CognitoUserId))
	}
	return nil
}

func (c *LambdaParameter[any]) GenerateContext(ctx context.Context, validations ...func(metadataContext *MetadataContext) error) (context.Context, error) {
	if validations != nil {
		for _, validation := range validations {
			if err := validation(c.Metadata); err != nil {
				return nil, err
			}
		}
	}
	ctx = context.WithValue(ctx, ContextId, c.Metadata.ContextId)
	ctx = context.WithValue(ctx, CognitoPoolId, c.Metadata.CognitoPoolId)
	ctx = context.WithValue(ctx, CognitoUserId, c.Metadata.CognitoUserId)
	return ctx, nil
}
