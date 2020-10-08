// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a specified VPC link under the caller's account in a region.
func (c *Client) GetVpcLink(ctx context.Context, params *GetVpcLinkInput, optFns ...func(*Options)) (*GetVpcLinkOutput, error) {
	if params == nil {
		params = &GetVpcLinkInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetVpcLink", params, optFns, addOperationGetVpcLinkMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetVpcLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets a specified VPC link under the caller's account in a region.
type GetVpcLinkInput struct {

	// [Required] The identifier of the VpcLink (). It is used in an Integration () to
	// reference this VpcLink ().
	//
	// This member is required.
	VpcLinkId *string

	Name *string

	Template *bool

	TemplateSkipList []*string

	Title *string
}

// An API Gateway VPC link for a RestApi () to access resources in an Amazon
// Virtual Private Cloud (VPC). To enable access to a resource in an Amazon Virtual
// Private Cloud through Amazon API Gateway, you, as an API developer, create a
// VpcLink () resource targeted for one or more network load balancers of the VPC
// and then integrate an API method with a private integration that uses the
// VpcLink (). The private integration has an integration type of HTTP or
// HTTP_PROXY and has a connection type of VPC_LINK. The integration uses the
// connectionId property to identify the VpcLink () used.
type GetVpcLinkOutput struct {

	// The description of the VPC link.
	Description *string

	// The identifier of the VpcLink (). It is used in an Integration () to reference
	// this VpcLink ().
	Id *string

	// The name used to label and identify the VPC link.
	Name *string

	// The status of the VPC link. The valid values are AVAILABLE, PENDING, DELETING,
	// or FAILED. Deploying an API will wait if the status is PENDING and will fail if
	// the status is DELETING.
	Status types.VpcLinkStatus

	// A description about the VPC link status.
	StatusMessage *string

	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// The ARN of the network load balancer of the VPC targeted by the VPC link. The
	// network load balancer must be owned by the same AWS account of the API owner.
	TargetArns []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetVpcLinkMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetVpcLink{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVpcLink{}, middleware.After)
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
	addOpGetVpcLinkValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetVpcLink(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	addAcceptHeader(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetVpcLink(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetVpcLink",
	}
}
