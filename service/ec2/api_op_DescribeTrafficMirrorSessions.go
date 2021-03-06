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

// Describes one or more Traffic Mirror sessions. By default, all Traffic Mirror
// sessions are described. Alternatively, you can filter the results.
func (c *Client) DescribeTrafficMirrorSessions(ctx context.Context, params *DescribeTrafficMirrorSessionsInput, optFns ...func(*Options)) (*DescribeTrafficMirrorSessionsOutput, error) {
	stack := middleware.NewStack("DescribeTrafficMirrorSessions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeTrafficMirrorSessionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTrafficMirrorSessions(options.Region), middleware.Before)
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
			OperationName: "DescribeTrafficMirrorSessions",
			Err:           err,
		}
	}
	out := result.(*DescribeTrafficMirrorSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeTrafficMirrorSessionsInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// One or more filters. The possible values are:
	//
	//     * description: The Traffic
	// Mirror session description.
	//
	//     * network-interface-id: The ID of the Traffic
	// Mirror session network interface.
	//
	//     * owner-id: The ID of the account that
	// owns the Traffic Mirror session.
	//
	//     * packet-length: The assigned number of
	// packets to mirror.
	//
	//     * session-number: The assigned session number.
	//
	//     *
	// traffic-mirror-filter-id: The ID of the Traffic Mirror filter.
	//
	//     *
	// traffic-mirror-session-id: The ID of the Traffic Mirror session.
	//
	//     *
	// traffic-mirror-target-id: The ID of the Traffic Mirror target.
	//
	//     *
	// virtual-network-id: The virtual network ID of the Traffic Mirror session.
	Filters []*types.Filter

	// The maximum number of results to return with a single call. To retrieve the
	// remaining results, make another call with the returned nextToken value.
	MaxResults *int32

	// The token for the next page of results.
	NextToken *string

	// The ID of the Traffic Mirror session.
	TrafficMirrorSessionIds []*string
}

type DescribeTrafficMirrorSessionsOutput struct {

	// The token to use to retrieve the next page of results. The value is null when
	// there are no more results to return.
	NextToken *string

	// Describes one or more Traffic Mirror sessions. By default, all Traffic Mirror
	// sessions are described. Alternatively, you can filter the results.
	TrafficMirrorSessions []*types.TrafficMirrorSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeTrafficMirrorSessionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeTrafficMirrorSessions{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeTrafficMirrorSessions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeTrafficMirrorSessions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeTrafficMirrorSessions",
	}
}
