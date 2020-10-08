// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Stops an execution. This API action is not supported by EXPRESS state machines.
func (c *Client) StopExecution(ctx context.Context, params *StopExecutionInput, optFns ...func(*Options)) (*StopExecutionOutput, error) {
	if params == nil {
		params = &StopExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopExecution", params, optFns, addOperationStopExecutionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*StopExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopExecutionInput struct {

	// The Amazon Resource Name (ARN) of the execution to stop.
	//
	// This member is required.
	ExecutionArn *string

	// A more detailed explanation of the cause of the failure.
	Cause *string

	// The error code of the failure.
	Error *string
}

type StopExecutionOutput struct {

	// The date the execution is stopped.
	//
	// This member is required.
	StopDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopExecutionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpStopExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpStopExecution{}, middleware.After)
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
	addOpStopExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopExecution(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "StopExecution",
	}
}
