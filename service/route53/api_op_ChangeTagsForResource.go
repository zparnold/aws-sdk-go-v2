// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds, edits, or deletes tags for a health check or a hosted zone. For
// information about using tags for cost allocation, see Using Cost Allocation Tags
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html)
// in the AWS Billing and Cost Management User Guide.
func (c *Client) ChangeTagsForResource(ctx context.Context, params *ChangeTagsForResourceInput, optFns ...func(*Options)) (*ChangeTagsForResourceOutput, error) {
	if params == nil {
		params = &ChangeTagsForResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ChangeTagsForResource", params, optFns, addOperationChangeTagsForResourceMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ChangeTagsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about the tags that you want to add,
// edit, or delete.
type ChangeTagsForResourceInput struct {

	// The ID of the resource for which you want to add, change, or delete tags.
	//
	// This member is required.
	ResourceId *string

	// The type of the resource.
	//
	//     * The resource type for health checks is
	// healthcheck.
	//
	//     * The resource type for hosted zones is hostedzone.
	//
	// This member is required.
	ResourceType types.TagResourceType

	// A complex type that contains a list of the tags that you want to add to the
	// specified health check or hosted zone and/or the tags that you want to edit
	// Value for. You can add a maximum of 10 tags to a health check or a hosted zone.
	AddTags []*types.Tag

	// A complex type that contains a list of the tags that you want to delete from the
	// specified health check or hosted zone. You can specify up to 10 keys.
	RemoveTagKeys []*string
}

// Empty response for the request.
type ChangeTagsForResourceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationChangeTagsForResourceMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpChangeTagsForResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpChangeTagsForResource{}, middleware.After)
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
	addOpChangeTagsForResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opChangeTagsForResource(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opChangeTagsForResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "ChangeTagsForResource",
	}
}
