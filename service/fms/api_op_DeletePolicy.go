// Code generated by smithy-go-codegen DO NOT EDIT.

package fms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Permanently deletes an AWS Firewall Manager policy.
func (c *Client) DeletePolicy(ctx context.Context, params *DeletePolicyInput, optFns ...func(*Options)) (*DeletePolicyOutput, error) {
	if params == nil {
		params = &DeletePolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeletePolicy", params, optFns, addOperationDeletePolicyMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeletePolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeletePolicyInput struct {

	// The ID of the policy that you want to delete. You can retrieve this ID from
	// PutPolicy and ListPolicies.
	//
	// This member is required.
	PolicyId *string

	// If True, the request performs cleanup according to the policy type. For AWS WAF
	// and Shield Advanced policies, the cleanup does the following:
	//
	//     * Deletes
	// rule groups created by AWS Firewall Manager
	//
	//     * Removes web ACLs from
	// in-scope resources
	//
	//     * Deletes web ACLs that contain no rules or rule
	// groups
	//
	// For security group policies, the cleanup does the following for each
	// security group in the policy:
	//
	//     * Disassociates the security group from
	// in-scope resources
	//
	//     * Deletes the security group if it was created through
	// Firewall Manager and if it's no longer associated with any resources through
	// another policy
	//
	// After the cleanup, in-scope resources are no longer protected by
	// web ACLs in this policy. Protection of out-of-scope resources remains unchanged.
	// Scope is determined by tags that you create and accounts that you associate with
	// the policy. When creating the policy, if you specify that only resources in
	// specific accounts or with specific tags are in scope of the policy, those
	// accounts and resources are handled by the policy. All others are out of scope.
	// If you don't specify tags or accounts, all resources are in scope.
	DeleteAllPolicyResources *bool
}

type DeletePolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeletePolicyMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeletePolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeletePolicy{}, middleware.After)
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
	addOpDeletePolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeletePolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeletePolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "fms",
		OperationName: "DeletePolicy",
	}
}
