// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Clients that perform this test extract all headers from the response.
func (c *Client) HttpPrefixHeadersResponse(ctx context.Context, params *HttpPrefixHeadersResponseInput, optFns ...func(*Options)) (*HttpPrefixHeadersResponseOutput, error) {
	stack := middleware.NewStack("HttpPrefixHeadersResponse", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpHttpPrefixHeadersResponseMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opHttpPrefixHeadersResponse(options.Region), middleware.Before)
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
			OperationName: "HttpPrefixHeadersResponse",
			Err:           err,
		}
	}
	out := result.(*HttpPrefixHeadersResponseOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpPrefixHeadersResponseInput struct {
}

type HttpPrefixHeadersResponseOutput struct {
	PrefixHeaders map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpHttpPrefixHeadersResponseMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpHttpPrefixHeadersResponse{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpHttpPrefixHeadersResponse{}, middleware.After)
}

func newServiceMetadataMiddleware_opHttpPrefixHeadersResponse(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpPrefixHeadersResponse",
	}
}
