// Code generated by smithy-go-codegen DO NOT EDIT.

package elastictranscoder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The CancelJob operation cancels an unfinished job. You can only cancel a job
// that has a status of Submitted. To prevent a pipeline from starting to process a
// job while you're getting the job identifier, use UpdatePipelineStatus () to
// temporarily pause the pipeline.
func (c *Client) CancelJob(ctx context.Context, params *CancelJobInput, optFns ...func(*Options)) (*CancelJobOutput, error) {
	if params == nil {
		params = &CancelJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CancelJob", params, optFns, addOperationCancelJobMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CancelJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The CancelJobRequest structure.
type CancelJobInput struct {

	// The identifier of the job that you want to cancel. To get a list of the jobs
	// (including their jobId) that have a status of Submitted, use the
	// ListJobsByStatus () API action.
	//
	// This member is required.
	Id *string
}

// The response body contains a JSON object. If the job is successfully canceled,
// the value of Success is true.
type CancelJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCancelJobMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCancelJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelJob{}, middleware.After)
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
	addOpCancelJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCancelJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elastictranscoder",
		OperationName: "CancelJob",
	}
}
