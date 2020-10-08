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

// Describes the ClassicLink DNS support status of one or more VPCs. If enabled,
// the DNS hostname of a linked EC2-Classic instance resolves to its private IP
// address when addressed from an instance in the VPC to which it's linked.
// Similarly, the DNS hostname of an instance in a VPC resolves to its private IP
// address when addressed from a linked EC2-Classic instance. For more information,
// see ClassicLink
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-classiclink.html) in
// the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DescribeVpcClassicLinkDnsSupport(ctx context.Context, params *DescribeVpcClassicLinkDnsSupportInput, optFns ...func(*Options)) (*DescribeVpcClassicLinkDnsSupportOutput, error) {
	if params == nil {
		params = &DescribeVpcClassicLinkDnsSupportInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVpcClassicLinkDnsSupport", params, optFns, addOperationDescribeVpcClassicLinkDnsSupportMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVpcClassicLinkDnsSupportOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeVpcClassicLinkDnsSupportInput struct {

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// One or more VPC IDs.
	VpcIds []*string
}

type DescribeVpcClassicLinkDnsSupportOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the ClassicLink DNS support status of the VPCs.
	Vpcs []*types.ClassicLinkDnsSupport

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVpcClassicLinkDnsSupportMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeVpcClassicLinkDnsSupport{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeVpcClassicLinkDnsSupport{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVpcClassicLinkDnsSupport(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeVpcClassicLinkDnsSupport(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeVpcClassicLinkDnsSupport",
	}
}
