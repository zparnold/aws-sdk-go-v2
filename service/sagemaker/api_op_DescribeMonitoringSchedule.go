// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes the schedule for a monitoring job.
func (c *Client) DescribeMonitoringSchedule(ctx context.Context, params *DescribeMonitoringScheduleInput, optFns ...func(*Options)) (*DescribeMonitoringScheduleOutput, error) {
	stack := middleware.NewStack("DescribeMonitoringSchedule", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeMonitoringScheduleMiddlewares(stack)
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
	addOpDescribeMonitoringScheduleValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMonitoringSchedule(options.Region), middleware.Before)
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
			OperationName: "DescribeMonitoringSchedule",
			Err:           err,
		}
	}
	out := result.(*DescribeMonitoringScheduleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMonitoringScheduleInput struct {

	// Name of a previously created monitoring schedule.
	//
	// This member is required.
	MonitoringScheduleName *string
}

type DescribeMonitoringScheduleOutput struct {

	// The time at which the monitoring job was created.
	//
	// This member is required.
	CreationTime *time.Time

	// The time at which the monitoring job was last modified.
	//
	// This member is required.
	LastModifiedTime *time.Time

	// The Amazon Resource Name (ARN) of the monitoring schedule.
	//
	// This member is required.
	MonitoringScheduleArn *string

	// The configuration object that specifies the monitoring schedule and defines the
	// monitoring job.
	//
	// This member is required.
	MonitoringScheduleConfig *types.MonitoringScheduleConfig

	// Name of the monitoring schedule.
	//
	// This member is required.
	MonitoringScheduleName *string

	// The status of an monitoring job.
	//
	// This member is required.
	MonitoringScheduleStatus types.ScheduleStatus

	// The name of the endpoint for the monitoring job.
	EndpointName *string

	// A string, up to one KB in size, that contains the reason a monitoring job
	// failed, if it failed.
	FailureReason *string

	// Describes metadata on the last execution to run, if there was one.
	LastMonitoringExecutionSummary *types.MonitoringExecutionSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeMonitoringScheduleMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMonitoringSchedule{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMonitoringSchedule{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeMonitoringSchedule(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DescribeMonitoringSchedule",
	}
}
