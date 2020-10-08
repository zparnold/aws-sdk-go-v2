// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a policy of a specified type that you can attach to a root, an
// organizational unit (OU), or an individual AWS account. For more information
// about policies and their use, see Managing Organization Policies
// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies.html).
// This operation can be called only from the organization's master account.
func (c *Client) CreatePolicy(ctx context.Context, params *CreatePolicyInput, optFns ...func(*Options)) (*CreatePolicyOutput, error) {
	if params == nil {
		params = &CreatePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreatePolicy", params, optFns, addOperationCreatePolicyMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreatePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreatePolicyInput struct {

	// The policy text content to add to the new policy. The text that you supply must
	// adhere to the rules of the policy type you specify in the Type parameter.
	//
	// This member is required.
	Content *string

	// An optional description to assign to the policy.
	//
	// This member is required.
	Description *string

	// The friendly name to assign to the policy. The regex pattern
	// (http://wikipedia.org/wiki/regex) that is used to validate this parameter is a
	// string of any of the characters in the ASCII character range.
	//
	// This member is required.
	Name *string

	// The type of policy to create. You can specify one of the following values:
	//
	//
	// * AISERVICES_OPT_OUT_POLICY
	// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_ai-opt-out.html)
	//
	//
	// * BACKUP_POLICY
	// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_backup.html)
	//
	//
	// * SERVICE_CONTROL_POLICY
	// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_scp.html)
	//
	//
	// * TAG_POLICY
	// (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html)
	//
	// This member is required.
	Type types.PolicyType
}

type CreatePolicyOutput struct {

	// A structure that contains details about the newly created policy.
	Policy *types.Policy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreatePolicyMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreatePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreatePolicy{}, middleware.After)
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
	addOpCreatePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreatePolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreatePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "CreatePolicy",
	}
}
