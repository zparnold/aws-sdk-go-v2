// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Defines an action that can be applied to audit findings by using
// StartAuditMitigationActionsTask. Each mitigation action can apply only one type
// of change.
func (c *Client) CreateMitigationAction(ctx context.Context, params *CreateMitigationActionInput, optFns ...func(*Options)) (*CreateMitigationActionOutput, error) {
	stack := middleware.NewStack("CreateMitigationAction", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateMitigationActionMiddlewares(stack)
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
	addOpCreateMitigationActionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMitigationAction(options.Region), middleware.Before)

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
			OperationName: "CreateMitigationAction",
			Err:           err,
		}
	}
	out := result.(*CreateMitigationActionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMitigationActionInput struct {
	// Metadata that can be used to manage the mitigation action.
	Tags []*types.Tag
	// A friendly name for the action. Choose a friendly name that accurately describes
	// the action (for example, EnableLoggingAction).
	ActionName *string
	// The ARN of the IAM role that is used to apply the mitigation action.
	RoleArn *string
	// Defines the type of action and the parameters for that action.
	ActionParams *types.MitigationActionParams
}

type CreateMitigationActionOutput struct {
	// A unique identifier for the new mitigation action.
	ActionId *string
	// The ARN for the new mitigation action.
	ActionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateMitigationActionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateMitigationAction{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMitigationAction{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateMitigationAction(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateMitigationAction",
	}
}
