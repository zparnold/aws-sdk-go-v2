// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or overwrites one or more tags for the specified AppStream 2.0 resource.
// You can tag AppStream 2.0 image builders, images, fleets, and stacks. Each tag
// consists of a key and an optional value. If a resource already has a tag with
// the same key, this operation updates its value.  <p>To list the current tags for
// your resources, use <a>ListTagsForResource</a>. To disassociate tags from your
// resources, use <a>UntagResource</a>.</p> <p>For more information about tags, see
// <a
// href="https://docs.aws.amazon.com/appstream2/latest/developerguide/tagging-basic.html">Tagging
// Your Resources</a> in the <i>Amazon AppStream 2.0 Administration Guide</i>.</p>
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

type TagResourceInput struct {

	// The Amazon Resource Name (ARN) of the resource.
	//
	// This member is required.
	ResourceArn *string

	// The tags to associate. A tag is a key-value pair, and the value is optional. For
	// example, Environment=Test. If you do not specify a value, Environment=.  <p>If
	// you do not specify a value, the value is set to an empty string.</p>
	// <p>Generally allowed characters are: letters, numbers, and spaces representable
	// in UTF-8, and the following special characters: </p> <p>_ . : / = + \ - @</p>
	//
	// This member is required.
	Tags map[string]*string
}

type TagResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTagResourceMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpTagResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpTagResource{}, middleware.After)
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
		SigningName:   "appstream",
		OperationName: "TagResource",
	}
}
