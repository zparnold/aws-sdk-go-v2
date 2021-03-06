// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a source identifier to an existing event notification subscription.
func (c *Client) AddSourceIdentifierToSubscription(ctx context.Context, params *AddSourceIdentifierToSubscriptionInput, optFns ...func(*Options)) (*AddSourceIdentifierToSubscriptionOutput, error) {
	stack := middleware.NewStack("AddSourceIdentifierToSubscription", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpAddSourceIdentifierToSubscriptionMiddlewares(stack)
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
	addOpAddSourceIdentifierToSubscriptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddSourceIdentifierToSubscription(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

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
			OperationName: "AddSourceIdentifierToSubscription",
			Err:           err,
		}
	}
	out := result.(*AddSourceIdentifierToSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddSourceIdentifierToSubscriptionInput struct {

	// The identifier of the event source to be added. Constraints:
	//
	//     * If the
	// source type is a DB instance, then a DBInstanceIdentifier must be supplied.
	//
	//
	// * If the source type is a DB security group, a DBSecurityGroupName must be
	// supplied.
	//
	//     * If the source type is a DB parameter group, a
	// DBParameterGroupName must be supplied.
	//
	//     * If the source type is a DB
	// snapshot, a DBSnapshotIdentifier must be supplied.
	//
	// This member is required.
	SourceIdentifier *string

	// The name of the event notification subscription you want to add a source
	// identifier to.
	//
	// This member is required.
	SubscriptionName *string
}

type AddSourceIdentifierToSubscriptionOutput struct {

	// Contains the results of a successful invocation of the
	// DescribeEventSubscriptions () action.
	EventSubscription *types.EventSubscription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpAddSourceIdentifierToSubscriptionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpAddSourceIdentifierToSubscription{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpAddSourceIdentifierToSubscription{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddSourceIdentifierToSubscription(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "AddSourceIdentifierToSubscription",
	}
}
