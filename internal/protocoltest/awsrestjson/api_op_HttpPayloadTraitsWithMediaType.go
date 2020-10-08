// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This examples uses a @mediaType trait on the payload to force a custom
// content-type to be serialized.
func (c *Client) HttpPayloadTraitsWithMediaType(ctx context.Context, params *HttpPayloadTraitsWithMediaTypeInput, optFns ...func(*Options)) (*HttpPayloadTraitsWithMediaTypeOutput, error) {
	if params == nil {
		params = &HttpPayloadTraitsWithMediaTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "HttpPayloadTraitsWithMediaType", params, optFns, addOperationHttpPayloadTraitsWithMediaTypeMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*HttpPayloadTraitsWithMediaTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type HttpPayloadTraitsWithMediaTypeInput struct {
	Blob []byte

	Foo *string
}

type HttpPayloadTraitsWithMediaTypeOutput struct {
	Blob []byte

	Foo *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationHttpPayloadTraitsWithMediaTypeMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpHttpPayloadTraitsWithMediaType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpHttpPayloadTraitsWithMediaType{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opHttpPayloadTraitsWithMediaType(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opHttpPayloadTraitsWithMediaType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "HttpPayloadTraitsWithMediaType",
	}
}
