// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables I/O operations for a volume that had I/O operations disabled because the
// data on the volume was potentially inconsistent.
func (c *Client) EnableVolumeIO(ctx context.Context, params *EnableVolumeIOInput, optFns ...func(*Options)) (*EnableVolumeIOOutput, error) {
	if params == nil {
		params = &EnableVolumeIOInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableVolumeIO", params, optFns, addOperationEnableVolumeIOMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableVolumeIOOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableVolumeIOInput struct {

	// The ID of the volume.
	//
	// This member is required.
	VolumeId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type EnableVolumeIOOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableVolumeIOMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpEnableVolumeIO{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpEnableVolumeIO{}, middleware.After)
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
	addOpEnableVolumeIOValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableVolumeIO(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opEnableVolumeIO(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "EnableVolumeIO",
	}
}
