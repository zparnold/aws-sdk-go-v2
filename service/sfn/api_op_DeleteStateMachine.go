// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a state machine. This is an asynchronous operation: It sets the state
// machine's status to DELETING and begins the deletion process. For EXPRESSstate
// machines, the deletion will happen eventually (usually less than a minute).
// Running executions may emit logs after DeleteStateMachine API is called.
func (c *Client) DeleteStateMachine(ctx context.Context, params *DeleteStateMachineInput, optFns ...func(*Options)) (*DeleteStateMachineOutput, error) {
	stack := middleware.NewStack("DeleteStateMachine", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpDeleteStateMachineMiddlewares(stack)
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
	addOpDeleteStateMachineValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteStateMachine(options.Region), middleware.Before)

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
			OperationName: "DeleteStateMachine",
			Err:           err,
		}
	}
	out := result.(*DeleteStateMachineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteStateMachineInput struct {
	// The Amazon Resource Name (ARN) of the state machine to delete.
	StateMachineArn *string
}

type DeleteStateMachineOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpDeleteStateMachineMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpDeleteStateMachine{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpDeleteStateMachine{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteStateMachine(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "DeleteStateMachine",
	}
}
