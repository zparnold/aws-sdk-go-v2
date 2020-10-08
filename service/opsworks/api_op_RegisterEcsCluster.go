// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Registers a specified Amazon ECS cluster with a stack. You can register only one
// cluster with a stack. A cluster can be registered with only one stack. For more
// information, see  Resource Management
// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinglayers-ecscluster.html).
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack or an attached policy that explicitly grants
// permissions. For more information on user permissions, see  Managing User
// Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) RegisterEcsCluster(ctx context.Context, params *RegisterEcsClusterInput, optFns ...func(*Options)) (*RegisterEcsClusterOutput, error) {
	if params == nil {
		params = &RegisterEcsClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterEcsCluster", params, optFns, addOperationRegisterEcsClusterMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterEcsClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterEcsClusterInput struct {

	// The cluster's ARN.
	//
	// This member is required.
	EcsClusterArn *string

	// The stack ID.
	//
	// This member is required.
	StackId *string
}

// Contains the response to a RegisterEcsCluster request.
type RegisterEcsClusterOutput struct {

	// The cluster's ARN.
	EcsClusterArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterEcsClusterMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterEcsCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterEcsCluster{}, middleware.After)
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
	addOpRegisterEcsClusterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterEcsCluster(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRegisterEcsCluster(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "RegisterEcsCluster",
	}
}
