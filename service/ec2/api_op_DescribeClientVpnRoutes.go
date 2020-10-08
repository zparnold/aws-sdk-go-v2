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

// Describes the routes for the specified Client VPN endpoint.
func (c *Client) DescribeClientVpnRoutes(ctx context.Context, params *DescribeClientVpnRoutesInput, optFns ...func(*Options)) (*DescribeClientVpnRoutesOutput, error) {
	if params == nil {
		params = &DescribeClientVpnRoutesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClientVpnRoutes", params, optFns, addOperationDescribeClientVpnRoutesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClientVpnRoutesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeClientVpnRoutesInput struct {

	// The ID of the Client VPN endpoint.
	//
	// This member is required.
	ClientVpnEndpointId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters. Filter names and values are case-sensitive.
	//
	//     *
	// destination-cidr - The CIDR of the route destination.
	//
	//     * origin - How the
	// route was associated with the Client VPN endpoint (associate | add-route).
	//
	//
	// * target-subnet - The ID of the subnet through which traffic is routed.
	Filters []*types.Filter

	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the nextToken
	// value.
	MaxResults *int32

	// The token to retrieve the next page of results.
	NextToken *string
}

type DescribeClientVpnRoutesOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the Client VPN endpoint routes.
	Routes []*types.ClientVpnRoute

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeClientVpnRoutesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeClientVpnRoutes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeClientVpnRoutes{}, middleware.After)
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
	addOpDescribeClientVpnRoutesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClientVpnRoutes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeClientVpnRoutes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeClientVpnRoutes",
	}
}
