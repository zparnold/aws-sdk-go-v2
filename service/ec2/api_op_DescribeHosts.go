// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified Dedicated Hosts or all your Dedicated Hosts. The results
// describe only the Dedicated Hosts in the Region you're currently using. All
// listed instances consume capacity on your Dedicated Host. Dedicated Hosts that
// have recently been released are listed with the state released.
func (c *Client) DescribeHosts(ctx context.Context, params *DescribeHostsInput, optFns ...func(*Options)) (*DescribeHostsOutput, error) {
	stack := middleware.NewStack("DescribeHosts", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeHostsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeHosts(options.Region), middleware.Before)
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
			OperationName: "DescribeHosts",
			Err:           err,
		}
	}
	out := result.(*DescribeHostsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeHostsInput struct {

	// The filters.
	//
	//     * auto-placement - Whether auto-placement is enabled or
	// disabled (on | off).
	//
	//     * availability-zone - The Availability Zone of the
	// host.
	//
	//     * client-token - The idempotency token that you provided when you
	// allocated the host.
	//
	//     * host-reservation-id - The ID of the reservation
	// assigned to this host.
	//
	//     * instance-type - The instance type size that the
	// Dedicated Host is configured to support.
	//
	//     * state - The allocation state of
	// the Dedicated Host (available | under-assessment | permanent-failure | released
	// | released-permanent-failure).
	//
	//     * tag-key - The key of a tag assigned to the
	// resource. Use this filter to find all resources assigned a tag with a specific
	// key, regardless of the tag value.
	Filter []*types.Filter

	// The IDs of the Dedicated Hosts. The IDs are used for targeted instance launches.
	HostIds []*string

	// The maximum number of results to return for the request in a single page. The
	// remaining results can be seen by sending another request with the returned
	// nextToken value. This value can be between 5 and 500. If maxResults is given a
	// larger value than 500, you receive an error. You cannot specify this parameter
	// and the host IDs parameter in the same request.
	MaxResults *int32

	// The token to use to retrieve the next page of results.
	NextToken *string
}

type DescribeHostsOutput struct {

	// Information about the Dedicated Hosts.
	Hosts []*types.Host

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeHostsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeHosts{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeHosts{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeHosts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeHosts",
	}
}
