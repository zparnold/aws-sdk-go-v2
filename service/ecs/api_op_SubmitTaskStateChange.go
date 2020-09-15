// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This action is only used by the Amazon ECS agent, and it is not intended for use
// outside of the agent. Sent to acknowledge that a task changed states.
func (c *Client) SubmitTaskStateChange(ctx context.Context, params *SubmitTaskStateChangeInput, optFns ...func(*Options)) (*SubmitTaskStateChangeOutput, error) {
	stack := middleware.NewStack("SubmitTaskStateChange", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSubmitTaskStateChangeMiddlewares(stack)
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
	addOpSubmitTaskStateChangeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSubmitTaskStateChange(options.Region), middleware.Before)

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
			OperationName: "SubmitTaskStateChange",
			Err:           err,
		}
	}
	out := result.(*SubmitTaskStateChangeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SubmitTaskStateChangeInput struct {
	// The status of the state change request.
	Status *string
	// The Unix timestamp for when the task execution stopped.
	ExecutionStoppedAt *time.Time
	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts the
	// task.
	Cluster *string
	// Any attachments associated with the state change request.
	Attachments []*types.AttachmentStateChange
	// Any containers associated with the state change request.
	Containers []*types.ContainerStateChange
	// The reason for the state change request.
	Reason *string
	// The Unix timestamp for when the container image pull began.
	PullStartedAt *time.Time
	// The Unix timestamp for when the container image pull completed.
	PullStoppedAt *time.Time
	// The task ID or full ARN of the task in the state change request.
	Task *string
}

type SubmitTaskStateChangeOutput struct {
	// Acknowledgement of the state change.
	Acknowledgment *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSubmitTaskStateChangeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSubmitTaskStateChange{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSubmitTaskStateChange{}, middleware.After)
}

func newServiceMetadataMiddleware_opSubmitTaskStateChange(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "SubmitTaskStateChange",
	}
}
