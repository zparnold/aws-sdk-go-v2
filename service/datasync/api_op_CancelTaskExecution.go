// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels execution of a task. When you cancel a task execution, the transfer of
// some files is abruptly interrupted. The contents of files that are transferred
// to the destination might be incomplete or inconsistent with the source files.
// However, if you start a new task execution on the same task and you allow the
// task execution to complete, file content on the destination is complete and
// consistent. This applies to other unexpected failures that interrupt a task
// execution. In all of these cases, AWS DataSync successfully complete the
// transfer when you start the next task execution.
func (c *Client) CancelTaskExecution(ctx context.Context, params *CancelTaskExecutionInput, optFns ...func(*Options)) (*CancelTaskExecutionOutput, error) {
	if params == nil {
		params = &CancelTaskExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelTaskExecution", params, optFns, addOperationCancelTaskExecutionMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelTaskExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// CancelTaskExecutionRequest
type CancelTaskExecutionInput struct {

	// The Amazon Resource Name (ARN) of the task execution to cancel.
	//
	// This member is required.
	TaskExecutionArn *string
}

type CancelTaskExecutionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCancelTaskExecutionMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCancelTaskExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCancelTaskExecution{}, middleware.After)
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
	addOpCancelTaskExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelTaskExecution(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCancelTaskExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "CancelTaskExecution",
	}
}
