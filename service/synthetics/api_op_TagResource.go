// Code generated by smithy-go-codegen DO NOT EDIT.

package synthetics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Assigns one or more tags (key-value pairs) to the specified canary. Tags can
// help you organize and categorize your resources. You can also use them to scope
// user permissions, by granting a user permission to access or change only
// resources with certain tag values. Tags don't have any semantic meaning to AWS
// and are interpreted strictly as strings of characters. You can use the
// TagResource action with a canary that already has tags. If you specify a new tag
// key for the alarm, this tag is appended to the list of tags associated with the
// alarm. If you specify a tag key that is already associated with the alarm, the
// new tag value that you specify replaces the previous value for that tag. You can
// associate as many as 50 tags with a canary.
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

	// The ARN of the canary that you're adding tags to. The ARN format of a canary is
	// arn:aws:synthetics:Region:account-id:canary:canary-name .
	//
	// This member is required.
	ResourceArn *string

	// The list of key-value pairs to associate with the canary.
	//
	// This member is required.
	Tags map[string]*string
}

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
		SigningName:   "synthetics",
		OperationName: "TagResource",
	}
}
