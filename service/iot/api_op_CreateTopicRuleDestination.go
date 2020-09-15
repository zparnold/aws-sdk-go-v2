// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a topic rule destination. The destination must be confirmed prior to
// use.
func (c *Client) CreateTopicRuleDestination(ctx context.Context, params *CreateTopicRuleDestinationInput, optFns ...func(*Options)) (*CreateTopicRuleDestinationOutput, error) {
	stack := middleware.NewStack("CreateTopicRuleDestination", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateTopicRuleDestinationMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateTopicRuleDestinationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTopicRuleDestination(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "CreateTopicRuleDestination",
			Err:           err,
		}
	}
	out := result.(*CreateTopicRuleDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTopicRuleDestinationInput struct {
	// The topic rule destination configuration.
	DestinationConfiguration *types.TopicRuleDestinationConfiguration
}

type CreateTopicRuleDestinationOutput struct {
	// The topic rule destination.
	TopicRuleDestination *types.TopicRuleDestination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateTopicRuleDestinationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateTopicRuleDestination{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTopicRuleDestination{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTopicRuleDestination(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateTopicRuleDestination",
	}
}
