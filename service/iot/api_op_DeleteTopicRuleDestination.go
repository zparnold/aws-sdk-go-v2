// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a topic rule destination.
func (c *Client) DeleteTopicRuleDestination(ctx context.Context, params *DeleteTopicRuleDestinationInput, optFns ...func(*Options)) (*DeleteTopicRuleDestinationOutput, error) {
	if params == nil {
		params = &DeleteTopicRuleDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTopicRuleDestination", params, optFns, addOperationDeleteTopicRuleDestinationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTopicRuleDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTopicRuleDestinationInput struct {

	// The ARN of the topic rule destination to delete.
	//
	// This member is required.
	Arn *string
}

type DeleteTopicRuleDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteTopicRuleDestinationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteTopicRuleDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteTopicRuleDestination{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpDeleteTopicRuleDestinationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTopicRuleDestination(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteTopicRuleDestination(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DeleteTopicRuleDestination",
	}
}
