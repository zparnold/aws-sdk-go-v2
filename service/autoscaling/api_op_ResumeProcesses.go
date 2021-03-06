// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resumes the specified suspended automatic scaling processes, or all suspended
// process, for the specified Auto Scaling group. For more information, see
// Suspending and Resuming Scaling Processes
// (https://docs.aws.amazon.com/autoscaling/ec2/userguide/as-suspend-resume-processes.html)
// in the Amazon EC2 Auto Scaling User Guide.
func (c *Client) ResumeProcesses(ctx context.Context, params *ResumeProcessesInput, optFns ...func(*Options)) (*ResumeProcessesOutput, error) {
	stack := middleware.NewStack("ResumeProcesses", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpResumeProcessesMiddlewares(stack)
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
	addOpResumeProcessesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opResumeProcesses(options.Region), middleware.Before)
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
			OperationName: "ResumeProcesses",
			Err:           err,
		}
	}
	out := result.(*ResumeProcessesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ResumeProcessesInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// One or more of the following processes:
	//
	//     * Launch
	//
	//     * Terminate
	//
	//     *
	// AddToLoadBalancer
	//
	//     * AlarmNotification
	//
	//     * AZRebalance
	//
	//     *
	// HealthCheck
	//
	//     * InstanceRefresh
	//
	//     * ReplaceUnhealthy
	//
	//     *
	// ScheduledActions
	//
	// If you omit this parameter, all processes are specified.
	ScalingProcesses []*string
}

type ResumeProcessesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpResumeProcessesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpResumeProcesses{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpResumeProcesses{}, middleware.After)
}

func newServiceMetadataMiddleware_opResumeProcesses(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "ResumeProcesses",
	}
}
