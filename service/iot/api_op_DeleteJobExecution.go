// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a job execution.
func (c *Client) DeleteJobExecution(ctx context.Context, params *DeleteJobExecutionInput, optFns ...func(*Options)) (*DeleteJobExecutionOutput, error) {
	stack := middleware.NewStack("DeleteJobExecution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteJobExecutionMiddlewares(stack)
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
	addOpDeleteJobExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteJobExecution(options.Region), middleware.Before)
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
			OperationName: "DeleteJobExecution",
			Err:           err,
		}
	}
	out := result.(*DeleteJobExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteJobExecutionInput struct {

	// The ID of the job execution to be deleted. The executionNumber refers to the
	// execution of a particular job on a particular device. Note that once a job
	// execution is deleted, the executionNumber may be reused by IoT, so be sure you
	// get and use the correct value here.
	//
	// This member is required.
	ExecutionNumber *int64

	// The ID of the job whose execution on a particular device will be deleted.
	//
	// This member is required.
	JobId *string

	// The name of the thing whose job execution will be deleted.
	//
	// This member is required.
	ThingName *string

	// (Optional) When true, you can delete a job execution which is "IN_PROGRESS".
	// Otherwise, you can only delete a job execution which is in a terminal state
	// ("SUCCEEDED", "FAILED", "REJECTED", "REMOVED" or "CANCELED") or an exception
	// will occur. The default is false. Deleting a job execution which is
	// "IN_PROGRESS", will cause the device to be unable to access job information or
	// update the job execution status. Use caution and ensure that the device is able
	// to recover to a valid state.
	Force *bool
}

type DeleteJobExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteJobExecutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteJobExecution{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteJobExecution{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteJobExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DeleteJobExecution",
	}
}
