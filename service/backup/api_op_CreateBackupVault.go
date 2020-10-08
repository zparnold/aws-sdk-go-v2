// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Creates a logical container where backups are stored. A CreateBackupVault
// request includes a name, optionally one or more resource tags, an encryption
// key, and a request ID. Sensitive data, such as passport numbers, should not be
// included the name of a backup vault.
func (c *Client) CreateBackupVault(ctx context.Context, params *CreateBackupVaultInput, optFns ...func(*Options)) (*CreateBackupVaultOutput, error) {
	if params == nil {
		params = &CreateBackupVaultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBackupVault", params, optFns, addOperationCreateBackupVaultMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBackupVaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateBackupVaultInput struct {

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// AWS Region where they are created. They consist of lowercase letters, numbers,
	// and hyphens.
	//
	// This member is required.
	BackupVaultName *string

	// Metadata that you can assign to help organize the resources that you create.
	// Each tag is a key-value pair.
	BackupVaultTags map[string]*string

	// A unique string that identifies the request and allows failed requests to be
	// retried without the risk of executing the operation twice.
	CreatorRequestId *string

	// The server-side encryption key that is used to protect your backups; for
	// example,
	// arn:aws:kms:us-west-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab.
	EncryptionKeyArn *string
}

type CreateBackupVaultOutput struct {

	// An Amazon Resource Name (ARN) that uniquely identifies a backup vault; for
	// example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.
	BackupVaultArn *string

	// The name of a logical container where backups are stored. Backup vaults are
	// identified by names that are unique to the account used to create them and the
	// Region where they are created. They consist of lowercase letters, numbers, and
	// hyphens.
	BackupVaultName *string

	// The date and time a backup vault is created, in Unix format and Coordinated
	// Universal Time (UTC). The value of CreationDate is accurate to milliseconds. For
	// example, the value 1516925490.087 represents Friday, January 26, 2018
	// 12:11:30.087 AM.
	CreationDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateBackupVaultMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBackupVault{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBackupVault{}, middleware.After)
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
	addOpCreateBackupVaultValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBackupVault(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateBackupVault(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "CreateBackupVault",
	}
}
