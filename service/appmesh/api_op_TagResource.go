// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates the specified tags to a resource with the specified resourceArn. If
// existing tags on a resource aren't specified in the request parameters, they
// aren't changed. When a resource is deleted, the tags associated with that
// resource are also deleted.
func (c *Client) TagResource(ctx context.Context, params *TagResourceInput, optFns ...func(*Options)) (*TagResourceOutput, error) {
	if params == nil {
		params = &TagResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TagResource", params, optFns, addOperationTagResourceMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*TagResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type TagResourceInput struct {

	// The Amazon Resource Name (ARN) of the resource to add tags to.
	//
	// This member is required.
	ResourceArn *string

	// The tags to add to the resource. A tag is an array of key-value pairs. Tag keys
	// can have a maximum character length of 128 characters, and tag values can have a
	// maximum length of 256 characters.
	//
	// This member is required.
	Tags []*types.TagRef
}

//
type TagResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTagResourceMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTagResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTagResource{}, middleware.After)
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
	addOpTagResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTagResource(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opTagResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appmesh",
		OperationName: "TagResource",
	}
}
