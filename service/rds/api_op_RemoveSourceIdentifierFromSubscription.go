// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes a source identifier from an existing RDS event notification
// subscription.
func (c *Client) RemoveSourceIdentifierFromSubscription(ctx context.Context, params *RemoveSourceIdentifierFromSubscriptionInput, optFns ...func(*Options)) (*RemoveSourceIdentifierFromSubscriptionOutput, error) {
	if params == nil {
		params = &RemoveSourceIdentifierFromSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RemoveSourceIdentifierFromSubscription", params, optFns, addOperationRemoveSourceIdentifierFromSubscriptionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*RemoveSourceIdentifierFromSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type RemoveSourceIdentifierFromSubscriptionInput struct {

	// The source identifier to be removed from the subscription, such as the DB
	// instance identifier for a DB instance or the name of a security group.
	//
	// This member is required.
	SourceIdentifier *string

	// The name of the RDS event notification subscription you want to remove a source
	// identifier from.
	//
	// This member is required.
	SubscriptionName *string
}

type RemoveSourceIdentifierFromSubscriptionOutput struct {

	// Contains the results of a successful invocation of the
	// DescribeEventSubscriptions action.
	EventSubscription *types.EventSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRemoveSourceIdentifierFromSubscriptionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRemoveSourceIdentifierFromSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRemoveSourceIdentifierFromSubscription{}, middleware.After)
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
	addOpRemoveSourceIdentifierFromSubscriptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveSourceIdentifierFromSubscription(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRemoveSourceIdentifierFromSubscription(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RemoveSourceIdentifierFromSubscription",
	}
}
