// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Sets the size of the specified Auto Scaling group. If a scale-in activity occurs
// as a result of a new DesiredCapacity value that is lower than the current size
// of the group, the Auto Scaling group uses its termination policy to determine
// which instances to terminate. For more information, see Manual Scaling
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-manual-scaling.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) SetDesiredCapacity(ctx context.Context, params *SetDesiredCapacityInput, optFns ...func(*Options)) (*SetDesiredCapacityOutput, error) {
	stack := middleware.NewStack("SetDesiredCapacity", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpSetDesiredCapacityMiddlewares(stack)
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
	addOpSetDesiredCapacityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetDesiredCapacity(options.Region), middleware.Before)

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
			OperationName: "SetDesiredCapacity",
			Err:           err,
		}
	}
	out := result.(*SetDesiredCapacityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetDesiredCapacityInput struct {
	// Indicates whether Amazon EC2 Auto Scaling waits for the cooldown period to
	// complete before initiating a scaling activity to set your Auto Scaling group to
	// its new capacity. By default, Amazon EC2 Auto Scaling does not honor the
	// cooldown period during manual scaling activities.
	HonorCooldown *bool
	// The desired capacity is the initial capacity of the Auto Scaling group after
	// this operation completes and the capacity it attempts to maintain.
	DesiredCapacity *int32
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string
}

type SetDesiredCapacityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpSetDesiredCapacityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpSetDesiredCapacity{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpSetDesiredCapacity{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetDesiredCapacity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "SetDesiredCapacity",
	}
}
