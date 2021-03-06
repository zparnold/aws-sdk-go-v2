// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts an execution of the workflow type in the specified domain using the
// provided workflowId and input data.  <p>This action returns the newly started
// workflow execution.</p> <p> <b>Access Control</b> </p> <p>You can use IAM
// policies to control this action's access to Amazon SWF resources as follows:</p>
// <ul> <li> <p>Use a <code>Resource</code> element with the domain name to limit
// the action to only specified domains.</p> </li> <li> <p>Use an
// <code>Action</code> element to allow or deny permission to call this action.</p>
// </li> <li> <p>Constrain the following parameters by using a
// <code>Condition</code> element with the appropriate keys.</p> <ul> <li> <p>
// <code>tagList.member.0</code>: The key is <code>swf:tagList.member.0</code>.</p>
// </li> <li> <p> <code>tagList.member.1</code>: The key is
// <code>swf:tagList.member.1</code>.</p> </li> <li> <p>
// <code>tagList.member.2</code>: The key is <code>swf:tagList.member.2</code>.</p>
// </li> <li> <p> <code>tagList.member.3</code>: The key is
// <code>swf:tagList.member.3</code>.</p> </li> <li> <p>
// <code>tagList.member.4</code>: The key is <code>swf:tagList.member.4</code>.</p>
// </li> <li> <p> <code>taskList</code>: String constraint. The key is
// <code>swf:taskList.name</code>.</p> </li> <li> <p>
// <code>workflowType.name</code>: String constraint. The key is
// <code>swf:workflowType.name</code>.</p> </li> <li> <p>
// <code>workflowType.version</code>: String constraint. The key is
// <code>swf:workflowType.version</code>.</p> </li> </ul> </li> </ul> <p>If the
// caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's <code>cause</code> parameter is set to
// <code>OPERATION_NOT_PERMITTED</code>. For details and example IAM policies, see
// <a
// href="https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html">Using
// IAM to Manage Access to Amazon SWF Workflows</a> in the <i>Amazon SWF Developer
// Guide</i>.</p>
func (c *Client) StartWorkflowExecution(ctx context.Context, params *StartWorkflowExecutionInput, optFns ...func(*Options)) (*StartWorkflowExecutionOutput, error) {
	stack := middleware.NewStack("StartWorkflowExecution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpStartWorkflowExecutionMiddlewares(stack)
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
	addOpStartWorkflowExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartWorkflowExecution(options.Region), middleware.Before)
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
			OperationName: "StartWorkflowExecution",
			Err:           err,
		}
	}
	out := result.(*StartWorkflowExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartWorkflowExecutionInput struct {

	// The name of the domain in which the workflow execution is created.
	//
	// This member is required.
	Domain *string

	// The user defined identifier associated with the workflow execution. You can use
	// this to associate a custom identifier with the workflow execution. You may
	// specify the same identifier if a workflow execution is logically a restart of a
	// previous execution. You cannot have two open workflow executions with the same
	// workflowId at the same time within the same domain.  <p>The specified string
	// must not start or end with whitespace. It must not contain a <code>:</code>
	// (colon), <code>/</code> (slash), <code>|</code> (vertical bar), or any control
	// characters (<code>\u0000-\u001f</code> | <code>\u007f-\u009f</code>). Also, it
	// must not <i>be</i> the literal string <code>arn</code>.</p>
	//
	// This member is required.
	WorkflowId *string

	// The type of the workflow to start.
	//
	// This member is required.
	WorkflowType *types.WorkflowType

	// If set, specifies the policy to use for the child workflow executions of this
	// workflow execution if it is terminated, by calling the
	// TerminateWorkflowExecution () action explicitly or due to an expired timeout.
	// This policy overrides the default child policy specified when registering the
	// workflow type using RegisterWorkflowType (). The supported child policies are:
	//
	//
	// * TERMINATE – The child executions are terminated.
	//
	//     * REQUEST_CANCEL – A
	// request to cancel is attempted for each child execution by recording a
	// WorkflowExecutionCancelRequested event in its history. It is up to the decider
	// to take appropriate actions when it receives an execution history with this
	// event.
	//
	//     * ABANDON – No action is taken. The child executions continue to
	// run.
	//
	// A child policy for this workflow execution must be specified either as a
	// default for the workflow type or through this parameter. If neither this
	// parameter is set nor a default child policy was specified at registration time
	// then a fault is returned.
	ChildPolicy types.ChildPolicy

	// The total duration for this workflow execution. This overrides the
	// defaultExecutionStartToCloseTimeout specified when registering the workflow
	// type.  <p>The duration is specified in seconds; an integer greater than or equal
	// to <code>0</code>. Exceeding this limit causes the workflow execution to time
	// out. Unlike some of the other timeout parameters in Amazon SWF, you cannot
	// specify a value of "NONE" for this timeout; there is a one-year max limit on the
	// time that a workflow execution can run.</p> <note> <p>An execution
	// start-to-close timeout must be specified either through this parameter or as a
	// default when the workflow type is registered. If neither this parameter nor a
	// default execution start-to-close timeout is specified, a fault is returned.</p>
	// </note>
	ExecutionStartToCloseTimeout *string

	// The input for the workflow execution. This is a free form string which should be
	// meaningful to the workflow you are starting. This input is made available to the
	// new workflow execution in the WorkflowExecutionStarted history event.
	Input *string

	// The IAM role to attach to this workflow execution. Executions of this workflow
	// type need IAM roles to invoke Lambda functions. If you don't attach an IAM role,
	// any attempt to schedule a Lambda task fails. This results in a
	// ScheduleLambdaFunctionFailed history event. For more information, see
	// https://docs.aws.amazon.com/amazonswf/latest/developerguide/lambda-task.html
	// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/lambda-task.html)
	// in the Amazon SWF Developer Guide.
	LambdaRole *string

	// The list of tags to associate with the workflow execution. You can specify a
	// maximum of 5 tags. You can list workflow executions with a specific tag by
	// calling ListOpenWorkflowExecutions () or ListClosedWorkflowExecutions () and
	// specifying a TagFilter ().
	TagList []*string

	// The task list to use for the decision tasks generated for this workflow
	// execution. This overrides the defaultTaskList specified when registering the
	// workflow type. A task list for this workflow execution must be specified either
	// as a default for the workflow type or through this parameter. If neither this
	// parameter is set nor a default task list was specified at registration time then
	// a fault is returned.  <p>The specified string must not start or end with
	// whitespace. It must not contain a <code>:</code> (colon), <code>/</code>
	// (slash), <code>|</code> (vertical bar), or any control characters
	// (<code>\u0000-\u001f</code> | <code>\u007f-\u009f</code>). Also, it must not
	// <i>be</i> the literal string <code>arn</code>.</p>
	TaskList *types.TaskList

	// The task priority to use for this workflow execution. This overrides any default
	// priority that was assigned when the workflow type was registered. If not set,
	// then the default task priority for the workflow type is used. Valid values are
	// integers that range from Java's Integer.MIN_VALUE (-2147483648) to
	// Integer.MAX_VALUE (2147483647). Higher numbers indicate higher priority. For
	// more information about setting task priority, see Setting Task Priority
	// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/programming-priority.html)
	// in the Amazon SWF Developer Guide.
	TaskPriority *string

	// Specifies the maximum duration of decision tasks for this workflow execution.
	// This parameter overrides the defaultTaskStartToCloseTimout specified when
	// registering the workflow type using RegisterWorkflowType (). The duration is
	// specified in seconds, an integer greater than or equal to 0. You can use NONE to
	// specify unlimited duration. A task start-to-close timeout for this workflow
	// execution must be specified either as a default for the workflow type or through
	// this parameter. If neither this parameter is set nor a default task
	// start-to-close timeout was specified at registration time then a fault is
	// returned.
	TaskStartToCloseTimeout *string
}

// Specifies the runId of a workflow execution.
type StartWorkflowExecutionOutput struct {

	// The runId of a workflow execution. This ID is generated by the service and can
	// be used to uniquely identify the workflow execution within a domain.
	RunId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpStartWorkflowExecutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpStartWorkflowExecution{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpStartWorkflowExecution{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartWorkflowExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "StartWorkflowExecution",
	}
}
