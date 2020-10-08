// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This action is only used by the Amazon ECS agent, and it is not intended for use
// outside of the agent. Registers an EC2 instance into the specified cluster. This
// instance becomes available to place containers on.
func (c *Client) RegisterContainerInstance(ctx context.Context, params *RegisterContainerInstanceInput, optFns ...func(*Options)) (*RegisterContainerInstanceOutput, error) {
	if params == nil {
		params = &RegisterContainerInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterContainerInstance", params, optFns, addOperationRegisterContainerInstanceMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterContainerInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterContainerInstanceInput struct {

	// The container instance attributes that this container instance supports.
	Attributes []*types.Attribute

	// The short name or full Amazon Resource Name (ARN) of the cluster with which to
	// register your container instance. If you do not specify a cluster, the default
	// cluster is assumed.
	Cluster *string

	// The ARN of the container instance (if it was previously registered).
	ContainerInstanceArn *string

	// The instance identity document for the EC2 instance to register. This document
	// can be found by running the following command from the instance: curl
	// http://169.254.169.254/latest/dynamic/instance-identity/document/
	InstanceIdentityDocument *string

	// The instance identity document signature for the EC2 instance to register. This
	// signature can be found by running the following command from the instance: curl
	// http://169.254.169.254/latest/dynamic/instance-identity/signature/
	InstanceIdentityDocumentSignature *string

	// The devices that are available on the container instance. The only supported
	// device type is a GPU.
	PlatformDevices []*types.PlatformDevice

	// The metadata that you apply to the container instance to help you categorize and
	// organize them. Each tag consists of a key and an optional value, both of which
	// you define. The following basic restrictions apply to tags:
	//
	//     * Maximum
	// number of tags per resource - 50
	//
	//     * For each resource, each tag key must be
	// unique, and each tag key can have only one value.
	//
	//     * Maximum key length -
	// 128 Unicode characters in UTF-8
	//
	//     * Maximum value length - 256 Unicode
	// characters in UTF-8
	//
	//     * If your tagging schema is used across multiple
	// services and resources, remember that other services may have restrictions on
	// allowed characters. Generally allowed characters are: letters, numbers, and
	// spaces representable in UTF-8, and the following characters: + - = . _ : / @.
	//
	//
	// * Tag keys and values are case-sensitive.
	//
	//     * Do not use aws:, AWS:, or any
	// upper or lowercase combination of such as a prefix for either keys or values as
	// it is reserved for AWS use. You cannot edit or delete tag keys or values with
	// this prefix. Tags with this prefix do not count against your tags per resource
	// limit.
	Tags []*types.Tag

	// The resources available on the instance.
	TotalResources []*types.Resource

	// The version information for the Amazon ECS container agent and Docker daemon
	// running on the container instance.
	VersionInfo *types.VersionInfo
}

type RegisterContainerInstanceOutput struct {

	// The container instance that was registered.
	ContainerInstance *types.ContainerInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterContainerInstanceMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterContainerInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterContainerInstance{}, middleware.After)
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
	addOpRegisterContainerInstanceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterContainerInstance(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRegisterContainerInstance(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "RegisterContainerInstance",
	}
}
