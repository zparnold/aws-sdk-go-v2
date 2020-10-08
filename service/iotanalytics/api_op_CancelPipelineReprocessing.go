// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels the reprocessing of data through the pipeline.
func (c *Client) CancelPipelineReprocessing(ctx context.Context, params *CancelPipelineReprocessingInput, optFns ...func(*Options)) (*CancelPipelineReprocessingOutput, error) {
	if params == nil {
		params = &CancelPipelineReprocessingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelPipelineReprocessing", params, optFns, addOperationCancelPipelineReprocessingMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelPipelineReprocessingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelPipelineReprocessingInput struct {

	// The name of pipeline for which data reprocessing is canceled.
	//
	// This member is required.
	PipelineName *string

	// The ID of the reprocessing task (returned by "StartPipelineReprocessing").
	//
	// This member is required.
	ReprocessingId *string
}

type CancelPipelineReprocessingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCancelPipelineReprocessingMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelPipelineReprocessing{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelPipelineReprocessing{}, middleware.After)
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
	addOpCancelPipelineReprocessingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelPipelineReprocessing(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCancelPipelineReprocessing(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "CancelPipelineReprocessing",
	}
}
