// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an Amazon FSx backup, deleting its contents. After deletion, the backup
// no longer exists, and its data is gone.  <p>The <code>DeleteBackup</code> call
// returns instantly. The backup will not show up in later
// <code>DescribeBackups</code> calls.</p> <important> <p>The data in a deleted
// backup is also deleted and can't be recovered by any means.</p> </important>
func (c *Client) DeleteBackup(ctx context.Context, params *DeleteBackupInput, optFns ...func(*Options)) (*DeleteBackupOutput, error) {
	if params == nil {
		params = &DeleteBackupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteBackup", params, optFns, addOperationDeleteBackupMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteBackupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object for DeleteBackup operation.
type DeleteBackupInput struct {

	// The ID of the backup you want to delete.
	//
	// This member is required.
	BackupId *string

	// A string of up to 64 ASCII characters that Amazon FSx uses to ensure idempotent
	// deletion. This is automatically filled on your behalf when using the AWS CLI or
	// SDK.
	ClientRequestToken *string
}

// The response object for DeleteBackup operation.
type DeleteBackupOutput struct {

	// The ID of the backup deleted.
	BackupId *string

	// The lifecycle of the backup. Should be DELETED.
	Lifecycle types.BackupLifecycle

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteBackupMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteBackup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteBackup{}, middleware.After)
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
	addIdempotencyToken_opDeleteBackupMiddleware(stack, options)
	addOpDeleteBackupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteBackup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpDeleteBackup struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteBackup) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteBackup) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteBackupInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteBackupInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDeleteBackupMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpDeleteBackup{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteBackup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "DeleteBackup",
	}
}
