// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a new Facet () to an object. An object can have more than one facet applied
// on it.
func (c *Client) AddFacetToObject(ctx context.Context, params *AddFacetToObjectInput, optFns ...func(*Options)) (*AddFacetToObjectOutput, error) {
	stack := middleware.NewStack("AddFacetToObject", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpAddFacetToObjectMiddlewares(stack)
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
	addOpAddFacetToObjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAddFacetToObject(options.Region), middleware.Before)
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
			OperationName: "AddFacetToObject",
			Err:           err,
		}
	}
	out := result.(*AddFacetToObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AddFacetToObjectInput struct {

	// The Amazon Resource Name (ARN) that is associated with the Directory () where
	// the object resides. For more information, see arns ().
	//
	// This member is required.
	DirectoryArn *string

	// A reference to the object you are adding the specified facet to.
	//
	// This member is required.
	ObjectReference *types.ObjectReference

	// Identifiers for the facet that you are adding to the object. See SchemaFacet ()
	// for details.
	//
	// This member is required.
	SchemaFacet *types.SchemaFacet

	// Attributes on the facet that you are adding to the object.
	ObjectAttributeList []*types.AttributeKeyAndValue
}

type AddFacetToObjectOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpAddFacetToObjectMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpAddFacetToObject{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpAddFacetToObject{}, middleware.After)
}

func newServiceMetadataMiddleware_opAddFacetToObject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "AddFacetToObject",
	}
}
