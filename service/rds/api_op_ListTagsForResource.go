// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all tags on an Amazon RDS resource. For an overview on tagging an Amazon
// RDS resource, see Tagging Amazon RDS Resources
// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.Tagging.html)
// in the Amazon RDS User Guide.
func (c *Client) ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) {
	stack := middleware.NewStack("ListTagsForResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListTagsForResourceMiddlewares(stack)
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
	addOpListTagsForResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTagsForResource(options.Region), middleware.Before)

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
			OperationName: "ListTagsForResource",
			Err:           err,
		}
	}
	out := result.(*ListTagsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ListTagsForResourceInput struct {
	// This parameter isn't currently supported.
	Filters []*types.Filter
	// The Amazon RDS resource with tags to be listed. This value is an Amazon Resource
	// Name (ARN). For information about creating an ARN, see  Constructing an ARN for
	// Amazon RDS
	// (https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_Tagging.ARN.html#USER_Tagging.ARN.Constructing)
	// in the Amazon RDS User Guide.
	ResourceName *string
}

//
type ListTagsForResourceOutput struct {
	// List of tags returned by the ListTagsForResource operation.
	TagList []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListTagsForResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListTagsForResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListTagsForResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTagsForResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ListTagsForResource",
	}
}
