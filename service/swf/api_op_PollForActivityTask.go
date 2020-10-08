// Code generated by smithy-go-codegen DO NOT EDIT.

package swf

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/swf/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Used by workers to get an ActivityTask () from the specified activity taskList.
// This initiates a long poll, where the service holds the HTTP connection open and
// responds as soon as a task becomes available. The maximum time the service holds
// on to the request before responding is 60 seconds. If no task is available
// within 60 seconds, the poll returns an empty result. An empty result, in this
// context, means that an ActivityTask is returned, but that the value of taskToken
// is an empty string. If a task is returned, the worker should use its type to
// identify and process it correctly. Workers should set their client side socket
// timeout to at least 70 seconds (10 seconds higher than the maximum time service
// may hold the poll request). Access Control You can use IAM policies to control
// this action's access to Amazon SWF resources as follows:
//
//     * Use a Resource
// element with the domain name to limit the action to only specified domains.
//
//
// * Use an Action element to allow or deny permission to call this action.
//
//     *
// Constrain the taskList.name parameter by using a Condition element with the
// swf:taskList.name key to allow the action to access only certain task lists.
//
// If
// the caller doesn't have sufficient permissions to invoke the action, or the
// parameter values fall outside the specified constraints, the action fails. The
// associated event attribute's cause parameter is set to OPERATION_NOT_PERMITTED.
// For details and example IAM policies, see Using IAM to Manage Access to Amazon
// SWF Workflows
// (https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html)
// in the Amazon SWF Developer Guide.
func (c *Client) PollForActivityTask(ctx context.Context, params *PollForActivityTaskInput, optFns ...func(*Options)) (*PollForActivityTaskOutput, error) {
	if params == nil {
		params = &PollForActivityTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PollForActivityTask", params, optFns, addOperationPollForActivityTaskMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*PollForActivityTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PollForActivityTaskInput struct {

	// The name of the domain that contains the task lists being polled.
	//
	// This member is required.
	Domain *string

	// Specifies the task list to poll for activity tasks.  <p>The specified string
	// must not start or end with whitespace. It must not contain a <code>:</code>
	// (colon), <code>/</code> (slash), <code>|</code> (vertical bar), or any control
	// characters (<code>\u0000-\u001f</code> | <code>\u007f-\u009f</code>). Also, it
	// must not <i>be</i> the literal string <code>arn</code>.</p>
	//
	// This member is required.
	TaskList *types.TaskList

	// Identity of the worker making the request, recorded in the ActivityTaskStarted
	// event in the workflow history. This enables diagnostic tracing when problems
	// arise. The form of this identity is user defined.
	Identity *string
}

// Unit of work sent to an activity worker.
type PollForActivityTaskOutput struct {

	// The unique ID of the task.
	//
	// This member is required.
	ActivityId *string

	// The type of this activity task.
	//
	// This member is required.
	ActivityType *types.ActivityType

	// The ID of the ActivityTaskStarted event recorded in the history.
	//
	// This member is required.
	StartedEventId *int64

	// The opaque string used as a handle on the task. This token is used by workers to
	// communicate progress and response information back to the system about the task.
	//
	// This member is required.
	TaskToken *string

	// The workflow execution that started this activity task.
	//
	// This member is required.
	WorkflowExecution *types.WorkflowExecution

	// The inputs provided when the activity task was scheduled. The form of the input
	// is user defined and should be meaningful to the activity implementation.
	Input *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPollForActivityTaskMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpPollForActivityTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpPollForActivityTask{}, middleware.After)
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
	addOpPollForActivityTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPollForActivityTask(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPollForActivityTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "PollForActivityTask",
	}
}
