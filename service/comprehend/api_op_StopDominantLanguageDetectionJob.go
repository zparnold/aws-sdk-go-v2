// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops a dominant language detection job in progress. If the job state is
// IN_PROGRESS the job is marked for termination and put into the STOP_REQUESTED
// state. If the job completes before it can be stopped, it is put into the
// COMPLETED state; otherwise the job is stopped and put into the STOPPED state. If
// the job is in the COMPLETED or FAILED state when you call the
// StopDominantLanguageDetectionJob operation, the operation returns a 400 Internal
// Request Exception. When a job is stopped, any documents already processed are
// written to the output location.
func (c *Client) StopDominantLanguageDetectionJob(ctx context.Context, params *StopDominantLanguageDetectionJobInput, optFns ...func(*Options)) (*StopDominantLanguageDetectionJobOutput, error) {
	stack := middleware.NewStack("StopDominantLanguageDetectionJob", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStopDominantLanguageDetectionJobMiddlewares(stack)
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
	addOpStopDominantLanguageDetectionJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopDominantLanguageDetectionJob(options.Region), middleware.Before)
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
			OperationName: "StopDominantLanguageDetectionJob",
			Err:           err,
		}
	}
	out := result.(*StopDominantLanguageDetectionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopDominantLanguageDetectionJobInput struct {

	// The identifier of the dominant language detection job to stop.
	//
	// This member is required.
	JobId *string
}

type StopDominantLanguageDetectionJobOutput struct {

	// The identifier of the dominant language detection job to stop.
	JobId *string

	// Either STOP_REQUESTED if the job is currently running, or STOPPED if the job was
	// previously stopped with the StopDominantLanguageDetectionJob operation.
	JobStatus types.JobStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStopDominantLanguageDetectionJobMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStopDominantLanguageDetectionJob{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopDominantLanguageDetectionJob{}, middleware.After)
}

func newServiceMetadataMiddleware_opStopDominantLanguageDetectionJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "StopDominantLanguageDetectionJob",
	}
}
