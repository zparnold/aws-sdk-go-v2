// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/ec2query/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example serializes enums as top level properties, in lists, sets, and maps.
func (c *Client) XmlEnums(ctx context.Context, params *XmlEnumsInput, optFns ...func(*Options)) (*XmlEnumsOutput, error) {
	if params == nil {
		params = &XmlEnumsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "XmlEnums", params, optFns, addOperationXmlEnumsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*XmlEnumsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type XmlEnumsInput struct {
}

type XmlEnumsOutput struct {
	FooEnum1 types.FooEnum

	FooEnum2 types.FooEnum

	FooEnum3 types.FooEnum

	FooEnumList []types.FooEnum

	FooEnumMap map[string]types.FooEnum

	FooEnumSet []types.FooEnum

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationXmlEnumsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpXmlEnums{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpXmlEnums{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opXmlEnums(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opXmlEnums(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "XmlEnums",
	}
}
