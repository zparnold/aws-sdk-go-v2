// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a contact with the specified details.
func (c *Client) CreateContact(ctx context.Context, params *CreateContactInput, optFns ...func(*Options)) (*CreateContactOutput, error) {
	if params == nil {
		params = &CreateContactInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateContact", params, optFns, addOperationCreateContactMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateContactOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateContactInput struct {

	// The first name of the contact that is used to call the contact on the device.
	//
	// This member is required.
	FirstName *string

	// A unique, user-specified identifier for this request that ensures idempotency.
	ClientRequestToken *string

	// The name of the contact to display on the console.
	DisplayName *string

	// The last name of the contact that is used to call the contact on the device.
	LastName *string

	// The phone number of the contact in E.164 format. The phone number type defaults
	// to WORK. You can specify PhoneNumber or PhoneNumbers. We recommend that you use
	// PhoneNumbers, which lets you specify the phone number type and multiple numbers.
	PhoneNumber *string

	// The list of phone numbers for the contact.
	PhoneNumbers []*types.PhoneNumber

	// The list of SIP addresses for the contact.
	SipAddresses []*types.SipAddress
}

type CreateContactOutput struct {

	// The ARN of the newly created address book.
	ContactArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateContactMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateContact{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateContact{}, middleware.After)
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
	addIdempotencyToken_opCreateContactMiddleware(stack, options)
	addOpCreateContactValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateContact(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpCreateContact struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateContact) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateContact) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateContactInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateContactInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateContactMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateContact{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateContact(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "CreateContact",
	}
}
