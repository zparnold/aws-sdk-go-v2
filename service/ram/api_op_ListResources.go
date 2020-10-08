// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the resources that you added to a resource shares or the resources that
// are shared with you.
func (c *Client) ListResources(ctx context.Context, params *ListResourcesInput, optFns ...func(*Options)) (*ListResourcesOutput, error) {
	if params == nil {
		params = &ListResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListResources", params, optFns, addOperationListResourcesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListResourcesInput struct {

	// The type of owner.
	//
	// This member is required.
	ResourceOwner types.ResourceOwner

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The principal.
	Principal *string

	// The Amazon Resource Names (ARN) of the resources.
	ResourceArns []*string

	// The Amazon Resource Names (ARN) of the resource shares.
	ResourceShareArns []*string

	// The resource type. Valid values: codebuild:Project | codebuild:ReportGroup |
	// ec2:CapacityReservation | ec2:DedicatedHost | ec2:Subnet |
	// ec2:TrafficMirrorTarget | ec2:TransitGateway | imagebuilder:Component |
	// imagebuilder:Image | imagebuilder:ImageRecipe |
	// license-manager:LicenseConfiguration I resource-groups:Group | rds:Cluster |
	// route53resolver:ResolverRule
	ResourceType *string
}

type ListResourcesOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the resources.
	Resources []*types.Resource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListResourcesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListResources{}, middleware.After)
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
	addOpListResourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListResources(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "ListResources",
	}
}
