// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworkscm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an application-level backup of a server. While the server is in the
// BACKING_UP state, the server cannot be changed, and no additional backup can be
// created. Backups can be created for servers in RUNNING, HEALTHY, and UNHEALTHY
// states. By default, you can create a maximum of 50 manual backups. This
// operation is asynchronous. A LimitExceededException is thrown when the maximum
// number of manual backups is reached. An InvalidStateException is thrown when the
// server is not in any of the following states: RUNNING, HEALTHY, or UNHEALTHY. A
// ResourceNotFoundException is thrown when the server is not found. A
// ValidationException is thrown when parameters of the request are not valid.
func (c *Client) CreateBackup(ctx context.Context, params *CreateBackupInput, optFns ...func(*Options)) (*CreateBackupOutput, error) {
	stack := middleware.NewStack("CreateBackup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateBackupMiddlewares(stack)
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
	addOpCreateBackupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBackup(options.Region), middleware.Before)
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
			OperationName: "CreateBackup",
			Err:           err,
		}
	}
	out := result.(*CreateBackupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBackupInput struct {

	// The name of the server that you want to back up.
	//
	// This member is required.
	ServerName *string

	// A user-defined description of the backup.
	Description *string

	// A map that contains tag keys and tag values to attach to an AWS OpsWorks-CM
	// server backup.
	//
	//     * The key cannot be empty.
	//
	//     * The key can be a maximum
	// of 127 characters, and can contain only Unicode letters, numbers, or separators,
	// or the following special characters: + - = . _ : /
	//
	//     * The value can be a
	// maximum 255 characters, and contain only Unicode letters, numbers, or
	// separators, or the following special characters: + - = . _ : /
	//
	//     * Leading
	// and trailing white spaces are trimmed from both the key and value.
	//
	//     * A
	// maximum of 50 user-applied tags is allowed for tag-supported AWS OpsWorks-CM
	// resources.
	Tags []*types.Tag
}

type CreateBackupOutput struct {

	// Backup created by request.
	Backup *types.Backup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateBackupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateBackup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateBackup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateBackup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks-cm",
		OperationName: "CreateBackup",
	}
}
