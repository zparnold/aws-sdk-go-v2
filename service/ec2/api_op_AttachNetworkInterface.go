// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Attaches a network interface to an instance.
func (c *Client) AttachNetworkInterface(ctx context.Context, params *AttachNetworkInterfaceInput, optFns ...func(*Options)) (*AttachNetworkInterfaceOutput, error) {
	stack := middleware.NewStack("AttachNetworkInterface", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpAttachNetworkInterfaceMiddlewares(stack)
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
	addOpAttachNetworkInterfaceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAttachNetworkInterface(options.Region), middleware.Before)

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
			OperationName: "AttachNetworkInterface",
			Err:           err,
		}
	}
	out := result.(*AttachNetworkInterfaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for AttachNetworkInterface.
type AttachNetworkInterfaceInput struct {
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// The ID of the instance.
	InstanceId *string
	// The index of the device for the network interface attachment.
	DeviceIndex *int32
	// The ID of the network interface.
	NetworkInterfaceId *string
}

// Contains the output of AttachNetworkInterface.
type AttachNetworkInterfaceOutput struct {
	// The ID of the network interface attachment.
	AttachmentId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpAttachNetworkInterfaceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpAttachNetworkInterface{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpAttachNetworkInterface{}, middleware.After)
}

func newServiceMetadataMiddleware_opAttachNetworkInterface(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "AttachNetworkInterface",
	}
}
