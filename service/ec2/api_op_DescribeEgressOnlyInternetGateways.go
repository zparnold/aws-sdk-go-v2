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

// Describes one or more of your egress-only internet gateways.
func (c *Client) DescribeEgressOnlyInternetGateways(ctx context.Context, params *DescribeEgressOnlyInternetGatewaysInput, optFns ...func(*Options)) (*DescribeEgressOnlyInternetGatewaysOutput, error) {
	if params == nil {
		params = &DescribeEgressOnlyInternetGatewaysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeEgressOnlyInternetGateways", params, optFns, addOperationDescribeEgressOnlyInternetGatewaysMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeEgressOnlyInternetGatewaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeEgressOnlyInternetGatewaysInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more egress-only internet gateway IDs.
	EgressOnlyInternetGatewayIds []*string

	// One or more filters.
	//
	//     * tag: - The key/value combination of a tag assigned
	// to the resource. Use the tag key in the filter name and the tag value as the
	// filter value. For example, to find all resources that have a tag with the key
	// Owner and the value TeamA, specify tag:Owner for the filter name and TeamA for
	// the filter value.
	//
	//     * tag-key - The key of a tag assigned to the resource.
	// Use this filter to find all resources assigned a tag with a specific key,
	// regardless of the tag value.
	Filters []*types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string
}

type DescribeEgressOnlyInternetGatewaysOutput struct {

	// Information about the egress-only internet gateways.
	EgressOnlyInternetGateways []*types.EgressOnlyInternetGateway

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeEgressOnlyInternetGatewaysMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeEgressOnlyInternetGateways{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeEgressOnlyInternetGateways{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeEgressOnlyInternetGateways(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeEgressOnlyInternetGateways(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeEgressOnlyInternetGateways",
	}
}
