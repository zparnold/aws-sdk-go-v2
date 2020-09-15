// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds or overwrites only the specified tags for the specified Amazon EC2 resource
// or resources. When you specify an existing tag key, the value is overwritten
// with the new value. Each resource can have a maximum of 50 tags. Each tag
// consists of a key and optional value. Tag keys must be unique per resource.
// <p>For more information about tags, see <a
// href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html">Tagging
// Your Resources</a> in the <i>Amazon Elastic Compute Cloud User Guide</i>. For
// more information about creating IAM policies that control users' access to
// resources based on tags, see <a
// href="https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-supported-iam-actions-resources.html">Supported
// Resource-Level Permissions for Amazon EC2 API Actions</a> in the <i>Amazon
// Elastic Compute Cloud User Guide</i>.</p>
func (c *Client) CreateTags(ctx context.Context, params *CreateTagsInput, optFns ...func(*Options)) (*CreateTagsOutput, error) {
	stack := middleware.NewStack("CreateTags", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpCreateTagsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateTagsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTags(options.Region), middleware.Before)

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
			OperationName: "CreateTags",
			Err:           err,
		}
	}
	out := result.(*CreateTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTagsInput struct {
	// The IDs of the resources, separated by spaces. Constraints: Up to 1000 resource
	// IDs. We recommend breaking up this request into smaller batches.
	Resources []*string
	// The tags. The value parameter is required, but if you don't want the tag to have
	// a value, specify the parameter with no value, and we set the value to an empty
	// string.
	Tags []*types.Tag
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type CreateTagsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpCreateTagsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpCreateTags{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpCreateTags{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateTags(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "CreateTags",
	}
}
