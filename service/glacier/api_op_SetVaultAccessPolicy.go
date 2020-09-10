// Code generated by smithy-go-codegen DO NOT EDIT.

package glacier

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation configures an access policy for a vault and will overwrite an
// existing policy. To configure a vault access policy, send a PUT request to the
// access-policy subresource of the vault. An access policy is specific to a vault
// and is also called a vault subresource. You can set one access policy per vault
// and the policy can be up to 20 KB in size. For more information about vault
// access policies, see Amazon Glacier Access Control with Vault Access Policies
// (https://docs.aws.amazon.com/amazonglacier/latest/dev/vault-access-policy.html).
func (c *Client) SetVaultAccessPolicy(ctx context.Context, params *SetVaultAccessPolicyInput, optFns ...func(*Options)) (*SetVaultAccessPolicyOutput, error) {
	stack := middleware.NewStack("SetVaultAccessPolicy", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpSetVaultAccessPolicyMiddlewares(stack)
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
	addOpSetVaultAccessPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSetVaultAccessPolicy(options.Region), middleware.Before)

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
			OperationName: "SetVaultAccessPolicy",
			Err:           err,
		}
	}
	out := result.(*SetVaultAccessPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// SetVaultAccessPolicy input.
type SetVaultAccessPolicyInput struct {
	// The vault access policy as a JSON string.
	Policy *types.VaultAccessPolicy
	// The AccountId value is the AWS account ID of the account that owns the vault.
	// You can either specify an AWS account ID or optionally a single '-' (hyphen), in
	// which case Amazon S3 Glacier uses the AWS account ID associated with the
	// credentials used to sign the request. If you use an account ID, do not include
	// any hyphens ('-') in the ID.
	AccountId *string
	// The name of the vault.
	VaultName *string
}

type SetVaultAccessPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpSetVaultAccessPolicyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpSetVaultAccessPolicy{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpSetVaultAccessPolicy{}, middleware.After)
}

func newServiceMetadataMiddleware_opSetVaultAccessPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glacier",
		OperationName: "SetVaultAccessPolicy",
	}
}