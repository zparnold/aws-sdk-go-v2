// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds, edits, or deletes tags for a health check or a hosted zone. For
// information about using tags for cost allocation, see Using Cost Allocation Tags
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
// in the AWS Billing and Cost Management User Guide.
func (c *Client) ChangeTagsForResource(ctx context.Context, params *ChangeTagsForResourceInput, optFns ...func(*Options)) (*ChangeTagsForResourceOutput, error) {
	stack := middleware.NewStack("ChangeTagsForResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpChangeTagsForResourceMiddlewares(stack)
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
	addOpChangeTagsForResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opChangeTagsForResource(options.Region), middleware.Before)

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
			OperationName: "ChangeTagsForResource",
			Err:           err,
		}
	}
	out := result.(*ChangeTagsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the tags that you want to add,
// edit, or delete.
type ChangeTagsForResourceInput struct {
	// The ID of the resource for which you want to add, change, or delete tags.
	ResourceId *string
	// A complex type that contains a list of the tags that you want to delete from the
	// specified health check or hosted zone. You can specify up to 10 keys.
	RemoveTagKeys []*string
	// The type of the resource.
	//
	//     * The resource type for health checks is
	// healthcheck.
	//
	//     * The resource type for hosted zones is hostedzone.
	ResourceType types.TagResourceType
	// A complex type that contains a list of the tags that you want to add to the
	// specified health check or hosted zone and/or the tags that you want to edit
	// Value for. You can add a maximum of 10 tags to a health check or a hosted zone.
	AddTags []*types.Tag
}

// Empty response for the request.
type ChangeTagsForResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpChangeTagsForResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpChangeTagsForResource{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpChangeTagsForResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opChangeTagsForResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ChangeTagsForResource",
	}
}
