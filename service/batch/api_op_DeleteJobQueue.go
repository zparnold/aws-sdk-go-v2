// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified job queue. You must first disable submissions for a queue
// with the UpdateJobQueue () operation. All jobs in the queue are terminated when
// you delete a job queue. It is not necessary to disassociate compute environments
// from a queue before submitting a DeleteJobQueue request.
func (c *Client) DeleteJobQueue(ctx context.Context, params *DeleteJobQueueInput, optFns ...func(*Options)) (*DeleteJobQueueOutput, error) {
	if params == nil {
		params = &DeleteJobQueueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteJobQueue", params, optFns, addOperationDeleteJobQueueMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteJobQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteJobQueueInput struct {

	// The short name or full Amazon Resource Name (ARN) of the queue to delete.
	//
	// This member is required.
	JobQueue *string
}

type DeleteJobQueueOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteJobQueueMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteJobQueue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteJobQueue{}, middleware.After)
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
	addOpDeleteJobQueueValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteJobQueue(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteJobQueue(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "DeleteJobQueue",
	}
}
