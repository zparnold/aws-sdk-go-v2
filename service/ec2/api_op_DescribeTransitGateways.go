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

// Describes one or more transit gateways. By default, all transit gateways are
// described. Alternatively, you can filter the results.
func (c *Client) DescribeTransitGateways(ctx context.Context, params *DescribeTransitGatewaysInput, optFns ...func(*Options)) (*DescribeTransitGatewaysOutput, error) {
	if params == nil {
		params = &DescribeTransitGatewaysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTransitGateways", params, optFns, addOperationDescribeTransitGatewaysMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTransitGatewaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTransitGatewaysInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters. The possible values are:
	//
	//     *
	// options.propagation-default-route-table-id - The ID of the default propagation
	// route table.
	//
	//     * options.amazon-side-asn - The private ASN for the Amazon
	// side of a BGP session.
	//
	//     * options.association-default-route-table-id - The
	// ID of the default association route table.
	//
	//     *
	// options.auto-accept-shared-attachments - Indicates whether there is automatic
	// acceptance of attachment requests (enable | disable).
	//
	//     *
	// options.default-route-table-association - Indicates whether resource attachments
	// are automatically associated with the default association route table (enable |
	// disable).
	//
	//     * options.default-route-table-propagation - Indicates whether
	// resource attachments automatically propagate routes to the default propagation
	// route table (enable | disable).
	//
	//     * options.dns-support - Indicates whether
	// DNS support is enabled (enable | disable).
	//
	//     * options.vpn-ecmp-support -
	// Indicates whether Equal Cost Multipath Protocol support is enabled (enable |
	// disable).
	//
	//     * owner-id - The ID of the AWS account that owns the transit
	// gateway.
	//
	//     * state - The state of the attachment (available | deleted |
	// deleting | failed | modifying | pendingAcceptance | pending | rollingBack |
	// rejected | rejecting).
	//
	//     * transit-gateway-id - The ID of the transit
	// gateway.
	Filters []*types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The IDs of the transit gateways.
	TransitGatewayIds []*string
}

type DescribeTransitGatewaysOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the transit gateways.
	TransitGateways []*types.TransitGateway

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTransitGatewaysMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeTransitGateways{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTransitGateways{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTransitGateways(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeTransitGateways(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeTransitGateways",
	}
}
