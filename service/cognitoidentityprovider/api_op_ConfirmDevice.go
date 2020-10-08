// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Confirms tracking of the device. This API call is the call that begins device
// tracking.
func (c *Client) ConfirmDevice(ctx context.Context, params *ConfirmDeviceInput, optFns ...func(*Options)) (*ConfirmDeviceOutput, error) {
	if params == nil {
		params = &ConfirmDeviceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ConfirmDevice", params, optFns, addOperationConfirmDeviceMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*ConfirmDeviceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Confirms the device request.
type ConfirmDeviceInput struct {

	// The access token.
	//
	// This member is required.
	AccessToken *string

	// The device key.
	//
	// This member is required.
	DeviceKey *string

	// The device name.
	DeviceName *string

	// The configuration of the device secret verifier.
	DeviceSecretVerifierConfig *types.DeviceSecretVerifierConfigType
}

// Confirms the device response.
type ConfirmDeviceOutput struct {

	// Indicates whether the user confirmation is necessary to confirm the device
	// response.
	UserConfirmationNecessary *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationConfirmDeviceMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpConfirmDevice{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpConfirmDevice{}, middleware.After)
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
	addOpConfirmDeviceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opConfirmDevice(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opConfirmDevice(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "ConfirmDevice",
	}
}
