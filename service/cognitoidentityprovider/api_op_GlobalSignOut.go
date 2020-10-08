// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Signs out users from all devices. It also invalidates all refresh tokens issued
// to a user. The user's current access and Id tokens remain valid until their
// expiry. Access and Id tokens expire one hour after they are issued.
func (c *Client) GlobalSignOut(ctx context.Context, params *GlobalSignOutInput, optFns ...func(*Options)) (*GlobalSignOutOutput, error) {
	if params == nil {
		params = &GlobalSignOutInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GlobalSignOut", params, optFns, addOperationGlobalSignOutMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*GlobalSignOutOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to sign out all devices.
type GlobalSignOutInput struct {

	// The access token.
	//
	// This member is required.
	AccessToken *string
}

// The response to the request to sign out all devices.
type GlobalSignOutOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGlobalSignOutMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGlobalSignOut{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGlobalSignOut{}, middleware.After)
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
	addOpGlobalSignOutValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGlobalSignOut(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGlobalSignOut(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "GlobalSignOut",
	}
}
