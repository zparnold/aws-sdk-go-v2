// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns detailed information about command execution for an invocation or
// plugin.
func (c *Client) GetCommandInvocation(ctx context.Context, params *GetCommandInvocationInput, optFns ...func(*Options)) (*GetCommandInvocationOutput, error) {
	if params == nil {
		params = &GetCommandInvocationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCommandInvocation", params, optFns, addOperationGetCommandInvocationMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCommandInvocationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCommandInvocationInput struct {

	// (Required) The parent command ID of the invocation plugin.
	//
	// This member is required.
	CommandId *string

	// (Required) The ID of the managed instance targeted by the command. A managed
	// instance can be an EC2 instance or an instance in your hybrid environment that
	// is configured for Systems Manager.
	//
	// This member is required.
	InstanceId *string

	// (Optional) The name of the plugin for which you want detailed results. If the
	// document contains only one plugin, the name can be omitted and the details will
	// be returned. Plugin names are also referred to as step names in Systems Manager
	// documents.
	PluginName *string
}

type GetCommandInvocationOutput struct {

	// CloudWatch Logs information where Systems Manager sent the command output.
	CloudWatchOutputConfig *types.CloudWatchOutputConfig

	// The parent command ID of the invocation plugin.
	CommandId *string

	// The comment text for the command.
	Comment *string

	// The name of the document that was run. For example, AWS-RunShellScript.
	DocumentName *string

	// The SSM document version used in the request.
	DocumentVersion *string

	// Duration since ExecutionStartDateTime.
	ExecutionElapsedTime *string

	// The date and time the plugin was finished running. Date and time are written in
	// ISO 8601 format. For example, June 7, 2017 is represented as 2017-06-7. The
	// following sample AWS CLI command uses the InvokedAfter filter. aws ssm
	// list-commands --filters key=InvokedAfter,value=2017-06-07T00:00:00Z If the
	// plugin has not started to run, the string is empty.
	ExecutionEndDateTime *string

	// The date and time the plugin started running. Date and time are written in ISO
	// 8601 format. For example, June 7, 2017 is represented as 2017-06-7. The
	// following sample AWS CLI command uses the InvokedBefore filter. aws ssm
	// list-commands --filters key=InvokedBefore,value=2017-06-07T00:00:00Z If the
	// plugin has not started to run, the string is empty.
	ExecutionStartDateTime *string

	// The ID of the managed instance targeted by the command. A managed instance can
	// be an EC2 instance or an instance in your hybrid environment that is configured
	// for Systems Manager.
	InstanceId *string

	// The name of the plugin for which you want detailed results. For example,
	// aws:RunShellScript is a plugin.
	PluginName *string

	// The error level response code for the plugin script. If the response code is -1,
	// then the command has not started running on the instance, or it was not received
	// by the instance.
	ResponseCode *int32

	// The first 8,000 characters written by the plugin to stderr. If the command has
	// not finished running, then this string is empty.
	StandardErrorContent *string

	// The URL for the complete text written by the plugin to stderr. If the command
	// has not finished running, then this string is empty.
	StandardErrorUrl *string

	// The first 24,000 characters written by the plugin to stdout. If the command has
	// not finished running, if ExecutionStatus is neither Succeeded nor Failed, then
	// this string is empty.
	StandardOutputContent *string

	// The URL for the complete text written by the plugin to stdout in Amazon S3. If
	// an S3 bucket was not specified, then this string is empty.
	StandardOutputUrl *string

	// The status of this invocation plugin. This status can be different than
	// StatusDetails.
	Status types.CommandInvocationStatus

	// A detailed status of the command execution for an invocation. StatusDetails
	// includes more information than Status because it includes states resulting from
	// error and concurrency control parameters. StatusDetails can show different
	// results than Status. For more information about these statuses, see
	// Understanding command statuses
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/monitor-commands.html)
	// in the AWS Systems Manager User Guide. StatusDetails can be one of the following
	// values:  <ul> <li> <p>Pending: The command has not been sent to the
	// instance.</p> </li> <li> <p>In Progress: The command has been sent to the
	// instance but has not reached a terminal state.</p> </li> <li> <p>Delayed: The
	// system attempted to send the command to the target, but the target was not
	// available. The instance might not be available because of network issues,
	// because the instance was stopped, or for similar reasons. The system will try to
	// send the command again.</p> </li> <li> <p>Success: The command or plugin ran
	// successfully. This is a terminal state.</p> </li> <li> <p>Delivery Timed Out:
	// The command was not delivered to the instance before the delivery timeout
	// expired. Delivery timeouts do not count against the parent command's MaxErrors
	// limit, but they do contribute to whether the parent command status is Success or
	// Incomplete. This is a terminal state.</p> </li> <li> <p>Execution Timed Out: The
	// command started to run on the instance, but the execution was not complete
	// before the timeout expired. Execution timeouts count against the MaxErrors limit
	// of the parent command. This is a terminal state.</p> </li> <li> <p>Failed: The
	// command wasn't run successfully on the instance. For a plugin, this indicates
	// that the result code was not zero. For a command invocation, this indicates that
	// the result code for one or more plugins was not zero. Invocation failures count
	// against the MaxErrors limit of the parent command. This is a terminal state.</p>
	// </li> <li> <p>Canceled: The command was terminated before it was completed. This
	// is a terminal state.</p> </li> <li> <p>Undeliverable: The command can't be
	// delivered to the instance. The instance might not exist or might not be
	// responding. Undeliverable invocations don't count against the parent command's
	// MaxErrors limit and don't contribute to whether the parent command status is
	// Success or Incomplete. This is a terminal state.</p> </li> <li> <p>Terminated:
	// The parent command exceeded its MaxErrors limit and subsequent command
	// invocations were canceled by the system. This is a terminal state.</p> </li>
	// </ul>
	StatusDetails *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCommandInvocationMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetCommandInvocation{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetCommandInvocation{}, middleware.After)
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
	addOpGetCommandInvocationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetCommandInvocation(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetCommandInvocation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetCommandInvocation",
	}
}
