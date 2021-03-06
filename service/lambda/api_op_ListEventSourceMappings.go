// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists event source mappings. Specify an EventSourceArn to only show event source
// mappings for a single event source.
func (c *Client) ListEventSourceMappings(ctx context.Context, params *ListEventSourceMappingsInput, optFns ...func(*Options)) (*ListEventSourceMappingsOutput, error) {
	stack := middleware.NewStack("ListEventSourceMappings", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListEventSourceMappingsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListEventSourceMappings(options.Region), middleware.Before)
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
			OperationName: "ListEventSourceMappings",
			Err:           err,
		}
	}
	out := result.(*ListEventSourceMappingsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEventSourceMappingsInput struct {

	// The Amazon Resource Name (ARN) of the event source.
	//
	//     * Amazon Kinesis - The
	// ARN of the data stream or a stream consumer.
	//
	//     * Amazon DynamoDB Streams -
	// The ARN of the stream.
	//
	//     * Amazon Simple Queue Service - The ARN of the
	// queue.
	EventSourceArn *string

	// The name of the Lambda function. Name formats
	//
	//     * Function name -
	// MyFunction.
	//
	//     * Function ARN -
	// arn:aws:lambda:us-west-2:123456789012:function:MyFunction.
	//
	//     * Version or
	// Alias ARN - arn:aws:lambda:us-west-2:123456789012:function:MyFunction:PROD.
	//
	//
	// * Partial ARN - 123456789012:function:MyFunction.
	//
	// The length constraint applies
	// only to the full ARN. If you specify only the function name, it's limited to 64
	// characters in length.
	FunctionName *string

	// A pagination token returned by a previous call.
	Marker *string

	// The maximum number of event source mappings to return.
	MaxItems *int32
}

type ListEventSourceMappingsOutput struct {

	// A list of event source mappings.
	EventSourceMappings []*types.EventSourceMappingConfiguration

	// A pagination token that's returned when the response doesn't contain all event
	// source mappings.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListEventSourceMappingsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListEventSourceMappings{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListEventSourceMappings{}, middleware.After)
}

func newServiceMetadataMiddleware_opListEventSourceMappings(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "ListEventSourceMappings",
	}
}
