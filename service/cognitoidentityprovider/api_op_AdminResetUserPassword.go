// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Resets the specified user's password in a user pool as an administrator. Works
// on any user. When a developer calls this API, the current password is
// invalidated, so it must be changed. If a user tries to sign in after the API is
// called, the app will get a PasswordResetRequiredException exception back and
// should direct the user down the flow to reset the password, which is the same as
// the forgot password flow. In addition, if the user pool has phone verification
// selected and a verified phone number exists for the user, or if email
// verification is selected and a verified email exists for the user, calling this
// API will also result in sending a message to the end user with the code to
// change their password. Calling this action requires developer credentials.
func (c *Client) AdminResetUserPassword(ctx context.Context, params *AdminResetUserPasswordInput, optFns ...func(*Options)) (*AdminResetUserPasswordOutput, error) {
	stack := middleware.NewStack("AdminResetUserPassword", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAdminResetUserPasswordMiddlewares(stack)
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
	addOpAdminResetUserPasswordValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminResetUserPassword(options.Region), middleware.Before)

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
			OperationName: "AdminResetUserPassword",
			Err:           err,
		}
	}
	out := result.(*AdminResetUserPasswordOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to reset a user's password as an administrator.
type AdminResetUserPasswordInput struct {
	// The user pool ID for the user pool where you want to reset the user's password.
	UserPoolId *string
	// A map of custom key-value pairs that you can provide as input for any custom
	// workflows that this action triggers. You create custom workflows by assigning
	// AWS Lambda functions to user pool triggers. When you use the
	// AdminResetUserPassword API action, Amazon Cognito invokes the function that is
	// assigned to the custom message trigger. When Amazon Cognito invokes this
	// function, it passes a JSON payload, which the function receives as input. This
	// payload contains a clientMetadata attribute, which provides the data that you
	// assigned to the ClientMetadata parameter in your AdminResetUserPassword request.
	// In your function code in AWS Lambda, you can process the clientMetadata value to
	// enhance your workflow for your specific needs. For more information, see
	// Customizing User Pool Workflows with Lambda Triggers
	// (https://docs.aws.amazon.com/cognito/latest/developerguide/cognito-user-identity-pools-working-with-aws-lambda-triggers.html)
	// in the Amazon Cognito Developer Guide. Take the following limitations into
	// consideration when you use the ClientMetadata parameter:
	//
	//     * Amazon Cognito
	// does not store the ClientMetadata value. This data is available only to AWS
	// Lambda triggers that are assigned to a user pool to support custom workflows. If
	// your user pool configuration does not include triggers, the ClientMetadata
	// parameter serves no purpose.
	//
	//     * Amazon Cognito does not validate the
	// ClientMetadata value.
	//
	//     * Amazon Cognito does not encrypt the the
	// ClientMetadata value, so don't use it to provide sensitive information.
	ClientMetadata map[string]*string
	// The user name of the user whose password you wish to reset.
	Username *string
}

// Represents the response from the server to reset a user password as an
// administrator.
type AdminResetUserPasswordOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAdminResetUserPasswordMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAdminResetUserPassword{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminResetUserPassword{}, middleware.After)
}

func newServiceMetadataMiddleware_opAdminResetUserPassword(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminResetUserPassword",
	}
}
