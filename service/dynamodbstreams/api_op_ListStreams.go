// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodbstreams

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of stream ARNs associated with the current account and
// endpoint. If the TableName parameter is present, then ListStreams will return
// only the streams ARNs for that table. You can call ListStreams at a maximum rate
// of 5 times per second.
func (c *Client) ListStreams(ctx context.Context, params *ListStreamsInput, optFns ...func(*Options)) (*ListStreamsOutput, error) {
	stack := middleware.NewStack("ListStreams", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpListStreamsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListStreams(options.Region), middleware.Before)
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
			OperationName: "ListStreams",
			Err:           err,
		}
	}
	out := result.(*ListStreamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a ListStreams operation.
type ListStreamsInput struct {

	// The ARN (Amazon Resource Name) of the first item that this operation will
	// evaluate. Use the value that was returned for LastEvaluatedStreamArn in the
	// previous operation.
	ExclusiveStartStreamArn *string

	// The maximum number of streams to return. The upper limit is 100.
	Limit *int32

	// If this parameter is provided, then only the streams associated with this table
	// name are returned.
	TableName *string
}

// Represents the output of a ListStreams operation.
type ListStreamsOutput struct {

	// The stream ARN of the item where the operation stopped, inclusive of the
	// previous result set. Use this value to start a new operation, excluding this
	// value in the new request. If LastEvaluatedStreamArn is empty, then the "last
	// page" of results has been processed and there is no more data to be retrieved.
	// If LastEvaluatedStreamArn is not empty, it does not necessarily mean that there
	// is more data in the result set. The only way to know when you have reached the
	// end of the result set is when LastEvaluatedStreamArn is empty.
	LastEvaluatedStreamArn *string

	// A list of stream descriptors associated with the current account and endpoint.
	Streams []*types.Stream

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpListStreamsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpListStreams{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpListStreams{}, middleware.After)
}

func newServiceMetadataMiddleware_opListStreams(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "ListStreams",
	}
}
