// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconvert

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconvert/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieve the JSON for a specific queue.
func (c *Client) GetQueue(ctx context.Context, params *GetQueueInput, optFns ...func(*Options)) (*GetQueueOutput, error) {
	if params == nil {
		params = &GetQueueInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQueue", params, optFns, addOperationGetQueueMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQueueInput struct {

	// The name of the queue that you want information about.
	//
	// This member is required.
	Name *string
}

type GetQueueOutput struct {

	// You can use queues to manage the resources that are available to your AWS
	// account for running multiple transcoding jobs at the same time. If you don't
	// specify a queue, the service sends all jobs through the default queue. For more
	// information, see
	// https://docs.aws.amazon.com/mediaconvert/latest/ug/working-with-queues.html.
	Queue *types.Queue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetQueueMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetQueue{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetQueue{}, middleware.After)
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
	addOpGetQueueValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetQueue(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetQueue(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconvert",
		OperationName: "GetQueue",
	}
}
