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

// Describes one or more transit gateway route tables. By default, all transit
// gateway route tables are described. Alternatively, you can filter the results.
func (c *Client) DescribeTransitGatewayRouteTables(ctx context.Context, params *DescribeTransitGatewayRouteTablesInput, optFns ...func(*Options)) (*DescribeTransitGatewayRouteTablesOutput, error) {
	if params == nil {
		params = &DescribeTransitGatewayRouteTablesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTransitGatewayRouteTables", params, optFns, addOperationDescribeTransitGatewayRouteTablesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTransitGatewayRouteTablesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTransitGatewayRouteTablesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters. The possible values are:
	//
	//     *
	// default-association-route-table - Indicates whether this is the default
	// association route table for the transit gateway (true | false).
	//
	//     *
	// default-propagation-route-table - Indicates whether this is the default
	// propagation route table for the transit gateway (true | false).
	//
	//     * state -
	// The state of the attachment (available | deleted | deleting | failed | modifying
	// | pendingAcceptance | pending | rollingBack | rejected | rejecting).
	//
	//     *
	// transit-gateway-id - The ID of the transit gateway.
	//
	//     *
	// transit-gateway-route-table-id - The ID of the transit gateway route table.
	Filters []*types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The IDs of the transit gateway route tables.
	TransitGatewayRouteTableIds []*string
}

type DescribeTransitGatewayRouteTablesOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the transit gateway route tables.
	TransitGatewayRouteTables []*types.TransitGatewayRouteTable

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeTransitGatewayRouteTablesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeTransitGatewayRouteTables{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTransitGatewayRouteTables{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTransitGatewayRouteTables(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeTransitGatewayRouteTables(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeTransitGatewayRouteTables",
	}
}
