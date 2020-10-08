// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the IAM policy that is attached to the specified rule group. You must be
// the owner of the rule group to perform this operation.
func (c *Client) GetPermissionPolicy(ctx context.Context, params *GetPermissionPolicyInput, optFns ...func(*Options)) (*GetPermissionPolicyOutput, error) {
	if params == nil {
		params = &GetPermissionPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPermissionPolicy", params, optFns, addOperationGetPermissionPolicyMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPermissionPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPermissionPolicyInput struct {

	// The Amazon Resource Name (ARN) of the rule group for which you want to get the
	// policy.
	//
	// This member is required.
	ResourceArn *string
}

type GetPermissionPolicyOutput struct {

	// The IAM policy that is attached to the specified rule group.
	Policy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetPermissionPolicyMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPermissionPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPermissionPolicy{}, middleware.After)
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
	addOpGetPermissionPolicyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPermissionPolicy(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetPermissionPolicy(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "GetPermissionPolicy",
	}
}
