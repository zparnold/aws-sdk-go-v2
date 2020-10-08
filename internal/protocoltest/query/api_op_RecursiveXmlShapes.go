// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/query/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Recursive shapes
func (c *Client) RecursiveXmlShapes(ctx context.Context, params *RecursiveXmlShapesInput, optFns ...func(*Options)) (*RecursiveXmlShapesOutput, error) {
	if params == nil {
		params = &RecursiveXmlShapesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RecursiveXmlShapes", params, optFns, addOperationRecursiveXmlShapesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*RecursiveXmlShapesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RecursiveXmlShapesInput struct {
}

type RecursiveXmlShapesOutput struct {
	Nested *types.RecursiveXmlShapesOutputNested1

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRecursiveXmlShapesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpRecursiveXmlShapes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpRecursiveXmlShapes{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opRecursiveXmlShapes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRecursiveXmlShapes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RecursiveXmlShapes",
	}
}
