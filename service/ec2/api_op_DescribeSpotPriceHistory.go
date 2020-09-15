// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes the Spot price history. For more information, see Spot Instance
// pricing history
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-spot-instances-history.html)
// in the Amazon EC2 User Guide for Linux Instances. When you specify a start and
// end time, this operation returns the prices of the instance types within the
// time range that you specified and the time when the price changed. The price is
// valid within the time period that you specified; the response merely indicates
// the last time that the price changed.
func (c *Client) DescribeSpotPriceHistory(ctx context.Context, params *DescribeSpotPriceHistoryInput, optFns ...func(*Options)) (*DescribeSpotPriceHistoryOutput, error) {
	stack := middleware.NewStack("DescribeSpotPriceHistory", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeSpotPriceHistoryMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSpotPriceHistory(options.Region), middleware.Before)

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
			OperationName: "DescribeSpotPriceHistory",
			Err:           err,
		}
	}
	out := result.(*DescribeSpotPriceHistoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribeSpotPriceHistory.
type DescribeSpotPriceHistoryInput struct {
	// The date and time, up to the current date, from which to stop retrieving the
	// price history data, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ).
	EndTime *time.Time
	// Filters the results by the specified instance types.
	InstanceTypes []types.InstanceType
	// The date and time, up to the past 90 days, from which to start retrieving the
	// price history data, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ).
	StartTime *time.Time
	// The token for the next set of results.
	NextToken *string
	// The maximum number of results to return in a single call. Specify a value
	// between 1 and 1000. The default value is 1000. To retrieve the remaining
	// results, make another call with the returned NextToken value.
	MaxResults *int32
	// Filters the results by the specified Availability Zone.
	AvailabilityZone *string
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// One or more filters.
	//
	//     * availability-zone - The Availability Zone for which
	// prices should be returned.
	//
	//     * instance-type - The type of instance (for
	// example, m3.medium).
	//
	//     * product-description - The product description for
	// the Spot price (Linux/UNIX | SUSE Linux | Windows | Linux/UNIX (Amazon VPC) |
	// SUSE Linux (Amazon VPC) | Windows (Amazon VPC)).
	//
	//     * spot-price - The Spot
	// price. The value must match exactly (or use wildcards; greater than or less than
	// comparison is not supported).
	//
	//     * timestamp - The time stamp of the Spot
	// price history, in UTC format (for example, YYYY-MM-DDTHH:MM:SSZ). You can use
	// wildcards (* and ?). Greater than or less than comparison is not supported.
	Filters []*types.Filter
	// Filters the results by the specified basic product descriptions.
	ProductDescriptions []*string
}

// Contains the output of DescribeSpotPriceHistory.
type DescribeSpotPriceHistoryOutput struct {
	// The token required to retrieve the next set of results. This value is null or an
	// empty string when there are no more results to return.
	NextToken *string
	// The historical Spot prices.
	SpotPriceHistory []*types.SpotPrice

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeSpotPriceHistoryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeSpotPriceHistory{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeSpotPriceHistory{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeSpotPriceHistory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeSpotPriceHistory",
	}
}
