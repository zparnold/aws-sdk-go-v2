// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	glaciercust "github.com/aws/aws-sdk-go-v2/service/glacier/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation lists all vaults owned by the calling user's account. The list
// returned in the response is ASCII-sorted by vault name.  <p>By default, this
// operation returns up to 10 items. If there are more vaults to list, the response
// <code>marker</code> field contains the vault Amazon Resource Name (ARN) at which
// to continue the list with a new List Vaults request; otherwise, the
// <code>marker</code> field is <code>null</code>. To return a list of vaults that
// begins at a specific vault, set the <code>marker</code> request parameter to the
// vault ARN you obtained from a previous List Vaults request. You can also limit
// the number of vaults returned in the response by specifying the
// <code>limit</code> parameter in the request. </p> <p>An AWS account has full
// permission to perform all operations (actions). However, AWS Identity and Access
// Management (IAM) users don't have any permissions by default. You must grant
// them explicit permission to perform specific actions. For more information, see
// <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html">Access
// Control Using AWS Identity and Access Management (IAM)</a>.</p> <p>For
// conceptual information and underlying REST API, see <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/retrieving-vault-info.html">Retrieving
// Vault Metadata in Amazon S3 Glacier</a> and <a
// href="https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vaults-get.html">List
// Vaults </a> in the <i>Amazon Glacier Developer Guide</i>. </p>
func (c *Client) ListVaults(ctx context.Context, params *ListVaultsInput, optFns ...func(*Options)) (*ListVaultsOutput, error) {
	if params == nil {
		params = &ListVaultsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVaults", params, optFns, addOperationListVaultsMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVaultsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Provides options to retrieve the vault list owned by the calling user's account.
// The list provides metadata information for each vault.
type ListVaultsInput struct {

	// The AccountId value is the AWS account ID. This value must match the AWS account
	// ID associated with the credentials used to sign the request. You can either
	// specify an AWS account ID or optionally a single '-' (hyphen), in which case
	// Amazon Glacier uses the AWS account ID associated with the credentials used to
	// sign the request. If you specify your account ID, do not include any hyphens
	// ('-') in the ID.
	//
	// This member is required.
	AccountId *string

	// The maximum number of vaults to be returned. The default limit is 10. The number
	// of vaults returned might be fewer than the specified limit, but the number of
	// returned vaults never exceeds the limit.
	Limit *string

	// A string used for pagination. The marker specifies the vault ARN after which the
	// listing of vaults should begin.
	Marker *string
}

// Contains the Amazon S3 Glacier response to your request.
type ListVaultsOutput struct {

	// The vault ARN at which to continue pagination of the results. You use the marker
	// in another List Vaults request to obtain more vaults in the list.
	Marker *string

	// List of vaults.
	VaultList []*types.DescribeVaultOutput

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListVaultsMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVaults{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVaults{}, middleware.After)
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
	addOpListVaultsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListVaults(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	glaciercust.AddTreeHashMiddleware(stack)
	glaciercust.AddGlacierAPIVersionMiddleware(stack, ServiceAPIVersion)
	glaciercust.AddDefaultAccountIDMiddleware(stack, setDefaultAccountID)
	return nil
}

func newServiceMetadataMiddleware_opListVaults(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "ListVaults",
	}
}
