// Code generated by smithy-go-codegen DO NOT EDIT.

package fsx

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/fsx/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Use this operation to update the configuration of an existing Amazon FSx file
// system. You can update multiple properties in a single request. For Amazon FSx
// for Windows File Server file systems, you can update the following properties:
//
//
// * AutomaticBackupRetentionDays
//
//     * DailyAutomaticBackupStartTime
//
//     *
// SelfManagedActiveDirectoryConfiguration
//
//     * StorageCapacity
//
//     *
// ThroughputCapacity
//
//     * WeeklyMaintenanceStartTime
//
// For Amazon FSx for Lustre
// file systems, you can update the following properties:
//
//     * AutoImportPolicy
//
//
// * AutomaticBackupRetentionDays
//
//     * DailyAutomaticBackupStartTime
//
//     *
// WeeklyMaintenanceStartTime
func (c *Client) UpdateFileSystem(ctx context.Context, params *UpdateFileSystemInput, optFns ...func(*Options)) (*UpdateFileSystemOutput, error) {
	stack := middleware.NewStack("UpdateFileSystem", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateFileSystemMiddlewares(stack)
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
	addIdempotencyToken_opUpdateFileSystemMiddleware(stack, options)
	addOpUpdateFileSystemValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateFileSystem(options.Region), middleware.Before)

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
			OperationName: "UpdateFileSystem",
			Err:           err,
		}
	}
	out := result.(*UpdateFileSystemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The request object for the UpdateFileSystem operation.
type UpdateFileSystemInput struct {
	// A string of up to 64 ASCII characters that Amazon FSx uses to ensure idempotent
	// updates. This string is automatically filled on your behalf when you use the AWS
	// Command Line Interface (AWS CLI) or an AWS SDK.
	ClientRequestToken *string
	// Use this parameter to increase the storage capacity of an Amazon FSx for Windows
	// File Server file system. Specifies the storage capacity target value, GiB, for
	// the file system you're updating. The storage capacity target value must be at
	// least 10 percent (%) greater than the current storage capacity value. In order
	// to increase storage capacity, the file system needs to have at least 16 MB/s of
	// throughput capacity. You cannot make a storage capacity increase request if
	// there is an existing storage capacity increase request in progress. For more
	// information, see Managing Storage Capacity
	// (https://docs.aws.amazon.com/fsx/latest/WindowsGuide/managing-storage-capacity.html).
	StorageCapacity *int32
	// Identifies the file system that you are updating.
	FileSystemId *string
	// The configuration updates for an Amazon FSx for Windows File Server file system.
	WindowsConfiguration *types.UpdateFileSystemWindowsConfiguration
	// The configuration object for Amazon FSx for Lustre file systems used in the
	// UpdateFileSystem operation.
	LustreConfiguration *types.UpdateFileSystemLustreConfiguration
}

// The response object for the UpdateFileSystem operation.
type UpdateFileSystemOutput struct {
	// A description of the file system that was updated.
	FileSystem *types.FileSystem

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateFileSystemMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateFileSystem{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateFileSystem{}, middleware.After)
}

type idempotencyToken_initializeOpUpdateFileSystem struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateFileSystem) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateFileSystem) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateFileSystemInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateFileSystemInput ")
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
func addIdempotencyToken_opUpdateFileSystemMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpUpdateFileSystem{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateFileSystem(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fsx",
		OperationName: "UpdateFileSystem",
	}
}