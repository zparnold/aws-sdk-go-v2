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

// Describes one or more of your VPN connections. For more information, see AWS
// Site-to-Site VPN (https://docs.aws.amazon.com/vpn/latest/s2svpn/VPC_VPN.html) in
// the AWS Site-to-Site VPN User Guide.
func (c *Client) DescribeVpnConnections(ctx context.Context, params *DescribeVpnConnectionsInput, optFns ...func(*Options)) (*DescribeVpnConnectionsOutput, error) {
	if params == nil {
		params = &DescribeVpnConnectionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpnConnections", params, optFns, addOperationDescribeVpnConnectionsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpnConnectionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeVpnConnections.
type DescribeVpnConnectionsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters.
	//
	//     * customer-gateway-configuration - The configuration
	// information for the customer gateway.
	//
	//     * customer-gateway-id - The ID of a
	// customer gateway associated with the VPN connection.
	//
	//     * state - The state of
	// the VPN connection (pending | available | deleting | deleted).
	//
	//     *
	// option.static-routes-only - Indicates whether the connection has static routes
	// only. Used for devices that do not support Border Gateway Protocol (BGP).
	//
	//     *
	// route.destination-cidr-block - The destination CIDR block. This corresponds to
	// the subnet used in a customer data center.
	//
	//     * bgp-asn - The BGP Autonomous
	// System Number (ASN) associated with a BGP device.
	//
	//     * tag: - The key/value
	// combination of a tag assigned to the resource. Use the tag key in the filter
	// name and the tag value as the filter value. For example, to find all resources
	// that have a tag with the key Owner and the value TeamA, specify tag:Owner for
	// the filter name and TeamA for the filter value.
	//
	//     * tag-key - The key of a
	// tag assigned to the resource. Use this filter to find all resources assigned a
	// tag with a specific key, regardless of the tag value.
	//
	//     * type - The type of
	// VPN connection. Currently the only supported type is ipsec.1.
	//
	//     *
	// vpn-connection-id - The ID of the VPN connection.
	//
	//     * vpn-gateway-id - The ID
	// of a virtual private gateway associated with the VPN connection.
	//
	//     *
	// transit-gateway-id - The ID of a transit gateway associated with the VPN
	// connection.
	Filters []*types.Filter

	// One or more VPN connection IDs. Default: Describes your VPN connections.
	VpnConnectionIds []*string
}

// Contains the output of DescribeVpnConnections.
type DescribeVpnConnectionsOutput struct {

	// Information about one or more VPN connections.
	VpnConnections []*types.VpnConnection

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVpnConnectionsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpnConnections{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpnConnections{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpnConnections(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeVpnConnections(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpnConnections",
	}
}
