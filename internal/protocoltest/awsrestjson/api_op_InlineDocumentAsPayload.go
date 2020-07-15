// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example serializes an inline document as the entire HTTP payload.
func (c *Client) InlineDocumentAsPayload(ctx context.Context, params *InlineDocumentAsPayloadInput, optFns ...func(*Options)) (*InlineDocumentAsPayloadOutput, error) {
	stack := middleware.NewStack("InlineDocumentAsPayload", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opInlineDocumentAsPayload(options.Region), middleware.Before)
	addawsRestjson1_serdeOpInlineDocumentAsPayloadMiddlewares(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "InlineDocumentAsPayload",
			Err:           err,
		}
	}
	out := result.(*InlineDocumentAsPayloadOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InlineDocumentAsPayloadInput struct {
	DocumentValue smithy.Document
}

type InlineDocumentAsPayloadOutput struct {
	DocumentValue smithy.Document

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpInlineDocumentAsPayloadMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpInlineDocumentAsPayload{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpInlineDocumentAsPayload{}, middleware.After)
}

func newServiceMetadataMiddleware_opInlineDocumentAsPayload(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Rest Json Protocol",
		ServiceID:      "restjsonprotocol",
		EndpointPrefix: "restjsonprotocol",
		OperationName:  "InlineDocumentAsPayload",
	}
}
