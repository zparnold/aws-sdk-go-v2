// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscalingplans

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscalingplans/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the scalable resources in the specified scaling plan.
func (c *Client) DescribeScalingPlanResources(ctx context.Context, params *DescribeScalingPlanResourcesInput, optFns ...func(*Options)) (*DescribeScalingPlanResourcesOutput, error) {
	if params == nil {
		params = &DescribeScalingPlanResourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeScalingPlanResources", params, optFns, addOperationDescribeScalingPlanResourcesMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScalingPlanResourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeScalingPlanResourcesInput struct {

	// The name of the scaling plan.
	//
	// This member is required.
	ScalingPlanName *string

	// The version number of the scaling plan.
	//
	// This member is required.
	ScalingPlanVersion *int64

	// The maximum number of scalable resources to return. The value must be between 1
	// and 50. The default value is 50.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string
}

type DescribeScalingPlanResourcesOutput struct {

	// The token required to get the next set of results. This value is null if there
	// are no more results to return.
	NextToken *string

	// Information about the scalable resources.
	ScalingPlanResources []*types.ScalingPlanResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeScalingPlanResourcesMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeScalingPlanResources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeScalingPlanResources{}, middleware.After)
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
	addOpDescribeScalingPlanResourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScalingPlanResources(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeScalingPlanResources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling-plans",
		OperationName: "DescribeScalingPlanResources",
	}
}
