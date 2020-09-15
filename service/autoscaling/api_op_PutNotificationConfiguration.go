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

// Configures an Auto Scaling group to send notifications when specified events
// take place. Subscribers to the specified topic can have messages delivered to an
// endpoint such as a web server or an email address. This configuration overwrites
// any existing configuration. For more information, see Getting Amazon SNS
// Notifications When Your Auto Scaling Group Scales
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/ASGettingNotifications.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) PutNotificationConfiguration(ctx context.Context, params *PutNotificationConfigurationInput, optFns ...func(*Options)) (*PutNotificationConfigurationOutput, error) {
	stack := middleware.NewStack("PutNotificationConfiguration", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpPutNotificationConfigurationMiddlewares(stack)
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
	addOpPutNotificationConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutNotificationConfiguration(options.Region), middleware.Before)

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
			OperationName: "PutNotificationConfiguration",
			Err:           err,
		}
	}
	out := result.(*PutNotificationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutNotificationConfigurationInput struct {
	// The name of the Auto Scaling group.
	AutoScalingGroupName *string
	// The Amazon Resource Name (ARN) of the Amazon Simple Notification Service (Amazon
	// SNS) topic.
	TopicARN *string
	// The type of event that causes the notification to be sent. To query the
	// notification types supported by Amazon EC2 Auto Scaling, call the
	// DescribeAutoScalingNotificationTypes () API.
	NotificationTypes []*string
}

type PutNotificationConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpPutNotificationConfigurationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpPutNotificationConfiguration{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpPutNotificationConfiguration{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutNotificationConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "PutNotificationConfiguration",
	}
}