// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes metadata tags from an Amazon DocumentDB resource.
func (c *Client) RemoveTagsFromResource(ctx context.Context, params *RemoveTagsFromResourceInput, optFns ...func(*Options)) (*RemoveTagsFromResourceOutput, error) {
	stack := middleware.NewStack("RemoveTagsFromResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpRemoveTagsFromResourceMiddlewares(stack)
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
	addOpRemoveTagsFromResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveTagsFromResource(options.Region), middleware.Before)

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
			OperationName: "RemoveTagsFromResource",
			Err:           err,
		}
	}
	out := result.(*RemoveTagsFromResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to RemoveTagsFromResource ().
type RemoveTagsFromResourceInput struct {
	// The Amazon DocumentDB resource that the tags are removed from. This value is an
	// Amazon Resource Name (ARN).
	ResourceName *string
	// The tag key (name) of the tag to be removed.
	TagKeys []*string
}

type RemoveTagsFromResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpRemoveTagsFromResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpRemoveTagsFromResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpRemoveTagsFromResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveTagsFromResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "RemoveTagsFromResource",
	}
}