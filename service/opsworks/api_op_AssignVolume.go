// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Assigns one of the stack's registered Amazon EBS volumes to a specified
// instance. The volume must first be registered with the stack by calling
// RegisterVolume (). After you register the volume, you must call UpdateVolume ()
// to specify a mount point before calling AssignVolume. For more information, see
// Resource Management
// (https://docs.aws.amazon.com/opsworks/latest/userguide/resources.html). Required
// Permissions: To use this action, an IAM user must have a Manage permissions
// level for the stack, or an attached policy that explicitly grants permissions.
// For more information on user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) AssignVolume(ctx context.Context, params *AssignVolumeInput, optFns ...func(*Options)) (*AssignVolumeOutput, error) {
	stack := middleware.NewStack("AssignVolume", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssignVolumeMiddlewares(stack)
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
	addOpAssignVolumeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssignVolume(options.Region), middleware.Before)
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
			OperationName: "AssignVolume",
			Err:           err,
		}
	}
	out := result.(*AssignVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssignVolumeInput struct {

	// The volume ID.
	//
	// This member is required.
	VolumeId *string

	// The instance ID.
	InstanceId *string
}

type AssignVolumeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssignVolumeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssignVolume{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssignVolume{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssignVolume(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "AssignVolume",
	}
}
