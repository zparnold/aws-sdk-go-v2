// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves each Amazon Resource Name (ARN) of schemas in the development state.
func (c *Client) ListDevelopmentSchemaArns(ctx context.Context, params *ListDevelopmentSchemaArnsInput, optFns ...func(*Options)) (*ListDevelopmentSchemaArnsOutput, error) {
	if params == nil {
		params = &ListDevelopmentSchemaArnsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDevelopmentSchemaArns", params, optFns, addOperationListDevelopmentSchemaArnsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDevelopmentSchemaArnsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDevelopmentSchemaArnsInput struct {

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The pagination token.
	NextToken *string
}

type ListDevelopmentSchemaArnsOutput struct {

	// The pagination token.
	NextToken *string

	// The ARNs of retrieved development schemas.
	SchemaArns []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDevelopmentSchemaArnsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDevelopmentSchemaArns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDevelopmentSchemaArns{}, middleware.After)
	if err != nil {
		return err
	}
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDevelopmentSchemaArns(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListDevelopmentSchemaArns(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListDevelopmentSchemaArns",
	}
}
