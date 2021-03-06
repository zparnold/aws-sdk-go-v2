// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables the specified MFA device and associates it with the specified IAM user.
// When enabled, the MFA device is required for every subsequent login by the IAM
// user associated with the device.
func (c *Client) EnableMFADevice(ctx context.Context, params *EnableMFADeviceInput, optFns ...func(*Options)) (*EnableMFADeviceOutput, error) {
	stack := middleware.NewStack("EnableMFADevice", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpEnableMFADeviceMiddlewares(stack)
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
	addOpEnableMFADeviceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableMFADevice(options.Region), middleware.Before)
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
			OperationName: "EnableMFADevice",
			Err:           err,
		}
	}
	out := result.(*EnableMFADeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableMFADeviceInput struct {

	// An authentication code emitted by the device. The format for this parameter is a
	// string of six digits. Submit your request immediately after generating the
	// authentication codes. If you generate the codes and then wait too long to submit
	// the request, the MFA device successfully associates with the user but the MFA
	// device becomes out of sync. This happens because time-based one-time passwords
	// (TOTP) expire after a short period of time. If this happens, you can resync the
	// device
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_mfa_sync.html).
	//
	// This member is required.
	AuthenticationCode1 *string

	// A subsequent authentication code emitted by the device. The format for this
	// parameter is a string of six digits. Submit your request immediately after
	// generating the authentication codes. If you generate the codes and then wait too
	// long to submit the request, the MFA device successfully associates with the user
	// but the MFA device becomes out of sync. This happens because time-based one-time
	// passwords (TOTP) expire after a short period of time. If this happens, you can
	// resync the device
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_mfa_sync.html).
	//
	// This member is required.
	AuthenticationCode2 *string

	// The serial number that uniquely identifies the MFA device. For virtual MFA
	// devices, the serial number is the device ARN. This parameter allows (through its
	// regex pattern (http://wikipedia.org/wiki/regex)) a string of characters
	// consisting of upper and lowercase alphanumeric characters with no spaces. You
	// can also include any of the following characters: =,.@:/-
	//
	// This member is required.
	SerialNumber *string

	// The name of the IAM user for whom you want to enable the MFA device. This
	// parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a
	// string of characters consisting of upper and lowercase alphanumeric characters
	// with no spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	UserName *string
}

type EnableMFADeviceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpEnableMFADeviceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpEnableMFADevice{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpEnableMFADevice{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableMFADevice(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "EnableMFADevice",
	}
}
