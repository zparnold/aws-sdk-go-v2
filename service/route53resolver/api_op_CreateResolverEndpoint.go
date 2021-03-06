// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a resolver endpoint. There are two types of resolver endpoints, inbound
// and outbound:
//
//     * An inbound resolver endpoint forwards DNS queries to the
// DNS service for a VPC from your network or another VPC.
//
//     * An outbound
// resolver endpoint forwards DNS queries from the DNS service for a VPC to your
// network or another VPC.
func (c *Client) CreateResolverEndpoint(ctx context.Context, params *CreateResolverEndpointInput, optFns ...func(*Options)) (*CreateResolverEndpointOutput, error) {
	stack := middleware.NewStack("CreateResolverEndpoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateResolverEndpointMiddlewares(stack)
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
	addOpCreateResolverEndpointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResolverEndpoint(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

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
			OperationName: "CreateResolverEndpoint",
			Err:           err,
		}
	}
	out := result.(*CreateResolverEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResolverEndpointInput struct {

	// A unique string that identifies the request and that allows failed requests to
	// be retried without the risk of executing the operation twice. CreatorRequestId
	// can be any unique string, for example, a date/time stamp.
	//
	// This member is required.
	CreatorRequestId *string

	// Specify the applicable value:
	//
	//     * INBOUND: Resolver forwards DNS queries to
	// the DNS service for a VPC from your network or another VPC
	//
	//     * OUTBOUND:
	// Resolver forwards DNS queries from the DNS service for a VPC to your network or
	// another VPC
	//
	// This member is required.
	Direction types.ResolverEndpointDirection

	// The subnets and IP addresses in your VPC that you want DNS queries to pass
	// through on the way from your VPCs to your network (for outbound endpoints) or on
	// the way from your network to your VPCs (for inbound resolver endpoints).
	//
	// This member is required.
	IpAddresses []*types.IpAddressRequest

	// The ID of one or more security groups that you want to use to control access to
	// this VPC. The security group that you specify must include one or more inbound
	// rules (for inbound resolver endpoints) or outbound rules (for outbound resolver
	// endpoints).
	//
	// This member is required.
	SecurityGroupIds []*string

	// A friendly name that lets you easily find a configuration in the Resolver
	// dashboard in the Route 53 console.
	Name *string

	// A list of the tag keys and values that you want to associate with the endpoint.
	Tags []*types.Tag
}

type CreateResolverEndpointOutput struct {

	// Information about the CreateResolverEndpoint request, including the status of
	// the request.
	ResolverEndpoint *types.ResolverEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateResolverEndpointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateResolverEndpoint{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateResolverEndpoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateResolverEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "CreateResolverEndpoint",
	}
}
