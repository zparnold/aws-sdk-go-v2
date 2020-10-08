// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc10

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The example tests how requests and responses are serialized when there's no
// request or response payload because the operation has no input or output. While
// this should be rare, code generators must support this.
func (c *Client) NoInputAndNoOutput(ctx context.Context, params *NoInputAndNoOutputInput, optFns ...func(*Options)) (*NoInputAndNoOutputOutput, error) {
	if params == nil {
		params = &NoInputAndNoOutputInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "NoInputAndNoOutput", params, optFns, addOperationNoInputAndNoOutputMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*NoInputAndNoOutputOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type NoInputAndNoOutputInput struct {
}

type NoInputAndNoOutputOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationNoInputAndNoOutputMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpNoInputAndNoOutput{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpNoInputAndNoOutput{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opNoInputAndNoOutput(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opNoInputAndNoOutput(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "NoInputAndNoOutput",
	}
}
