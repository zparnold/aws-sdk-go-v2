// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops a processing job.
func (c *Client) StopProcessingJob(ctx context.Context, params *StopProcessingJobInput, optFns ...func(*Options)) (*StopProcessingJobOutput, error) {
	if params == nil {
		params = &StopProcessingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopProcessingJob", params, optFns, addOperationStopProcessingJobMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*StopProcessingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopProcessingJobInput struct {

	// The name of the processing job to stop.
	//
	// This member is required.
	ProcessingJobName *string
}

type StopProcessingJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopProcessingJobMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopProcessingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopProcessingJob{}, middleware.After)
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
	addOpStopProcessingJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopProcessingJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopProcessingJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "StopProcessingJob",
	}
}
