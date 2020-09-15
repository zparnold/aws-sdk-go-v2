// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes an automatic scaling policy from a specified instance group within an
// EMR cluster.
func (c *Client) RemoveAutoScalingPolicy(ctx context.Context, params *RemoveAutoScalingPolicyInput, optFns ...func(*Options)) (*RemoveAutoScalingPolicyOutput, error) {
	stack := middleware.NewStack("RemoveAutoScalingPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpRemoveAutoScalingPolicyMiddlewares(stack)
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
	addOpRemoveAutoScalingPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRemoveAutoScalingPolicy(options.Region), middleware.Before)

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
			OperationName: "RemoveAutoScalingPolicy",
			Err:           err,
		}
	}
	out := result.(*RemoveAutoScalingPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RemoveAutoScalingPolicyInput struct {
	// Specifies the ID of a cluster. The instance group to which the automatic scaling
	// policy is applied is within this cluster.
	ClusterId *string
	// Specifies the ID of the instance group to which the scaling policy is applied.
	InstanceGroupId *string
}

type RemoveAutoScalingPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpRemoveAutoScalingPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpRemoveAutoScalingPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpRemoveAutoScalingPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opRemoveAutoScalingPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "RemoveAutoScalingPolicy",
	}
}