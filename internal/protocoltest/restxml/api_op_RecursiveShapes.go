// Code generated by smithy-go-codegen DO NOT EDIT.
package restxml

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/restxml/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Recursive shapes
func (c *Client) RecursiveShapes(ctx context.Context, params *RecursiveShapesInput, optFns ...func(*Options)) (*RecursiveShapesOutput, error) {
	stack := middleware.NewStack("RecursiveShapes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	retry.AddRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRecursiveShapes(options.Region), middleware.Before)
	addawsRestxml_serdeOpRecursiveShapesMiddlewares(stack)

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
			OperationName: "RecursiveShapes",
			Err:           err,
		}
	}
	out := result.(*RecursiveShapesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RecursiveShapesInput struct {
	Nested *types.RecursiveShapesInputOutputNested1
}

type RecursiveShapesOutput struct {
	Nested *types.RecursiveShapesInputOutputNested1

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpRecursiveShapesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpRecursiveShapes{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpRecursiveShapes{}, middleware.After)
}

func newServiceMetadataMiddleware_opRecursiveShapes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "Rest Xml Protocol",
		ServiceID:      "restxmlprotocol",
		EndpointPrefix: "restxmlprotocol",
		OperationName:  "RecursiveShapes",
	}
}
