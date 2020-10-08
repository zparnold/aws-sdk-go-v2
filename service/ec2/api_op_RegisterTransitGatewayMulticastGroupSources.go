// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers sources (network interfaces) with the specified transit gateway
// multicast group. A multicast source is a network interface attached to a
// supported instance that sends multicast traffic. For information about supported
// instances, see Multicast Considerations
// (https://docs.aws.amazon.com/vpc/latest/tgw/transit-gateway-limits.html#multicast-limits)
// in Amazon VPC Transit Gateways. After you add the source, use
// SearchTransitGatewayMulticastGroups
// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SearchTransitGatewayMulticastGroups.html)
// to verify that the source was added to the multicast group.
func (c *Client) RegisterTransitGatewayMulticastGroupSources(ctx context.Context, params *RegisterTransitGatewayMulticastGroupSourcesInput, optFns ...func(*Options)) (*RegisterTransitGatewayMulticastGroupSourcesOutput, error) {
	if params == nil {
		params = &RegisterTransitGatewayMulticastGroupSourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterTransitGatewayMulticastGroupSources", params, optFns, addOperationRegisterTransitGatewayMulticastGroupSourcesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterTransitGatewayMulticastGroupSourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterTransitGatewayMulticastGroupSourcesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The IP address assigned to the transit gateway multicast group.
	GroupIpAddress *string

	// The group sources' network interface IDs to register with the transit gateway
	// multicast group.
	NetworkInterfaceIds []*string

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId *string
}

type RegisterTransitGatewayMulticastGroupSourcesOutput struct {

	// Information about the transit gateway multicast group sources.
	RegisteredMulticastGroupSources *types.TransitGatewayMulticastRegisteredGroupSources

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterTransitGatewayMulticastGroupSourcesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpRegisterTransitGatewayMulticastGroupSources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpRegisterTransitGatewayMulticastGroupSources{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterTransitGatewayMulticastGroupSources(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRegisterTransitGatewayMulticastGroupSources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "RegisterTransitGatewayMulticastGroupSources",
	}
}
