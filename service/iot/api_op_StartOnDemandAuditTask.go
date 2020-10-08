// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Starts an on-demand Device Defender audit.
func (c *Client) StartOnDemandAuditTask(ctx context.Context, params *StartOnDemandAuditTaskInput, optFns ...func(*Options)) (*StartOnDemandAuditTaskOutput, error) {
	if params == nil {
		params = &StartOnDemandAuditTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StartOnDemandAuditTask", params, optFns, addOperationStartOnDemandAuditTaskMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*StartOnDemandAuditTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StartOnDemandAuditTaskInput struct {

	// Which checks are performed during the audit. The checks you specify must be
	// enabled for your account or an exception occurs. Use
	// DescribeAccountAuditConfiguration to see the list of all checks, including those
	// that are enabled or UpdateAccountAuditConfiguration to select which checks are
	// enabled.
	//
	// This member is required.
	TargetCheckNames []*string
}

type StartOnDemandAuditTaskOutput struct {

	// The ID of the on-demand audit you started.
	TaskId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStartOnDemandAuditTaskMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpStartOnDemandAuditTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpStartOnDemandAuditTask{}, middleware.After)
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
	addOpStartOnDemandAuditTaskValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStartOnDemandAuditTask(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStartOnDemandAuditTask(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "StartOnDemandAuditTask",
	}
}
