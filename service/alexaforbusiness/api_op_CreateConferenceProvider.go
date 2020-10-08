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

// Adds a new conference provider under the user's AWS account.
func (c *Client) CreateConferenceProvider(ctx context.Context, params *CreateConferenceProviderInput, optFns ...func(*Options)) (*CreateConferenceProviderOutput, error) {
	if params == nil {
		params = &CreateConferenceProviderInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateConferenceProvider", params, optFns, addOperationCreateConferenceProviderMiddleware)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateConferenceProviderOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConferenceProviderInput struct {

	// The name of the conference provider.
	//
	// This member is required.
	ConferenceProviderName *string

	// Represents a type within a list of predefined types.
	//
	// This member is required.
	ConferenceProviderType types.ConferenceProviderType

	// The meeting settings for the conference provider.
	//
	// This member is required.
	MeetingSetting *types.MeetingSetting

	// The request token of the client.
	ClientRequestToken *string

	// The IP endpoint and protocol for calling.
	IPDialIn *types.IPDialIn

	// The information for PSTN conferencing.
	PSTNDialIn *types.PSTNDialIn
}

type CreateConferenceProviderOutput struct {

	// The ARN of the newly-created conference provider.
	ConferenceProviderArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateConferenceProviderMiddleware(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateConferenceProvider{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateConferenceProvider{}, middleware.After)
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
	addIdempotencyToken_opCreateConferenceProviderMiddleware(stack, options)
	addOpCreateConferenceProviderValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConferenceProvider(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpCreateConferenceProvider struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateConferenceProvider) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateConferenceProvider) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateConferenceProviderInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateConferenceProviderInput ")
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
func addIdempotencyToken_opCreateConferenceProviderMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateConferenceProvider{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateConferenceProvider(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "CreateConferenceProvider",
	}
}
