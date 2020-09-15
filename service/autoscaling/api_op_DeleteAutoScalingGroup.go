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

// Deletes the specified Auto Scaling group. If the group has instances or scaling
// activities in progress, you must specify the option to force the deletion in
// order for it to succeed. If the group has policies, deleting the group deletes
// the policies, the underlying alarm actions, and any alarm that no longer has an
// associated action. To remove instances from the Auto Scaling group before
// deleting it, call the DetachInstances () API with the list of instances and the
// option to decrement the desired capacity. This ensures that Amazon EC2 Auto
// Scaling does not launch replacement instances. To terminate all instances before
// deleting the Auto Scaling group, call the UpdateAutoScalingGroup () API and set
// the minimum size and desired capacity of the Auto Scaling group to zero.
func (c *Client) DeleteAutoScalingGroup(ctx context.Context, params *DeleteAutoScalingGroupInput, optFns ...func(*Options)) (*DeleteAutoScalingGroupOutput, error) {
	stack := middleware.NewStack("DeleteAutoScalingGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDeleteAutoScalingGroupMiddlewares(stack)
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
	addOpDeleteAutoScalingGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAutoScalingGroup(options.Region), middleware.Before)

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
			OperationName: "DeleteAutoScalingGroup",
			Err:           err,
		}
	}
	out := result.(*DeleteAutoScalingGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAutoScalingGroupInput struct {
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string
	// Specifies that the group is to be deleted along with all instances associated
	// with the group, without waiting for all instances to be terminated. This
	// parameter also deletes any lifecycle actions associated with the group.
	ForceDelete *bool
}

type DeleteAutoScalingGroupOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDeleteAutoScalingGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDeleteAutoScalingGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteAutoScalingGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteAutoScalingGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DeleteAutoScalingGroup",
	}
}
