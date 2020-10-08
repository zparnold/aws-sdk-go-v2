// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disables the specified user. Calling this action requires developer credentials.
func (c *Client) AdminDisableUser(ctx context.Context, params *AdminDisableUserInput, optFns ...func(*Options)) (*AdminDisableUserOutput, error) {
	if params == nil {
		params = &AdminDisableUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AdminDisableUser", params, optFns, addOperationAdminDisableUserMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*AdminDisableUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to disable any user as an administrator.
type AdminDisableUserInput struct {

	// The user pool ID for the user pool where you want to disable the user.
	//
	// This member is required.
	UserPoolId *string

	// The user name of the user you wish to disable.
	//
	// This member is required.
	Username *string
}

// Represents the response received from the server to disable the user as an
// administrator.
type AdminDisableUserOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAdminDisableUserMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAdminDisableUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminDisableUser{}, middleware.After)
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
	addOpAdminDisableUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminDisableUser(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAdminDisableUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminDisableUser",
	}
}
