// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all attributes that are associated with an object.
func (c *Client) ListObjectAttributes(ctx context.Context, params *ListObjectAttributesInput, optFns ...func(*Options)) (*ListObjectAttributesOutput, error) {
	if params == nil {
		params = &ListObjectAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListObjectAttributes", params, optFns, addOperationListObjectAttributesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListObjectAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListObjectAttributesInput struct {

	// The Amazon Resource Name (ARN) that is associated with the Directory () where
	// the object resides. For more information, see arns ().
	//
	// This member is required.
	DirectoryArn *string

	// The reference that identifies the object whose attributes will be listed.
	//
	// This member is required.
	ObjectReference *types.ObjectReference

	// Represents the manner and timing in which the successful write or update of an
	// object is reflected in a subsequent read operation of that same object.
	ConsistencyLevel types.ConsistencyLevel

	// Used to filter the list of object attributes that are associated with a certain
	// facet.
	FacetFilter *types.SchemaFacet

	// The maximum number of items to be retrieved in a single call. This is an
	// approximate number.
	MaxResults *int32

	// The pagination token.
	NextToken *string
}

type ListObjectAttributesOutput struct {

	// Attributes map that is associated with the object. AttributeArn is the key, and
	// attribute value is the value.
	Attributes []*types.AttributeKeyAndValue

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListObjectAttributesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListObjectAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListObjectAttributes{}, middleware.After)
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
	addOpListObjectAttributesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListObjectAttributes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListObjectAttributes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListObjectAttributes",
	}
}
