// Code generated by smithy-go-codegen DO NOT EDIT.

package globalaccelerator

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Delete an accelerator. Before you can delete an accelerator, you must disable it
// and remove all dependent resources (listeners and endpoint groups). To disable
// the accelerator, update the accelerator to set Enabled to false. When you create
// an accelerator, by default, Global Accelerator provides you with a set of two
// static IP addresses. Alternatively, you can bring your own IP address ranges to
// Global Accelerator and assign IP addresses from those ranges. The IP addresses
// are assigned to your accelerator for as long as it exists, even if you disable
// the accelerator and it no longer accepts or routes traffic. However, when you
// delete an accelerator, you lose the static IP addresses that are assigned to the
// accelerator, so you can no longer route traffic by using them. As a best
// practice, ensure that you have permissions in place to avoid inadvertently
// deleting accelerators. You can use IAM policies with Global Accelerator to limit
// the users who have permissions to delete an accelerator. For more information,
// see Authentication and Access Control
// (https://docs.aws.amazon.com/global-accelerator/latest/dg/auth-and-access-control.html)
// in the AWS Global Accelerator Developer Guide.
func (c *Client) DeleteAccelerator(ctx context.Context, params *DeleteAcceleratorInput, optFns ...func(*Options)) (*DeleteAcceleratorOutput, error) {
	if params == nil {
		params = &DeleteAcceleratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteAccelerator", params, optFns, addOperationDeleteAcceleratorMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteAcceleratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteAcceleratorInput struct {

	// The Amazon Resource Name (ARN) of an accelerator.
	//
	// This member is required.
	AcceleratorArn *string
}

type DeleteAcceleratorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteAcceleratorMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteAccelerator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteAccelerator{}, middleware.After)
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
	addOpDeleteAcceleratorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteAccelerator(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteAccelerator(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "globalaccelerator",
		OperationName: "DeleteAccelerator",
	}
}
