// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the queued purchases for the specified Reserved Instances.
func (c *Client) DeleteQueuedReservedInstances(ctx context.Context, params *DeleteQueuedReservedInstancesInput, optFns ...func(*Options)) (*DeleteQueuedReservedInstancesOutput, error) {
	if params == nil {
		params = &DeleteQueuedReservedInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteQueuedReservedInstances", params, optFns, addOperationDeleteQueuedReservedInstancesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteQueuedReservedInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteQueuedReservedInstancesInput struct {

	// The IDs of the Reserved Instances.
	//
	// This member is required.
	ReservedInstancesIds []*string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DeleteQueuedReservedInstancesOutput struct {

	// Information about the queued purchases that could not be deleted.
	FailedQueuedPurchaseDeletions []*types.FailedQueuedPurchaseDeletion

	// Information about the queued purchases that were successfully deleted.
	SuccessfulQueuedPurchaseDeletions []*types.SuccessfulQueuedPurchaseDeletion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteQueuedReservedInstancesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeleteQueuedReservedInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteQueuedReservedInstances{}, middleware.After)
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
	addOpDeleteQueuedReservedInstancesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteQueuedReservedInstances(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteQueuedReservedInstances(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteQueuedReservedInstances",
	}
}
