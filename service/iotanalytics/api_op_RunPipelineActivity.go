// Code generated by smithy-go-codegen DO NOT EDIT.

package iotanalytics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotanalytics/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Simulates the results of running a pipeline activity on a message payload.
func (c *Client) RunPipelineActivity(ctx context.Context, params *RunPipelineActivityInput, optFns ...func(*Options)) (*RunPipelineActivityOutput, error) {
	stack := middleware.NewStack("RunPipelineActivity", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpRunPipelineActivityMiddlewares(stack)
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
	addOpRunPipelineActivityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRunPipelineActivity(options.Region), middleware.Before)
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
			OperationName: "RunPipelineActivity",
			Err:           err,
		}
	}
	out := result.(*RunPipelineActivityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RunPipelineActivityInput struct {

	// The sample message payloads on which the pipeline activity is run.
	//
	// This member is required.
	Payloads [][]byte

	// The pipeline activity that is run. This must not be a 'channel' activity or a
	// 'datastore' activity because these activities are used in a pipeline only to
	// load the original message and to store the (possibly) transformed message. If a
	// 'lambda' activity is specified, only short-running Lambda functions (those with
	// a timeout of less than 30 seconds or less) can be used.
	//
	// This member is required.
	PipelineActivity *types.PipelineActivity
}

type RunPipelineActivityOutput struct {

	// In case the pipeline activity fails, the log message that is generated.
	LogResult *string

	// The enriched or transformed sample message payloads as base64-encoded strings.
	// (The results of running the pipeline activity on each input sample message
	// payload, encoded in base64.)
	Payloads [][]byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpRunPipelineActivityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpRunPipelineActivity{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpRunPipelineActivity{}, middleware.After)
}

func newServiceMetadataMiddleware_opRunPipelineActivity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotanalytics",
		OperationName: "RunPipelineActivity",
	}
}
