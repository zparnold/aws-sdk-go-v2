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
)

// Describes the Regions that are enabled for your account, or all Regions. For a
// list of the Regions supported by Amazon EC2, see  Regions and Endpoints
// (https://docs.aws.amazon.com/general/latest/gr/rande.html#ec2_region). For
// information about enabling and disabling Regions for your account, see Managing
// AWS Regions (https://docs.aws.amazon.com/general/latest/gr/rande-manage.html) in
// the AWS General Reference.
func (c *Client) DescribeRegions(ctx context.Context, params *DescribeRegionsInput, optFns ...func(*Options)) (*DescribeRegionsOutput, error) {
	stack := middleware.NewStack("DescribeRegions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeRegionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeRegions(options.Region), middleware.Before)

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
			OperationName: "DescribeRegions",
			Err:           err,
		}
	}
	out := result.(*DescribeRegionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeRegionsInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The names of the Regions. You can specify any Regions, whether they are enabled
	// and disabled for your account.
	RegionNames []*string
	// Indicates whether to display all Regions, including Regions that are disabled
	// for your account.
	AllRegions *bool
	// The filters.
	//
	//     * endpoint - The endpoint of the Region (for example,
	// ec2.us-east-1.amazonaws.com).
	//
	//     * opt-in-status - The opt-in status of the
	// Region (opt-in-not-required | opted-in | not-opted-in).
	//
	//     * region-name - The
	// name of the Region (for example, us-east-1).
	Filters []*types.Filter
}

type DescribeRegionsOutput struct {
	// Information about the Regions.
	Regions []*types.Region

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeRegionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeRegions{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeRegions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeRegions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeRegions",
	}
}