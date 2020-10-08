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

// Retrieves information about a specific task running on a specific target.
func (c *Client) GetMaintenanceWindowExecutionTaskInvocation(ctx context.Context, params *GetMaintenanceWindowExecutionTaskInvocationInput, optFns ...func(*Options)) (*GetMaintenanceWindowExecutionTaskInvocationOutput, error) {
	if params == nil {
		params = &GetMaintenanceWindowExecutionTaskInvocationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMaintenanceWindowExecutionTaskInvocation", params, optFns, addOperationGetMaintenanceWindowExecutionTaskInvocationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMaintenanceWindowExecutionTaskInvocationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMaintenanceWindowExecutionTaskInvocationInput struct {

	// The invocation ID to retrieve.
	//
	// This member is required.
	InvocationId *string

	// The ID of the specific task in the maintenance window task that should be
	// retrieved.
	//
	// This member is required.
	TaskId *string

	// The ID of the maintenance window execution for which the task is a part.
	//
	// This member is required.
	WindowExecutionId *string
}

type GetMaintenanceWindowExecutionTaskInvocationOutput struct {

	// The time that the task finished running on the target.
	EndTime *time.Time

	// The execution ID.
	ExecutionId *string

	// The invocation ID.
	InvocationId *string

	// User-provided value to be included in any CloudWatch events raised while running
	// tasks for these targets in this maintenance window.
	OwnerInformation *string

	// The parameters used at the time that the task ran.
	Parameters *string

	// The time that the task started running on the target.
	StartTime *time.Time

	// The task status for an invocation.
	Status types.MaintenanceWindowExecutionStatus

	// The details explaining the status. Details are only available for certain status
	// values.
	StatusDetails *string

	// The task execution ID.
	TaskExecutionId *string

	// Retrieves the task type for a maintenance window. Task types include the
	// following: LAMBDA, STEP_FUNCTIONS, AUTOMATION, RUN_COMMAND.
	TaskType types.MaintenanceWindowTaskType

	// The maintenance window execution ID.
	WindowExecutionId *string

	// The maintenance window target ID.
	WindowTargetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetMaintenanceWindowExecutionTaskInvocationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetMaintenanceWindowExecutionTaskInvocation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMaintenanceWindowExecutionTaskInvocation{}, middleware.After)
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
	addOpGetMaintenanceWindowExecutionTaskInvocationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMaintenanceWindowExecutionTaskInvocation(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetMaintenanceWindowExecutionTaskInvocation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetMaintenanceWindowExecutionTaskInvocation",
	}
}
