// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Initiates a stop request for the current test run. AWS Device Farm immediately
// stops the run on devices where tests have not started. You are not billed for
// these devices. On devices where tests have started executing, setup suite and
// teardown suite tests run to completion on those devices. You are billed for
// setup, teardown, and any tests that were in progress or already completed.
func (c *Client) StopRun(ctx context.Context, params *StopRunInput, optFns ...func(*Options)) (*StopRunOutput, error) {
	if params == nil {
		params = &StopRunInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopRun", params, optFns, addOperationStopRunMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*StopRunOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to stop a specific run.
type StopRunInput struct {

	// Represents the Amazon Resource Name (ARN) of the Device Farm run to stop.
	//
	// This member is required.
	Arn *string
}

// Represents the results of your stop run attempt.
type StopRunOutput struct {

	// The run that was stopped.
	Run *types.Run

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopRunMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopRun{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopRun{}, middleware.After)
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
	addOpStopRunValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopRun(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopRun(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "StopRun",
	}
}
