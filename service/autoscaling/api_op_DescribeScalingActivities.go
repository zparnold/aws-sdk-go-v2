// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more scaling activities for the specified Auto Scaling group.
func (c *Client) DescribeScalingActivities(ctx context.Context, params *DescribeScalingActivitiesInput, optFns ...func(*Options)) (*DescribeScalingActivitiesOutput, error) {
	stack := middleware.NewStack("DescribeScalingActivities", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribeScalingActivitiesMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScalingActivities(options.Region), middleware.Before)

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
			OperationName: "DescribeScalingActivities",
			Err:           err,
		}
	}
	out := result.(*DescribeScalingActivitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScalingActivitiesInput struct {
	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
	// The activity IDs of the desired scaling activities. You can specify up to 50
	// IDs. If you omit this parameter, all activities for the past six weeks are
	// described. If unknown activities are requested, they are ignored with no error.
	// If you specify an Auto Scaling group, the results are limited to that group.
	ActivityIds []*string
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string
	// The maximum number of items to return with this call. The default value is 100
	// and the maximum value is 100.
	MaxRecords *int32
}

type DescribeScalingActivitiesOutput struct {
	// The scaling activities. Activities are sorted by start time. Activities still in
	// progress are described first.
	Activities []*types.Activity
	// A string that indicates that the response contains more items than can be
	// returned in a single response. To receive additional items, specify this string
	// for the NextToken value when requesting the next set of items. This value is
	// null when there are no more items to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribeScalingActivitiesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribeScalingActivities{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeScalingActivities{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeScalingActivities(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeScalingActivities",
	}
}
