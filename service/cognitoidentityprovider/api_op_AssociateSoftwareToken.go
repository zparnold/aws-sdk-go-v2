// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a unique generated shared secret key code for the user account. The
// request takes an access token or a session string, but not both.
func (c *Client) AssociateSoftwareToken(ctx context.Context, params *AssociateSoftwareTokenInput, optFns ...func(*Options)) (*AssociateSoftwareTokenOutput, error) {
	stack := middleware.NewStack("AssociateSoftwareToken", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssociateSoftwareTokenMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateSoftwareToken(options.Region), middleware.Before)
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
			OperationName: "AssociateSoftwareToken",
			Err:           err,
		}
	}
	out := result.(*AssociateSoftwareTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateSoftwareTokenInput struct {

	// The access token.
	AccessToken *string

	// The session which should be passed both ways in challenge-response calls to the
	// service. This allows authentication of the user as part of the MFA setup
	// process.
	Session *string
}

type AssociateSoftwareTokenOutput struct {

	// A unique generated shared secret code that is used in the TOTP algorithm to
	// generate a one time code.
	SecretCode *string

	// The session which should be passed both ways in challenge-response calls to the
	// service. This allows authentication of the user as part of the MFA setup
	// process.
	Session *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssociateSoftwareTokenMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateSoftwareToken{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateSoftwareToken{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssociateSoftwareToken(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AssociateSoftwareToken",
	}
}
