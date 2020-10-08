// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation returns information about a vault, including the vault's Amazon
// Resource Name (ARN), the date the vault was created, the number of archives it
// contains, and the total size of all the archives in the vault. The number of
// archives and their total size are as of the last inventory generation. This
// means that if you add or remove an archive from a vault, and then immediately
// use Describe Vault, the change in contents will not be immediately reflected. If
// you want to retrieve the latest inventory of the vault, use InitiateJob ().
// Amazon S3 Glacier generates vault inventories approximately daily. For more
// information, see Downloading a Vault Inventory in Amazon S3 Glacier
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-inventory.html).
// <p>An AWS account has full permission to perform all operations (actions).
// However, AWS Identity and Access Management (IAM) users don't have any
// permissions by default. You must grant them explicit permission to perform
// specific actions. For more information, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html">Access
// Control Using AWS Identity and Access Management (IAM)</a>.</p> <p>For
// conceptual information and underlying REST API, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/retrieving-vault-info.html">Retrieving
// Vault Metadata in Amazon S3 Glacier</a> and <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-get.html">Describe
// Vault </a> in the <i>Amazon Glacier Developer Guide</i>. </p>
func (c *Client) DescribeVault(ctx context.Context, params *DescribeVaultInput, optFns ...func(*Options)) (*DescribeVaultOutput, error) {
	if params == nil {
		params = &DescribeVaultInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVault", params, optFns, addOperationDescribeVaultMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options for retrieving metadata for a specific vault in Amazon Glacier.
type DescribeVaultInput struct {

	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The name of the vault.
	//
	// This member is required.
	VaultName *string
}

// Contains the Amazon S3 Glacier response to your request.
type DescribeVaultOutput struct {

	// The Universal Coordinated Time (UTC) date when the vault was created. This value
	// should be a string in the ISO 8601 date format, for example
	// 2012-03-20T17:03:43.221Z.
	CreationDate *string

	// The Universal Coordinated Time (UTC) date when Amazon S3 Glacier completed the
	// last vault inventory. This value should be a string in the ISO 8601 date format,
	// for example 2012-03-20T17:03:43.221Z.
	LastInventoryDate *string

	// The number of archives in the vault as of the last inventory date. This field
	// will return null if an inventory has not yet run on the vault, for example if
	// you just created the vault.
	NumberOfArchives *int64

	// Total size, in bytes, of the archives in the vault as of the last inventory
	// date. This field will return null if an inventory has not yet run on the vault,
	// for example if you just created the vault.
	SizeInBytes *int64

	// The Amazon Resource Name (ARN) of the vault.
	VaultARN *string

	// The name of the vault.
	VaultName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVaultMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeVault{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeVault{}, middleware.After)
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
	addOpDescribeVaultValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVault(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	glaciercust.AddTreeHashMiddleware(stack)
	glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion)
	glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID)
	return nil
}

func newServiceMetadataMiddleware_opDescribeVault(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "DescribeVault",
	}
}
