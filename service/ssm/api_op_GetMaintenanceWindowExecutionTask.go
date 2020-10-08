// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves the details about a specific task run as part of a maintenance window
// execution.
func (c *Client) GetMaintenanceWindowExecutionTask(ctx context.Context, params *GetMaintenanceWindowExecutionTaskInput, optFns ...func(*Options)) (*GetMaintenanceWindowExecutionTaskOutput, error) {
	if params == nil {
		params = &GetMaintenanceWindowExecutionTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMaintenanceWindowExecutionTask", params, optFns, addOperationGetMaintenanceWindowExecutionTaskMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMaintenanceWindowExecutionTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMaintenanceWindowExecutionTaskInput struct {

	// The ID of the specific task execution in the maintenance window task that should
	// be retrieved.
	//
	// This member is required.
	TaskId *string

	// The ID of the maintenance window execution that includes the task.
	//
	// This member is required.
	WindowExecutionId *string
}

type GetMaintenanceWindowExecutionTaskOutput struct {

	// The time the task execution completed.
	EndTime *time.Time

	// The defined maximum number of task executions that could be run in parallel.
	MaxConcurrency *string

	// The defined maximum number of task execution errors allowed before scheduling of
	// the task execution would have been stopped.
	MaxErrors *string

	// The priority of the task.
	Priority *int32

	// The role that was assumed when running the task.
	ServiceRole *string

	// The time the task execution started.
	StartTime *time.Time

	// The status of the task.
	Status types.MaintenanceWindowExecutionStatus

	// The details explaining the Status. Only available for certain status values.
	StatusDetails *string

	// The ARN of the task that ran.
	TaskArn *string

	// The ID of the specific task execution in the maintenance window task that was
	// retrieved.
	TaskExecutionId *string

	// The parameters passed to the task when it was run. TaskParameters has been
	// deprecated. To specify parameters to pass to a task when it runs, instead use
	// the Parameters option in the TaskInvocationParameters structure. For information
	// about how Systems Manager handles these options for the supported maintenance
	// window task types, see MaintenanceWindowTaskInvocationParameters (). The map has
	// the following format: Key: string, between 1 and 255 characters Value: an array
	// of strings, each string is between 1 and 255 characters
	TaskParameters []map[string]*types.MaintenanceWindowTaskParameterValueExpression

	// The type of task that was run.
	Type types.MaintenanceWindowTaskType

	// The ID of the maintenance window execution that includes the task.
	WindowExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetMaintenanceWindowExecutionTaskMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMaintenanceWindowExecutionTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMaintenanceWindowExecutionTask{}, middleware.After)
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
	addOpGetMaintenanceWindowExecutionTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMaintenanceWindowExecutionTask(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetMaintenanceWindowExecutionTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetMaintenanceWindowExecutionTask",
	}
}
