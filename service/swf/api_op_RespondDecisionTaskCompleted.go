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

// Used by deciders to tell the service that the DecisionTask () identified by the
// taskToken has successfully completed. The decisions argument specifies the list
// of decisions made while processing the task.  <p>A
// <code>DecisionTaskCompleted</code> event is added to the workflow history. The
// <code>executionContext</code> specified is attached to the event in the workflow
// execution history.</p> <p> <b>Access Control</b> </p> <p>If an IAM policy grants
// permission to use <code>RespondDecisionTaskCompleted</code>, it can express
// permissions for the list of decisions in the <code>decisions</code> parameter.
// Each of the decisions has one or more parameters, much like a regular API call.
// To allow for policies to be as readable as possible, you can express permissions
// on decisions as if they were actual API calls, including applying conditions to
// some parameters. For more information, see <a
// href="https://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dev-iam.html">Using
// IAM to Manage Access to Amazon SWF Workflows</a> in the <i>Amazon SWF Developer
// Guide</i>.</p>
func (c *Client) RespondDecisionTaskCompleted(ctx context.Context, params *RespondDecisionTaskCompletedInput, optFns ...func(*Options)) (*RespondDecisionTaskCompletedOutput, error) {
	stack := middleware.NewStack("RespondDecisionTaskCompleted", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpRespondDecisionTaskCompletedMiddlewares(stack)
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
	addOpRespondDecisionTaskCompletedValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRespondDecisionTaskCompleted(options.Region), middleware.Before)
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
			OperationName: "RespondDecisionTaskCompleted",
			Err:           err,
		}
	}
	out := result.(*RespondDecisionTaskCompletedOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input data for a TaskCompleted response to a decision task.
type RespondDecisionTaskCompletedInput struct {

	// The taskToken from the DecisionTask (). taskToken is generated by the service
	// and should be treated as an opaque value. If the task is passed to another
	// process, its taskToken must also be passed. This enables it to provide its
	// progress and respond with results.
	//
	// This member is required.
	TaskToken *string

	// The list of decisions (possibly empty) made by the decider while processing this
	// decision task. See the docs for the Decision () structure for details.
	Decisions []*types.Decision

	// User defined context to add to workflow execution.
	ExecutionContext *string
}

type RespondDecisionTaskCompletedOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpRespondDecisionTaskCompletedMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpRespondDecisionTaskCompleted{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpRespondDecisionTaskCompleted{}, middleware.After)
}

func newServiceMetadataMiddleware_opRespondDecisionTaskCompleted(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "swf",
		OperationName: "RespondDecisionTaskCompleted",
	}
}
