// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Initiates execution of an Automation document.
func (c *Client) StartAutomationExecution(ctx context.Context, params *StartAutomationExecutionInput, optFns ...func(*Options)) (*StartAutomationExecutionOutput, error) {
	stack := middleware.NewStack("StartAutomationExecution", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpStartAutomationExecutionMiddlewares(stack)
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
	addOpStartAutomationExecutionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartAutomationExecution(options.Region), middleware.Before)

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
			OperationName: "StartAutomationExecution",
			Err:           err,
		}
	}
	out := result.(*StartAutomationExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartAutomationExecutionInput struct {
	// Optional metadata that you assign to a resource. You can specify a maximum of
	// five tags for an automation. Tags enable you to categorize a resource in
	// different ways, such as by purpose, owner, or environment. For example, you
	// might want to tag an automation to identify an environment or operating system.
	// In this case, you could specify the following key name/value pairs:
	//
	//     *
	// Key=environment,Value=test
	//
	//     * Key=OS,Value=Windows
	//
	// To add tags to an
	// existing patch baseline, use the AddTagsToResource () action.
	Tags []*types.Tag
	// The name of the Automation document to use for this execution.
	DocumentName *string
	// The number of errors that are allowed before the system stops running the
	// automation on additional targets. You can specify either an absolute number of
	// errors, for example 10, or a percentage of the target set, for example 10%. If
	// you specify 3, for example, the system stops running the automation when the
	// fourth error is received. If you specify 0, then the system stops running the
	// automation on additional targets after the first error result is returned. If
	// you run an automation on 50 resources and set max-errors to 10%, then the system
	// stops running the automation on additional targets when the sixth error is
	// received. Executions that are already running an automation when max-errors is
	// reached are allowed to complete, but some of these executions may fail as well.
	// If you need to ensure that there won't be more than max-errors failed
	// executions, set max-concurrency to 1 so the executions proceed one at a time.
	MaxErrors *string
	// A location is a combination of AWS Regions and/or AWS accounts where you want to
	// run the Automation. Use this action to start an Automation in multiple Regions
	// and multiple accounts. For more information, see Running Automation workflows in
	// multiple AWS Regions and accounts
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-automation-multiple-accounts-and-regions.html)
	// in the AWS Systems Manager User Guide.
	TargetLocations []*types.TargetLocation
	// A key-value mapping to target resources. Required if you specify
	// TargetParameterName.
	Targets []*types.Target
	// User-provided idempotency token. The token must be unique, is case insensitive,
	// enforces the UUID format, and can't be reused.
	ClientToken *string
	// A key-value mapping of document parameters to target resources. Both Targets and
	// TargetMaps cannot be specified together.
	TargetMaps []map[string][]*string
	// The version of the Automation document to use for this execution.
	DocumentVersion *string
	// The execution mode of the automation. Valid modes include the following: Auto
	// and Interactive. The default mode is Auto.
	Mode types.ExecutionMode
	// The maximum number of targets allowed to run this task in parallel. You can
	// specify a number, such as 10, or a percentage, such as 10%. The default value is
	// 10.
	MaxConcurrency *string
	// A key-value map of execution parameters, which match the declared parameters in
	// the Automation document.
	Parameters map[string][]*string
	// The name of the parameter used as the target resource for the rate-controlled
	// execution. Required if you specify targets.
	TargetParameterName *string
}

type StartAutomationExecutionOutput struct {
	// The unique ID of a newly scheduled automation execution.
	AutomationExecutionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpStartAutomationExecutionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpStartAutomationExecution{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpStartAutomationExecution{}, middleware.After)
}

func newServiceMetadataMiddleware_opStartAutomationExecution(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "StartAutomationExecution",
	}
}
