// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the tasks in a maintenance window.
func (c *Client) GetMaintenanceWindowTask(ctx context.Context, params *GetMaintenanceWindowTaskInput, optFns ...func(*Options)) (*GetMaintenanceWindowTaskOutput, error) {
	stack := middleware.NewStack("GetMaintenanceWindowTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetMaintenanceWindowTaskMiddlewares(stack)
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
	addOpGetMaintenanceWindowTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMaintenanceWindowTask(options.Region), middleware.Before)
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
			OperationName: "GetMaintenanceWindowTask",
			Err:           err,
		}
	}
	out := result.(*GetMaintenanceWindowTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMaintenanceWindowTaskInput struct {

	// The maintenance window ID that includes the task to retrieve.
	//
	// This member is required.
	WindowId *string

	// The maintenance window task ID to retrieve.
	//
	// This member is required.
	WindowTaskId *string
}

type GetMaintenanceWindowTaskOutput struct {

	// The retrieved task description.
	Description *string

	// The location in Amazon S3 where the task results are logged. LoggingInfo has
	// been deprecated. To specify an S3 bucket to contain logs, instead use the
	// OutputS3BucketName and OutputS3KeyPrefix options in the TaskInvocationParameters
	// structure. For information about how Systems Manager handles these options for
	// the supported maintenance window task types, see
	// MaintenanceWindowTaskInvocationParameters ().
	LoggingInfo *types.LoggingInfo

	// The maximum number of targets allowed to run this task in parallel.
	MaxConcurrency *string

	// The maximum number of errors allowed before the task stops being scheduled.
	MaxErrors *string

	// The retrieved task name.
	Name *string

	// The priority of the task when it runs. The lower the number, the higher the
	// priority. Tasks that have the same priority are scheduled in parallel.
	Priority *int32

	// The ARN of the IAM service role to use to publish Amazon Simple Notification
	// Service (Amazon SNS) notifications for maintenance window Run Command tasks.
	ServiceRoleArn *string

	// The targets where the task should run.
	Targets []*types.Target

	// The resource that the task used during execution. For RUN_COMMAND and AUTOMATION
	// task types, the TaskArn is the Systems Manager Document name/ARN. For LAMBDA
	// tasks, the value is the function name/ARN. For STEP_FUNCTIONS tasks, the value
	// is the state machine ARN.
	TaskArn *string

	// The parameters to pass to the task when it runs.
	TaskInvocationParameters *types.MaintenanceWindowTaskInvocationParameters

	// The parameters to pass to the task when it runs. TaskParameters has been
	// deprecated. To specify parameters to pass to a task when it runs, instead use
	// the Parameters option in the TaskInvocationParameters structure. For information
	// about how Systems Manager handles these options for the supported maintenance
	// window task types, see MaintenanceWindowTaskInvocationParameters ().
	TaskParameters map[string]*types.MaintenanceWindowTaskParameterValueExpression

	// The type of task to run.
	TaskType types.MaintenanceWindowTaskType

	// The retrieved maintenance window ID.
	WindowId *string

	// The retrieved maintenance window task ID.
	WindowTaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetMaintenanceWindowTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetMaintenanceWindowTask{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMaintenanceWindowTask{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetMaintenanceWindowTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetMaintenanceWindowTask",
	}
}
