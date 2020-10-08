// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The xmlName trait on the output structure is ignored in AWS Query. The wrapping
// element is always operation name + "Response", and inside of that wrapper is
// another wrapper named operation name + "Result".
func (c *Client) IgnoresWrappingXmlName(ctx context.Context, params *IgnoresWrappingXmlNameInput, optFns ...func(*Options)) (*IgnoresWrappingXmlNameOutput, error) {
	if params == nil {
		params = &IgnoresWrappingXmlNameInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "IgnoresWrappingXmlName", params, optFns, addOperationIgnoresWrappingXmlNameMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*IgnoresWrappingXmlNameOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type IgnoresWrappingXmlNameInput struct {
}

type IgnoresWrappingXmlNameOutput struct {
	Foo *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationIgnoresWrappingXmlNameMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpIgnoresWrappingXmlName{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpIgnoresWrappingXmlName{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opIgnoresWrappingXmlName(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opIgnoresWrappingXmlName(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "IgnoresWrappingXmlName",
	}
}
