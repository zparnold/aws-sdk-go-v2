// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a specified VPC link under the caller's account in a region.
func (c *Client) GetVpcLink(ctx context.Context, params *GetVpcLinkInput, optFns ...func(*Options)) (*GetVpcLinkOutput, error) {
	stack := middleware.NewStack("GetVpcLink", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetVpcLinkMiddlewares(stack)
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
	addOpGetVpcLinkValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetVpcLink(options.Region), middleware.Before)

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
			OperationName: "GetVpcLink",
			Err:           err,
		}
	}
	out := result.(*GetVpcLinkOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Gets a specified VPC link under the caller's account in a region.
type GetVpcLinkInput struct {
	Template *bool
	// [Required] The identifier of the VpcLink (). It is used in an Integration () to
	// reference this VpcLink ().
	VpcLinkId        *string
	Name             *string
	TemplateSkipList []*string
	Title            *string
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
	// The identifier of the VpcLink (). It is used in an Integration () to reference
	// this VpcLink ().
	Id *string
	// The name used to label and identify the VPC link.
	Name *string
	// The description of the VPC link.
	Description *string
	// A description about the VPC link status.
	StatusMessage *string
	// The ARN of the network load balancer of the VPC targeted by the VPC link. The
	// network load balancer must be owned by the same AWS account of the API owner.
	TargetArns []*string
	// The status of the VPC link. The valid values are AVAILABLE, PENDING, DELETING,
	// or FAILED. Deploying an API will wait if the status is PENDING and will fail if
	// the status is DELETING.
	Status types.VpcLinkStatus
	// The collection of tags. Each tag element is associated with a given resource.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetVpcLinkMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetVpcLink{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetVpcLink{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetVpcLink(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetVpcLink",
	}
}
