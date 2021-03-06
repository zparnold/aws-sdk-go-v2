// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Cancels a mitigation action task that is in progress. If the task is not in
// progress, an InvalidRequestException occurs.
func (c *Client) CancelAuditMitigationActionsTask(ctx context.Context, params *CancelAuditMitigationActionsTaskInput, optFns ...func(*Options)) (*CancelAuditMitigationActionsTaskOutput, error) {
	stack := middleware.NewStack("CancelAuditMitigationActionsTask", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCancelAuditMitigationActionsTaskMiddlewares(stack)
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
	addOpCancelAuditMitigationActionsTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCancelAuditMitigationActionsTask(options.Region), middleware.Before)
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
			OperationName: "CancelAuditMitigationActionsTask",
			Err:           err,
		}
	}
	out := result.(*CancelAuditMitigationActionsTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CancelAuditMitigationActionsTaskInput struct {

	// The unique identifier for the task that you want to cancel.
	//
	// This member is required.
	TaskId *string
}

type CancelAuditMitigationActionsTaskOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCancelAuditMitigationActionsTaskMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCancelAuditMitigationActionsTask{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCancelAuditMitigationActionsTask{}, middleware.After)
}

func newServiceMetadataMiddleware_opCancelAuditMitigationActionsTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CancelAuditMitigationActionsTask",
	}
}
