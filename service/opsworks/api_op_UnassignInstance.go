// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Unassigns a registered instance from all layers that are using the instance. The
// instance remains in the stack as an unassigned instance, and can be assigned to
// another layer as needed. You cannot use this action with instances that were
// created with AWS OpsWorks Stacks. Required Permissions: To use this action, an
// IAM user must have a Manage permissions level for the stack or an attached
// policy that explicitly grants permissions. For more information about user
// permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) UnassignInstance(ctx context.Context, params *UnassignInstanceInput, optFns ...func(*Options)) (*UnassignInstanceOutput, error) {
	if params == nil {
		params = &UnassignInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UnassignInstance", params, optFns, addOperationUnassignInstanceMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*UnassignInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UnassignInstanceInput struct {

	// The instance ID.
	//
	// This member is required.
	InstanceId *string
}

type UnassignInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUnassignInstanceMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUnassignInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUnassignInstance{}, middleware.After)
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
	addOpUnassignInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUnassignInstance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUnassignInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "UnassignInstance",
	}
}
