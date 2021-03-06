// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsecuretunneling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Closes a tunnel identified by the unique tunnel id. When a CloseTunnel request
// is received, we close the WebSocket connections between the client and proxy
// server so no data can be transmitted.
func (c *Client) CloseTunnel(ctx context.Context, params *CloseTunnelInput, optFns ...func(*Options)) (*CloseTunnelOutput, error) {
	stack := middleware.NewStack("CloseTunnel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCloseTunnelMiddlewares(stack)
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
	addOpCloseTunnelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCloseTunnel(options.Region), middleware.Before)
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
			OperationName: "CloseTunnel",
			Err:           err,
		}
	}
	out := result.(*CloseTunnelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CloseTunnelInput struct {

	// The ID of the tunnel to close.
	//
	// This member is required.
	TunnelId *string

	// When set to true, AWS IoT Secure Tunneling deletes the tunnel data immediately.
	Delete *bool
}

type CloseTunnelOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCloseTunnelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCloseTunnel{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCloseTunnel{}, middleware.After)
}

func newServiceMetadataMiddleware_opCloseTunnel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsecuredtunneling",
		OperationName: "CloseTunnel",
	}
}
